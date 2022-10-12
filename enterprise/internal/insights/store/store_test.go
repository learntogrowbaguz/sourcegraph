package store

import (
	"context"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hexops/autogold"
	"github.com/keegancsmith/sqlf"
	"github.com/sourcegraph/log/logtest"

	edb "github.com/sourcegraph/sourcegraph/enterprise/internal/database"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/insights/types"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/database/dbtest"
	"github.com/sourcegraph/sourcegraph/internal/timeutil"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

func TestSeriesPoints(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsDB := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))

	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	store := NewWithClock(insightsDB, permStore, clock)

	// Confirm we get no results initially.
	points, err := store.SeriesPoints(ctx, SeriesPointsOpts{})
	if err != nil {
		t.Fatal(err)
	}
	autogold.Want("SeriesPoints", []SeriesPoint{}).Equal(t, points)

	// Insert some fake data.
	_, err = insightsDB.ExecContext(context.Background(), `
INSERT INTO repo_names(name) VALUES ('github.com/gorilla/mux-original');
INSERT INTO repo_names(name) VALUES ('github.com/gorilla/mux-renamed');
SELECT setseed(0.5);
INSERT INTO series_points(
    time,
	series_id,
    value,
    repo_id,
    repo_name_id,
    original_repo_name_id)
SELECT time,
    'somehash',
    random()*80 - 40,
    2,
    (SELECT id FROM repo_names WHERE name = 'github.com/gorilla/mux-renamed'),
    (SELECT id FROM repo_names WHERE name = 'github.com/gorilla/mux-original')
	FROM GENERATE_SERIES(CURRENT_TIMESTAMP::date - INTERVAL '30 weeks', CURRENT_TIMESTAMP::date, '2 weeks') AS time;
`)
	if err != nil {
		t.Fatal(err)
	}

	time := func(s string) *time.Time {
		v, err := time.Parse(time.RFC3339, s)
		if err != nil {
			t.Fatal(err)
		}
		return &v
	}

	t.Run("all data points", func(t *testing.T) {
		// Confirm we get all data points.
		points, err = store.SeriesPoints(ctx, SeriesPointsOpts{})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(points)
		autogold.Want("SeriesPoints(2).len", int(16)).Equal(t, len(points))
	})

	t.Run("subset of data", func(t *testing.T) {
		// Confirm we can get a subset of data points.
		points, err = store.SeriesPoints(ctx, SeriesPointsOpts{
			From: time("2020-03-01T00:00:00Z"),
			To:   time("2020-06-01T00:00:00Z"),
		})
		if err != nil {
			t.Fatal(err)
		}
		autogold.Want("SeriesPoints(3).len", int(0)).Equal(t, len(points))
	})

	t.Run("latest 3 points", func(t *testing.T) {
		// Confirm we can get a subset of data points.
		points, err = store.SeriesPoints(ctx, SeriesPointsOpts{
			Limit: 3,
		})
		if err != nil {
			t.Fatal(err)
		}
		autogold.Want("SeriesPoints(4).len", int(3)).Equal(t, len(points))
	})

	t.Run("include list", func(t *testing.T) {
		points, err = store.SeriesPoints(ctx, SeriesPointsOpts{Included: []api.RepoID{2}})
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(16, len(points)); diff != "" {
			t.Errorf("unexpected results from include list: %v", diff)
		}
	})
	t.Run("exclude list", func(t *testing.T) {
		points, err = store.SeriesPoints(ctx, SeriesPointsOpts{Excluded: []api.RepoID{2}})
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(0, len(points)); diff != "" {
			t.Errorf("unexpected results from include list: %v", diff)
		}
	})
}

