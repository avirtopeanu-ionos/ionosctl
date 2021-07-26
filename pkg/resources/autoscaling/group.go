package autoscaling

import (
	"context"

	ionoscloudautoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
)

type Group struct {
	ionoscloudautoscaling.Group
}

type GroupProperties struct {
	ionoscloudautoscaling.GroupProperties
}

type Groups struct {
	ionoscloudautoscaling.GroupCollection
}

type Server struct {
	ionoscloudautoscaling.Server
}

type ServerProperties struct {
	ionoscloudautoscaling.ServerProperties
}

type Servers struct {
	ionoscloudautoscaling.ServerCollection
}

type Action struct {
	ionoscloudautoscaling.Action
}

type ActionProperties struct {
	ionoscloudautoscaling.ActionProperties
}

type Actions struct {
	ionoscloudautoscaling.ActionCollection
}

// GroupsService is a wrapper around ionoscloudautoscaling.Group
type GroupsService interface {
	List() (Groups, *Response, error)
	Get(groupId string) (*Group, *Response, error)
	Create(input Group) (*Group, *Response, error)
	Update(groupId string, input Group) (*Group, *Response, error)
	Delete(GroupId string) (*Response, error)
	ListServers(groupId string) (Servers, *Response, error)
	GetServer(groupId, serverId string) (*Server, *Response, error)
	ListActions(groupId string) (Actions, *Response, error)
	GetAction(groupId, actionId string) (*Action, *Response, error)
}

type groupsService struct {
	client  *Client
	context context.Context
}

var _ GroupsService = &groupsService{}

func NewGroupService(client *Client, ctx context.Context) GroupsService {
	return &groupsService{
		client:  client,
		context: ctx,
	}
}

func (gs *groupsService) List() (Groups, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsGet(gs.context)
	groups, res, err := gs.client.GroupsApi.AutoscalingGroupsGetExecute(req)
	return Groups{groups}, &Response{*res}, err
}

func (gs *groupsService) Get(groupId string) (*Group, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsFindById(gs.context, groupId)
	group, res, err := gs.client.GroupsApi.AutoscalingGroupsFindByIdExecute(req)
	return &Group{group}, &Response{*res}, err
}

func (gs *groupsService) Create(input Group) (*Group, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsPost(gs.context).Group(input.Group)
	group, res, err := gs.client.GroupsApi.AutoscalingGroupsPostExecute(req)
	return &Group{group}, &Response{*res}, err
}

func (gs *groupsService) Update(groupId string, input Group) (*Group, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsPut(gs.context, groupId).Group(input.Group)
	group, res, err := gs.client.GroupsApi.AutoscalingGroupsPutExecute(req)
	return &Group{group}, &Response{*res}, err
}

func (gs *groupsService) Delete(GroupId string) (*Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsDelete(context.Background(), GroupId)
	res, err := gs.client.GroupsApi.AutoscalingGroupsDeleteExecute(req)
	return &Response{*res}, err
}

func (gs *groupsService) ListServers(groupId string) (Servers, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsServersGet(gs.context, groupId)
	servers, res, err := gs.client.GroupsApi.AutoscalingGroupsServersGetExecute(req)
	return Servers{servers}, &Response{*res}, err
}

func (gs *groupsService) GetServer(groupId, serverId string) (*Server, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsServersFindById(gs.context, groupId, serverId)
	server, res, err := gs.client.GroupsApi.AutoscalingGroupsServersFindByIdExecute(req)
	return &Server{server}, &Response{*res}, err
}

func (gs *groupsService) ListActions(groupId string) (Actions, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsActionsGet(gs.context, groupId)
	actions, res, err := gs.client.GroupsApi.AutoscalingGroupsActionsGetExecute(req)
	return Actions{actions}, &Response{*res}, err
}

func (gs *groupsService) GetAction(groupId, serverId string) (*Action, *Response, error) {
	req := gs.client.GroupsApi.AutoscalingGroupsActionsFindById(gs.context, groupId, serverId)
	action, res, err := gs.client.GroupsApi.AutoscalingGroupsActionsFindByIdExecute(req)
	return &Action{action}, &Response{*res}, err
}
