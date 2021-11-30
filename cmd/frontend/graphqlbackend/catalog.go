package graphqlbackend

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
)

// This file just contains stub GraphQL resolvers and data types for the Catalog which merely return
// an error if not running in enterprise mode. The actual resolvers are in
// enterprise/cmd/frontend/internal/catalog/resolvers.

// CatalogRootResolver is the root resolver.
type CatalogRootResolver interface {
	Catalog(context.Context) (CatalogResolver, error)
	CatalogEntity(context.Context, *CatalogEntityArgs) (*CatalogEntityResolver, error)

	NodeResolvers() map[string]NodeByIDFunc
}

type CatalogEntityArgs struct {
	Name string
}

type CatalogResolver interface {
	Entities(context.Context, *CatalogEntitiesArgs) (CatalogEntityConnectionResolver, error)
	Graph(context.Context) (CatalogGraphResolver, error)
}

type CatalogEntitiesArgs struct {
	Query *string
	First *int32
	After *string
}

type CatalogGraphResolver interface {
	Nodes() []*CatalogEntityResolver
	Edges() []CatalogEntityRelationEdgeResolver
}

type CatalogEntityType string

type CatalogEntity interface {
	Node
	Type() CatalogEntityType
	Name() string
	Description() *string
	URL() string

	RelatedEntities(context.Context) (CatalogEntityRelatedEntityConnectionResolver, error)
}

type CatalogEntityResolver struct {
	CatalogEntity
}

func (r *CatalogEntityResolver) Typename() string {
	panic("X")
	switch r.CatalogEntity.(type) {
	case CatalogComponentResolver:
		return "CatalogComponent"
	default:
		panic("no __typename for CatalogEntity")
	}
}

func (r *CatalogEntityResolver) TypeName() string { return "CatalogComponent" }

func (r *CatalogEntityResolver) ToCatalogComponent() (CatalogComponentResolver, bool) {
	e, ok := r.CatalogEntity.(CatalogComponentResolver)
	return e, ok
}

type CatalogEntityRelationType string

type CatalogEntityRelationEdgeResolver interface {
	Type() CatalogEntityRelationType
	OutNode() *CatalogEntityResolver
	InNode() *CatalogEntityResolver
}

type CatalogEntityRelatedEntityConnectionResolver interface {
	Edges() []CatalogEntityRelatedEntityEdgeResolver
}

type CatalogEntityRelatedEntityEdgeResolver interface {
	Node() *CatalogEntityResolver
	Type() CatalogEntityRelationType
}

type CatalogEntityConnectionResolver interface {
	Nodes(context.Context) ([]*CatalogEntityResolver, error)
	TotalCount(context.Context) (int32, error)
	PageInfo(context.Context) (*graphqlutil.PageInfo, error)
}

type CatalogComponentResolver interface {
	CatalogEntity
	Kind() CatalogComponentKind

	Readme(context.Context) (FileResolver, error)
	SourceLocations(context.Context) ([]*GitTreeEntryResolver, error)
	Commits(context.Context, *graphqlutil.ConnectionArgs) (GitCommitConnectionResolver, error)
	Authors(context.Context) (*[]CatalogComponentAuthorEdgeResolver, error)
	Usage(context.Context, *CatalogComponentUsageArgs) (CatalogComponentUsageResolver, error)
	API(context.Context, *CatalogComponentAPIArgs) (CatalogComponentAPIResolver, error)
}

type CatalogComponentKind string

type CatalogComponentAuthorEdgeResolver interface {
	Component() CatalogComponentResolver
	Person() *PersonResolver
	AuthoredLineCount() int32
	AuthoredLineProportion() float64
	LastCommit(context.Context) (*GitCommitResolver, error)
}

type CatalogComponentUsageArgs struct {
	Query *string
}

type CatalogComponentUsageResolver interface {
	Locations(context.Context) (LocationConnectionResolver, error)
	People(context.Context) ([]CatalogComponentUsedByPersonEdgeResolver, error)
	Components(context.Context) ([]CatalogComponentUsedByComponentEdgeResolver, error)
}

type CatalogComponentUsedByPersonEdgeResolver interface {
	Node() *PersonResolver
	Locations(context.Context) (LocationConnectionResolver, error)
	AuthoredLineCount() int32
	LastCommit(context.Context) (*GitCommitResolver, error)
}

type CatalogComponentUsedByComponentEdgeResolver interface {
	Node() CatalogComponentResolver
	Locations(context.Context) (LocationConnectionResolver, error)
}

type CatalogComponentAPIArgs struct {
	Query *string
}

type CatalogComponentAPIResolver interface {
	Symbols(context.Context, *CatalogComponentAPISymbolsArgs) (*SymbolConnectionResolver, error)
	Schema(context.Context) (FileResolver, error)
}

type CatalogComponentAPISymbolsArgs struct {
	graphqlutil.ConnectionArgs
	Query *string
}