func TestCountData(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsDB := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))
	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	store := NewWithClock(insightsDB, permStore, clock)

	timeValue := func(s string) time.Time {
		v, err := time.Parse(time.RFC3339, s)
		if err != nil {
			t.Fatal(err)
		}
		return v
	}
	timePtr := func(s string) *time.Time {
		t := timeValue(s)
		return &t
	}
	optionalString := func(v string) *string { return &v }
	optionalRepoID := func(v api.RepoID) *api.RepoID { return &v }

	// Record some duplicate data points.
	records := []RecordSeriesPointArgs{
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: timeValue("2020-03-01T00:00:00Z"), Value: 1.1},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: RecordMode,
		},
		{
			SeriesID:    "two",
			Point:       SeriesPoint{Time: timeValue("2020-03-02T00:00:00Z"), Value: 2.2},
			PersistMode: RecordMode,
		},
		{
			SeriesID:    "two",
			Point:       SeriesPoint{Time: timeValue("2020-03-02T00:01:00Z"), Value: 2.2},
			PersistMode: RecordMode,
		},
		{
			SeriesID:    "three",
			Point:       SeriesPoint{Time: timeValue("2020-03-03T00:00:00Z"), Value: 3.3},
			PersistMode: RecordMode,
		},
		{
			SeriesID:    "three",
			Point:       SeriesPoint{Time: timeValue("2020-03-03T00:01:00Z"), Value: 3.3},
			PersistMode: RecordMode,
		},
	}
	if err := store.RecordSeriesPoints(ctx, records); err != nil {
		t.Fatal(err)
	}

	// How many data points on 02-29?
	numDataPoints, err := store.CountData(ctx, CountDataOpts{
		From: timePtr("2020-02-29T00:00:00Z"),
		To:   timePtr("2020-02-29T23:59:59Z"),
	})
	if err != nil {
		t.Fatal(err)
	}
	autogold.Want("first", int(0)).Equal(t, numDataPoints)

	// How many data points on 03-01?
	numDataPoints, err = store.CountData(ctx, CountDataOpts{
		From: timePtr("2020-03-01T00:00:00Z"),
		To:   timePtr("2020-03-01T23:59:59Z"),
	})
	if err != nil {
		t.Fatal(err)
	}
	autogold.Want("second", int(1)).Equal(t, numDataPoints)

	// How many data points from 03-01 to 03-04?
	numDataPoints, err = store.CountData(ctx, CountDataOpts{
		From: timePtr("2020-03-01T00:00:00Z"),
		To:   timePtr("2020-03-04T23:59:59Z"),
	})
	if err != nil {
		t.Fatal(err)
	}
	autogold.Want("third", int(5)).Equal(t, numDataPoints)
}

func TestRecordSeriesPoints(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsDB := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))
	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	store := NewWithClock(insightsDB, permStore, clock)

	// First test it does not error with no records.
	if err := store.RecordSeriesPoints(ctx, []RecordSeriesPointArgs{}); err != nil {
		t.Fatal(err)
	}

	optionalString := func(v string) *string { return &v }
	optionalRepoID := func(v api.RepoID) *api.RepoID { return &v }

	current := time.Date(2021, time.September, 10, 10, 0, 0, 0, time.UTC)

	records := []RecordSeriesPointArgs{
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: current, Value: 1.1},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: RecordMode,
		},
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: current.Add(-time.Hour * 24 * 14), Value: 2.2},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: RecordMode,
		},
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: current.Add(-time.Hour * 24 * 28), Value: 3.3},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: SnapshotMode,
		},
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: current.Add(-time.Hour * 24 * 42), Value: 3.3},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: SnapshotMode,
		},
	}
	if err := store.RecordSeriesPoints(ctx, records); err != nil {
		t.Fatal(err)
	}

	want := []SeriesPoint{
		{
			SeriesID: "one",
			Time:     current.Add(-time.Hour * 24 * 42),
			Value:    3.3,
		},
		{
			SeriesID: "one",
			Time:     current.Add(-time.Hour * 24 * 28),
			Value:    3.3,
		},
		{
			SeriesID: "one",
			Time:     current.Add(-time.Hour * 24 * 14),
			Value:    2.2,
		},
		{
			SeriesID: "one",
			Time:     current,
			Value:    1.1,
		},
	}

	// Confirm we get the expected data back.
	points, err := store.SeriesPoints(ctx, SeriesPointsOpts{})
	if err != nil {
		t.Fatal(err)
	}
	autogold.Want("len(points)", int(4)).Equal(t, len(points))
	if diff := cmp.Diff(4, len(points)); diff != "" {
		t.Errorf("len(points): %v", diff)
	}
	if diff := cmp.Diff(want[0], points[0]); diff != "" {
		t.Errorf("points[0].String(): %v", diff)
	}
	if diff := cmp.Diff(want[1], points[1]); diff != "" {
		t.Errorf("points[1].String(): %v", diff)
	}
	if diff := cmp.Diff(want[2], points[2]); diff != "" {
		t.Errorf("points[2].String(): %v", diff)
	}
	if diff := cmp.Diff(want[3], points[3]); diff != "" {
		t.Errorf("points[3].String(): %v", diff)
	}
}

