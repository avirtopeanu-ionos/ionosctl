package autoscaling

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testGroupResourceVar = "test-group-resource"

func TestNewGroupService(t *testing.T) {
	ctx := context.Background()
	t.Run("list_groups_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.List()
		assert.Error(t, err)
	})
	t.Run("get_group_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.Get(testGroupResourceVar)
		assert.Error(t, err)
	})
	t.Run("create_group_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.Create(Group{})
		assert.Error(t, err)
	})
	t.Run("update_group_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.Update(testGroupResourceVar, Group{})
		assert.Error(t, err)
	})
	t.Run("delete_group_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, err := groupSvc.Delete(testGroupResourceVar)
		assert.Error(t, err)
	})
	t.Run("listservers_groups_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.ListServers(testGroupResourceVar)
		assert.Error(t, err)
	})
	t.Run("getserver_groups_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.GetServer(testGroupResourceVar, testGroupResourceVar)
		assert.Error(t, err)
	})
	t.Run("listactions_groups_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.ListActions(testGroupResourceVar)
		assert.Error(t, err)
	})
	t.Run("getaction_groups_error", func(t *testing.T) {
		svc := getTestClient(t)
		groupSvc := NewGroupService(svc.Get(), ctx)
		_, _, err := groupSvc.GetAction(testGroupResourceVar, testGroupResourceVar)
		assert.Error(t, err)
	})
}
