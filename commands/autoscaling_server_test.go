package commands

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"regexp"
	"testing"

	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	sdkautoscaling "github.com/ionos-cloud/ionosctl/pkg/resources/autoscaling"
	"github.com/ionos-cloud/ionosctl/pkg/utils/clierror"
	ionoscloudautoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	testAutoscalingServerGet = sdkautoscaling.Server{
		Server: ionoscloudautoscaling.Server{
			Id:         &testAutoscalingServerVar,
			Properties: testAutoscalingServer.Properties,
			Metadata: &ionoscloudautoscaling.Metadata{
				State: &testAutoscalingStateTestVar,
			},
		},
	}
	testAutoscalingServer = sdkautoscaling.Server{
		Server: ionoscloudautoscaling.Server{
			Properties: &ionoscloudautoscaling.ServerProperties{
				DatacenterServer: &ionoscloudautoscaling.Resource{
					Id: &testAutoscalingServerVar,
				},
				Name: &testAutoscalingServerVar,
			},
		},
	}
	testAutoscalingServers = sdkautoscaling.Servers{
		ServerCollection: ionoscloudautoscaling.ServerCollection{
			Id:    &testAutoscalingServerVar,
			Items: &[]ionoscloudautoscaling.Server{testAutoscalingServerGet.Server, testAutoscalingServerGet.Server},
		},
	}
	testAutoscalingServerVar = "test-autoscaling-server"
	testAutoscalingServerErr = errors.New("autoscaling server test error occurred")
)

func TestPreRunAutoscalingGroupServerIds(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgServerId), testAutoscalingServerVar)
		err := PreRunAutoscalingGroupServerIds(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunAutoscalingGroupServerIdsRequiredFlagErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		err := PreRunAutoscalingGroupServerIds(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingServerList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingServerVar)
		rm.AutoscalingGroup.EXPECT().ListServers(testAutoscalingServerVar).Return(testAutoscalingServers, nil, nil)
		err := RunAutoscalingServerList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingServerListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingServerVar)
		rm.AutoscalingGroup.EXPECT().ListServers(testAutoscalingServerVar).Return(testAutoscalingServers, nil, testAutoscalingServerErr)
		err := RunAutoscalingServerList(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingServerGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgServerId), testAutoscalingServerVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().GetServer(testAutoscalingServerVar, testAutoscalingServerVar).Return(&testAutoscalingServerGet, nil, nil)
		err := RunAutoscalingServerGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingServerGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingServerVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgServerId), testAutoscalingServerVar)
		rm.AutoscalingGroup.EXPECT().GetServer(testAutoscalingServerVar, testAutoscalingServerVar).Return(&testAutoscalingServerGet, nil, testAutoscalingServerErr)
		err := RunAutoscalingServerGet(cfg)
		assert.Error(t, err)
	})
}

func TestGetAutoscalingServersCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling server", config.ArgCols), []string{"Name"})
	getAutoscalingServerCols(core.GetGlobalFlagName("autoscaling server", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetAutoscalingServersColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling server", config.ArgCols), []string{"Unknown"})
	getAutoscalingServerCols(core.GetGlobalFlagName("autoscaling server", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestGetAutoscalingServersIds(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	err := os.Setenv(ionoscloudautoscaling.IonosUsernameEnvVar, "user")
	assert.NoError(t, err)
	err = os.Setenv(ionoscloudautoscaling.IonosPasswordEnvVar, "pass")
	assert.NoError(t, err)
	viper.Set(config.ArgServerUrl, config.DefaultApiURL)
	getAutoscalingServersIds(w, testAutoscalingServerVar)
	err = w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`401 Unauthorized`)
	assert.True(t, re.Match(b.Bytes()))
}