func TestRecordSeriesPointsSnapshotOnly(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsDB := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))
	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	store := NewWithClock(insightsDB, permStore, clock)

	optionalString := func(v string) *string { return &v }
	optionalRepoID := func(v api.RepoID) *api.RepoID { return &v }

	current := time.Date(2021, time.September, 10, 10, 0, 0, 0, time.UTC)

	records := []RecordSeriesPointArgs{
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: current, Value: 1.1},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: SnapshotMode,
		},
	}
	if err := store.RecordSeriesPoints(ctx, records); err != nil {
		t.Fatal(err)
	}

	// check snapshots table has a row
	row := store.QueryRow(ctx, sqlf.Sprintf("select count(*) from %s", sqlf.Sprintf(snapshotsTable)))
	if row.Err() != nil {
		t.Fatal(row.Err())
	}

	want := 1
	var got int
	err := row.Scan(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected count from snapshots table (want/got): %v", diff)
	}

	// check recordings table has no rows
	row = store.QueryRow(ctx, sqlf.Sprintf("select count(*) from %s", sqlf.Sprintf(recordingTable)))
	if row.Err() != nil {
		t.Fatal(row.Err())
	}

	want = 0
	err = row.Scan(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected count from recordings table (want/got): %v", diff)
	}
}

func TestRecordSeriesPointsRecordingOnly(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsDB := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))
	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	store := NewWithClock(insightsDB, permStore, clock)

	optionalString := func(v string) *string { return &v }
	optionalRepoID := func(v api.RepoID) *api.RepoID { return &v }

	current := time.Date(2021, time.September, 10, 10, 0, 0, 0, time.UTC)

	records := []RecordSeriesPointArgs{
		{
			SeriesID:    "one",
			Point:       SeriesPoint{Time: current, Value: 1.1},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: RecordMode,
		},
	}
	if err := store.RecordSeriesPoints(ctx, records); err != nil {
		t.Fatal(err)
	}

	// check snapshots table has a row
	row := store.QueryRow(ctx, sqlf.Sprintf("select count(*) from %s", sqlf.Sprintf(snapshotsTable)))
	if row.Err() != nil {
		t.Fatal(row.Err())
	}

	want := 0
	var got int
	err := row.Scan(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected count from snapshots table (want/got): %v", diff)
	}

	// check recordings table has no rows
	row = store.QueryRow(ctx, sqlf.Sprintf("select count(*) from %s", sqlf.Sprintf(recordingTable)))
	if row.Err() != nil {
		t.Fatal(row.Err())
	}

	want = 1
	err = row.Scan(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected count from recordings table (want/got): %v", diff)
	}
}

