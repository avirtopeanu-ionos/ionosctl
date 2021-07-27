package autoscaling

import (
	"context"

	ionoscloudautoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
)

type Template struct {
	ionoscloudautoscaling.Template
}

type TemplateVolume struct {
	ionoscloudautoscaling.TemplateVolume
}

type TemplateNic struct {
	ionoscloudautoscaling.TemplateNic
}

type TemplateProperties struct {
	ionoscloudautoscaling.TemplateProperties
}

type Templates struct {
	ionoscloudautoscaling.TemplateCollection
}

type Response struct {
	ionoscloudautoscaling.APIResponse
}

// TemplatesService is a wrapper around ionoscloudautoscaling.Template
type TemplatesService interface {
	List() (Templates, *Response, error)
	Get(TemplateId string) (*Template, *Response, error)
	Create(input Template) (*Template, *Response, error)
	Delete(TemplateId string) (*Response, error)
}

type templatesService struct {
	client  *Client
	context context.Context
}

var _ TemplatesService = &templatesService{}

func NewTemplateService(client *Client, ctx context.Context) TemplatesService {
	return &templatesService{
		client:  client,
		context: ctx,
	}
}

func (ts *templatesService) List() (Templates, *Response, error) {
	req := ts.client.TemplatesApi.AutoscalingTemplatesGet(ts.context)
	dcs, res, err := ts.client.TemplatesApi.AutoscalingTemplatesGetExecute(req)
	return Templates{dcs}, &Response{*res}, err
}

func (ts *templatesService) Get(TemplateId string) (*Template, *Response, error) {
	req := ts.client.TemplatesApi.AutoscalingTemplatesFindById(ts.context, TemplateId)
	template, res, err := ts.client.TemplatesApi.AutoscalingTemplatesFindByIdExecute(req)
	return &Template{template}, &Response{*res}, err
}

func (ts *templatesService) Create(input Template) (*Template, *Response, error) {
	req := ts.client.TemplatesApi.AutoscalingTemplatesPost(ts.context).Template(input.Template)
	template, res, err := ts.client.TemplatesApi.AutoscalingTemplatesPostExecute(req)
	return &Template{template}, &Response{*res}, err
}

func (ts *templatesService) Delete(TemplateId string) (*Response, error) {
	req := ts.client.TemplatesApi.AutoscalingTemplatesDelete(context.Background(), TemplateId)
	res, err := ts.client.TemplatesApi.AutoscalingTemplatesDeleteExecute(req)
	return &Response{*res}, err
}
