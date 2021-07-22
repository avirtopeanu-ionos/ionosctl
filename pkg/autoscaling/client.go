package autoscaling

import (
	"errors"

	ionoscloudautoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
)

type Client struct {
	ionoscloudautoscaling.APIClient
}

type ClientConfig struct {
	ionoscloudautoscaling.Configuration
}

// ClientService is a wrapper around ionoscloud.APIClient
type ClientService interface {
	Get() *Client
	GetConfig() *ClientConfig
}

type clientService struct {
	client *ionoscloudautoscaling.APIClient
}

var _ ClientService = &clientService{}

func NewClientService(name, pwd, token, url string) (ClientService, error) {
	if url == "" {
		return nil, errors.New("server-url incorrect")
	}
	if token == "" && (name == "" || pwd == "") {
		return nil, errors.New("username, password or token incorrect")
	}
	clientConfig := &ionoscloudautoscaling.Configuration{
		Username: name,
		Password: pwd,
		Token:    token,
		Servers: ionoscloudautoscaling.ServerConfigurations{
			ionoscloudautoscaling.ServerConfiguration{
				URL: url,
			},
		},
	}
	return &clientService{
		client: ionoscloudautoscaling.NewAPIClient(clientConfig),
	}, nil
}

func (c clientService) Get() *Client {
	return &Client{
		APIClient: *c.client,
	}
}

func (c clientService) GetConfig() *ClientConfig {
	return &ClientConfig{
		Configuration: *c.client.GetConfig(),
	}
}