func TestDeleteSnapshots(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsDB := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))
	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	store := NewWithClock(insightsDB, permStore, clock)

	optionalString := func(v string) *string { return &v }
	optionalRepoID := func(v api.RepoID) *api.RepoID { return &v }

	current := time.Date(2021, time.September, 10, 10, 0, 0, 0, time.UTC)

	seriesID := "one"
	records := []RecordSeriesPointArgs{
		{
			SeriesID:    seriesID,
			Point:       SeriesPoint{Time: current, Value: 1.1},
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: SnapshotMode,
		},
		{
			SeriesID:    seriesID,
			Point:       SeriesPoint{Time: current.Add(time.Hour), Value: 1.1}, // offsetting the time by an hour so that the point is not deduplicated
			RepoName:    optionalString("repo1"),
			RepoID:      optionalRepoID(3),
			PersistMode: RecordMode,
		},
	}
	if err := store.RecordSeriesPoints(ctx, records); err != nil {
		t.Fatal(err)
	}

	// first check that we have one recording and one snapshot
	points, err := store.SeriesPoints(ctx, SeriesPointsOpts{SeriesID: &seriesID})
	if err != nil {
		t.Fatal(err)
	}
	got := len(points)
	want := 2
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected count of series points prior to deleting snapshots (want/got): %v", diff)
	}
	err = store.DeleteSnapshots(ctx, &types.InsightSeries{SeriesID: seriesID})
	if err != nil {
		t.Fatal(err)
	}
	// now verify that the remaining point is the recording
	points, err = store.SeriesPoints(ctx, SeriesPointsOpts{SeriesID: &seriesID})
	if err != nil {
		t.Fatal(err)
	}
	got = len(points)
	want = 1
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected count of series points after deleting snapshots (want/got): %v", diff)
	}
	autogold.Equal(t, points, autogold.ExportedOnly())
}

func TestValues(t *testing.T) {
	ids := []api.RepoID{1, 2, 3, 4, 5, 6}
	got := values(ids)
	want := "VALUES (1),(2),(3),(4),(5),(6)"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected values string: %v", diff)
	}
}

func TestDelete(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	now := time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC)

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsdb := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))

	repoName := "reallygreatrepo"
	repoId := api.RepoID(5)

	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	timeseriesStore := NewWithClock(insightsdb, permStore, clock)

	err := timeseriesStore.RecordSeriesPoints(ctx, []RecordSeriesPointArgs{
		{
			SeriesID: "series1",
			Point: SeriesPoint{
				SeriesID: "series1",
				Time:     now,
				Value:    50,
			},
			RepoName:    &repoName,
			RepoID:      &repoId,
			PersistMode: RecordMode,
		},
		{
			SeriesID: "series1",
			Point: SeriesPoint{
				SeriesID: "series1",
				Time:     now,
				Value:    50,
			},
			RepoName:    &repoName,
			RepoID:      &repoId,
			PersistMode: SnapshotMode,
		},
		{
			SeriesID: "series2",
			Point: SeriesPoint{
				SeriesID: "series2",
				Time:     now,
				Value:    25,
			},
			RepoName:    &repoName,
			RepoID:      &repoId,
			PersistMode: RecordMode,
		},
		{
			SeriesID: "series2",
			Point: SeriesPoint{
				SeriesID: "series2",
				Time:     now,
				Value:    25,
			},
			RepoName:    &repoName,
			RepoID:      &repoId,
			PersistMode: SnapshotMode,
		},
	})
	if err != nil {
		t.Error(err)
	}

	err = timeseriesStore.SetInsightSeriesRecordingTimes(ctx, []types.InsightSeriesRecordingTimes{
		{
			"series1",
			[]time.Time{now},
		},
		{
			"series2",
			[]time.Time{now},
		},
	})
	if err != nil {
		t.Error(err)
	}

	err = timeseriesStore.Delete(ctx, "series1")
	if err != nil {
		t.Fatal(err)
	}

	getCountForSeries := func(ctx context.Context, timeseriesStore *Store, mode PersistMode, seriesId string) int {
		table, err := getTableForPersistMode(mode)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(table)
		q := sqlf.Sprintf("select count(*) from %s where series_id = %s;", sqlf.Sprintf(table), seriesId)
		row := timeseriesStore.QueryRow(ctx, q)
		val, err := basestore.ScanInt(row)
		if err != nil {
			t.Fatal(err)
		}
		return val
	}

	getTimesCountforSeries := func(ctx context.Context, timeseriesStore *Store, seriesId string) int {
		q := sqlf.Sprintf("select count(*) from insight_series_recording_times where series_id = %s;", seriesId)
		row := timeseriesStore.QueryRow(ctx, q)
		val, err := basestore.ScanInt(row)
		if err != nil {
			t.Fatal(err)
		}
		return val
	}

	if getCountForSeries(ctx, timeseriesStore, RecordMode, "series1") != 0 {
		t.Errorf("expected 0 count for series1 in record table")
	}
	if getCountForSeries(ctx, timeseriesStore, SnapshotMode, "series1") != 0 {
		t.Errorf("expected 0 count for series1 in snapshot table")
	}
	if getTimesCountforSeries(ctx, timeseriesStore, "series1") != 0 {
		t.Errorf("expected 0 recording times to remain for series1")
	}

	if getCountForSeries(ctx, timeseriesStore, RecordMode, "series2") != 1 {
		t.Errorf("expected 1 count for series2 in record table")
	}
	if getCountForSeries(ctx, timeseriesStore, SnapshotMode, "series2") != 1 {
		t.Errorf("expected 1 count for series2 in snapshot table")
	}
	if getTimesCountforSeries(ctx, timeseriesStore, "series2") != 1 {
		t.Errorf("expected 1 recording times to remain for series2")
	}
}

