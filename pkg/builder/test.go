package builder

import (
	"context"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/resources"
	mocks "github.com/ionos-cloud/ionosctl/pkg/resources/mocks"
	"github.com/ionos-cloud/ionosctl/pkg/utils"
	"github.com/spf13/viper"
)

type CmdRunnerTest func(c *CommandConfig, mocks *ResourcesMocks)

type ResourcesMocks struct {
	Client     *mocks.MockClientService
	Datacenter *mocks.MockDatacentersService
}

func CmdConfigTest(t *testing.T, writer io.Writer, runner CmdRunnerTest) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	printer := utils.NewPrinterRegistry(writer, writer)
	prt := printer[viper.GetString(config.ArgOutput)]

	tm := &ResourcesMocks{
		Client:     mocks.NewMockClientService(ctrl),
		Datacenter: mocks.NewMockDatacentersService(ctrl),
	}

	cmdConfig := &CommandConfig{
		Name:         "test",
		Printer:      prt,
		Context:      context.TODO(),
		initServices: func(c *CommandConfig) error { return nil },
		DataCenters: func() resources.DatacentersService {
			return tm.Datacenter
		},
	}

	runner(cmdConfig, tm)
}
