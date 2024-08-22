package auth

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/sourcegraph/sourcegraph/internal/encryption"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/auth"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/github"
	"github.com/sourcegraph/sourcegraph/internal/httpcli"
	"github.com/sourcegraph/sourcegraph/internal/rcache"
	"github.com/sourcegraph/sourcegraph/internal/redispool"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

// GitHubAppAuthenticator is used to authenticate requests to the GitHub API
// using a GitHub App. It contains the ID and private key associated with
// the GitHub App.
type GitHubAppAuthenticator struct {
	appID  int
	key    *rsa.PrivateKey
	rawKey []byte
}

var _ auth.Authenticator = (*GitHubAppAuthenticator)(nil)
var _ auth.Authenticator = (*InstallationAuthenticator)(nil)

// NewGitHubAppAuthenticator creates an Authenticator that can be used to authenticate requests
// to the GitHub API as a GitHub App. It requires the GitHub App ID and RSA private key.
//
// The returned Authenticator can be used to sign requests to the GitHub API on behalf of the GitHub App.
// The requests will contain a JSON Web Token (JWT) in the Authorization header with claims identifying
// the GitHub App.
func NewGitHubAppAuthenticator(appID int, privateKey []byte) (*GitHubAppAuthenticator, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "parse private key")
	}
	return &GitHubAppAuthenticator{
		appID:  appID,
		key:    key,
		rawKey: privateKey,
	}, nil
}

// Authenticate adds an Authorization header to the request containing
// a JSON Web Token (JWT) signed with the GitHub App's private key.
// The JWT contains claims identifying the GitHub App.
func (a *GitHubAppAuthenticator) Authenticate(r *http.Request) error {
	token, err := a.generateJWT()
	if err != nil {
		return err
	}
	r.Header.Set("Authorization", "Bearer "+token)
	return nil
}