func getTableForPersistMode(mode PersistMode) (string, error) {
	switch mode {
	case RecordMode:
		return recordingTable, nil
	case SnapshotMode:
		return snapshotsTable, nil
	default:
		return "", errors.Newf("unsupported insights series point persist mode: %v", mode)
	}
}

func TestInsightSeriesRecordingTimes(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	now := time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC)

	logger := logtest.Scoped(t)
	ctx := context.Background()
	clock := timeutil.Now
	insightsdb := edb.NewInsightsDB(dbtest.NewInsightsDB(logger, t))

	postgres := database.NewDB(logger, dbtest.NewDB(logger, t))
	permStore := NewInsightPermissionStore(postgres)
	timeseriesStore := NewWithClock(insightsdb, permStore, clock)

	series1Times := types.InsightSeriesRecordingTimes{
		"series1",
		[]time.Time{now, now.AddDate(0, 1, 0)},
	}

	err := timeseriesStore.SetInsightSeriesRecordingTimes(ctx, []types.InsightSeriesRecordingTimes{
		series1Times,
		{
			"series2",
			[]time.Time{now, now.AddDate(0, 1, 0)},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := timeseriesStore.GetInsightSeriesRecordingTimes(ctx, "series1")
	if err != nil {
		t.Fatal(err)
	}

	stringifyTimes := func(times []time.Time) string {
		s := []string{}
		for _, t := range times {
			s = append(s, t.String())
		}
		sort.Strings(s)
		return strings.Join(s, " ")
	}
	if got.SeriesID != series1Times.SeriesID || stringifyTimes(got.RecordingTimes) != stringifyTimes(series1Times.RecordingTimes) {
		t.Errorf("got %v, want %v", got, series1Times)
	}

	err = timeseriesStore.DeleteInsightSeriesRecordingTimes(ctx, types.InsightSeriesRecordingTimes{"series2", []time.Time{now}})
	if err != nil {
		t.Fatal(err)
	}

	got, err = timeseriesStore.GetInsightSeriesRecordingTimes(ctx, "series2")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.RecordingTimes) != 1 {
		t.Fatalf("got %d recording times, expected 1", len(got.RecordingTimes))
	}
	if got.RecordingTimes[0] != now.AddDate(0, 1, 0) {
		t.Errorf("unexpected date, got %v want %v", got.RecordingTimes[0], now.AddDate(0, 1, 0))
	}
}
