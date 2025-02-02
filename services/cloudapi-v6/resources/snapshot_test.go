package resources

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSnapshotResourceVar = "test-snapshot-resource"
)

func TestNewSnapshotService(t *testing.T) {
	ctx := context.Background()
	t.Run("list_snapshots_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, _, err := snapshotSvc.List(ListQueryParams{})
		assert.Error(t, err)
	})
	t.Run("list_snapshots_filters_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, _, err := snapshotSvc.List(testListQueryParam)
		assert.Error(t, err)
	})
	t.Run("get_snapshot_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, _, err := snapshotSvc.Get(testSnapshotResourceVar, QueryParams{})
		assert.Error(t, err)
	})
	t.Run("create_snapshot_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, _, err := snapshotSvc.Create(
			testSnapshotResourceVar,
			testSnapshotResourceVar,
			testSnapshotResourceVar,
			testSnapshotResourceVar,
			testSnapshotResourceVar,
			false,
			QueryParams{},
		)
		assert.Error(t, err)
	})
	t.Run("update_snapshot_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, _, err := snapshotSvc.Update(testSnapshotResourceVar, SnapshotProperties{}, QueryParams{})
		assert.Error(t, err)
	})
	t.Run("restore_snapshot_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, err := snapshotSvc.Restore(testSnapshotResourceVar, testSnapshotResourceVar, testSnapshotResourceVar, QueryParams{})
		assert.Error(t, err)
	})
	t.Run("delete_snapshot_error", func(t *testing.T) {
		svc := getTestClient(t)
		snapshotSvc := NewSnapshotService(svc, ctx)
		_, err := snapshotSvc.Delete(testSnapshotResourceVar, QueryParams{})
		assert.Error(t, err)
	})
}