// generateJWT generates a JSON Web Token (JWT) signed with the GitHub App's private key.
// The JWT contains claims identifying the GitHub App.
//
// The payload computation is following GitHub App's Ruby example shown in
// https://docs.github.com/en/developers/apps/building-github-apps/authenticating-with-github-apps#authenticating-as-a-github-app.
//
// NOTE: GitHub rejects expiry and issue timestamps that are not an integer,
// while the jwt-go library serializes to fractional timestamps. Truncate them
// before passing to jwt-go.
//
// The returned JWT can be used to authenticate requests to the GitHub API as the GitHub App.
func (a *GitHubAppAuthenticator) generateJWT() (string, error) {
	iss := time.Now().Add(-time.Minute).Truncate(time.Second)
	exp := iss.Add(10 * time.Minute)
	claims := &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(iss),
		ExpiresAt: jwt.NewNumericDate(exp),
		Issuer:    strconv.Itoa(a.appID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(a.key)
}

func (a *GitHubAppAuthenticator) Hash() string {
	shaSum := sha256.Sum256(a.rawKey)
	return hex.EncodeToString(shaSum[:])
}

type InstallationAccessToken struct {
	Token     string
	ExpiresAt time.Time
}

type installationAccessToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// InstallationAuthenticator is used to authenticate requests to the
// GitHub API using an installation access token from a GitHub App.
//
// It implements the auth.Authenticator interface.
type InstallationAuthenticator struct {
	installationID          int
	installationAccessToken installationAccessToken
	baseURL                 *url.URL
	appAuthenticator        auth.Authenticator
	cache                   *rcache.Cache
	encryptionKey           encryption.Key
}

// NewInstallationAccessToken implements the Authenticator interface
// for GitHub App installations. Installation access tokens are created
// for the given installationID, using the provided authenticator.
//
// appAuthenticator must not be nil.
func NewInstallationAccessToken(
	baseURL *url.URL,
	installationID int,
	appAuthenticator auth.Authenticator,
	encryptionKey encryption.Key, // Used to encrypt the token before caching it
) *InstallationAuthenticator {
	cache := rcache.NewWithTTL(redispool.Cache, "github_app_installation_token", 55*60)
	return &InstallationAuthenticator{
		baseURL:          baseURL,
		installationID:   installationID,
		appAuthenticator: appAuthenticator,
		cache:            cache,
		encryptionKey:    encryptionKey,
	}
}

func (t *InstallationAuthenticator) cacheKey() string {
	return t.baseURL.String() + strconv.Itoa(t.installationID)
}

// getFromCache returns a new installationAccessToken from the cache, and a boolean
// indicating whether the fetch from cache was successful.
func (t *InstallationAuthenticator) getFromCache(ctx context.Context) (iat installationAccessToken, ok bool) {
	token, ok := t.cache.Get(t.cacheKey())
	if !ok {
		return
	}
	if t.encryptionKey != nil {
		encrypted, err := t.encryptionKey.Decrypt(ctx, token)
		if err != nil {
			return iat, false
		}
		token = []byte(encrypted.String())
	}

	if err := json.Unmarshal(token, &iat); err != nil {
		return
	}

	return iat, true
}

// storeInCache updates the installationAccessToken in the cache.
func (t *InstallationAuthenticator) storeInCache(ctx context.Context) error {
	res, err := json.Marshal(t.installationAccessToken)
	if err != nil {
		return err
	}
	if t.encryptionKey != nil {
		res, err = t.encryptionKey.Encrypt(ctx, res)
		if err != nil {
			return err
		}
	}

	t.cache.Set(t.cacheKey(), res)
	return nil
}

// Refresh generates a new installation access token for the GitHub App installation.
//
// It makes a request to the GitHub API to generate a new installation access token for the
// installation associated with the Authenticator.
// Returns an error if the request fails.
func (t *InstallationAuthenticator) Refresh(ctx context.Context, cli httpcli.Doer) error {
	token, ok := t.getFromCache(ctx)
	if ok {
		if t.installationAccessToken.Token != token.Token { // Confirm that we have a different token now
			t.installationAccessToken = token
			if !t.NeedsRefresh() {
				// Return nil, indicating the refresh was "successful"
				return nil
			}
		}
	}

	apiURL, _ := github.APIRoot(t.baseURL)
	apiURL = apiURL.JoinPath(fmt.Sprintf("/app/installations/%d/access_tokens", t.installationID))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL.String(), nil)
	if err != nil {
		return err
	}
	if err := t.appAuthenticator.Authenticate(req); err != nil {
		return err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.Newf("failed to refresh the access token: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&t.installationAccessToken); err != nil {
		return err
	}
	// Ignore if storing in cache fails somehow, since the token should still be valid
	_ = t.storeInCache(ctx)

	return nil
}

// Authenticate adds an Authorization header to the request containing
// the installation access token associated with the GitHub App installation.
func (t *InstallationAuthenticator) Authenticate(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer "+t.installationAccessToken.Token)
	return nil
}

// Hash returns a hash of the GitHub App installation ID.
// We use the installation ID instead of the installation access
// token because installation access tokens are short-lived.
func (t *InstallationAuthenticator) Hash() string {
	sum := sha256.Sum256([]byte(strconv.Itoa(t.installationID)))
	return hex.EncodeToString(sum[:])
}

// NeedsRefresh checks whether the GitHub App installation access token
// needs to be refreshed. An access token needs to be refreshed if it has
// expired or will expire within the next few seconds.
func (t *InstallationAuthenticator) NeedsRefresh() bool {
	if t.installationAccessToken.Token == "" {
		return true
	}
	if t.installationAccessToken.ExpiresAt.IsZero() {
		return false
	}
	return time.Until(t.installationAccessToken.ExpiresAt) < 5*time.Minute
}

// SetURLUser sets the URL's User field to contain the installation access token.
func (t *InstallationAuthenticator) SetURLUser(u *url.URL) {
	u.User = url.UserPassword("x-access-token", t.installationAccessToken.Token)
}

func (t *InstallationAuthenticator) GetToken() InstallationAccessToken {
	return InstallationAccessToken(t.installationAccessToken)
}

func (t *InstallationAuthenticator) InstallationID() int {
	return t.installationID
}
