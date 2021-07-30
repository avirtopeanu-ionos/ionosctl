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
	testAutoscalingActionGet = sdkautoscaling.Action{
		Action: ionoscloudautoscaling.Action{
			Id:         &testAutoscalingActionVar,
			Properties: testAutoscalingAction.Properties,
			Metadata: &ionoscloudautoscaling.Metadata{
				State: &testAutoscalingStateTestVar,
			},
		},
	}
	testAutoscalingAction = sdkautoscaling.Action{
		Action: ionoscloudautoscaling.Action{
			Properties: &ionoscloudautoscaling.ActionProperties{
				ActionStatus:       (*ionoscloudautoscaling.ActionStatus)(&testAutoscalingActionVar),
				ActionType:         (*ionoscloudautoscaling.ActionType)(&testAutoscalingActionVar),
				TargetReplicaCount: &testAutoscalingActionIntVar,
			},
		},
	}
	testAutoscalingActions = sdkautoscaling.Actions{
		ActionCollection: ionoscloudautoscaling.ActionCollection{
			Id:    &testAutoscalingActionVar,
			Items: &[]ionoscloudautoscaling.Action{testAutoscalingActionGet.Action, testAutoscalingActionGet.Action},
		},
	}
	testAutoscalingActionVar    = "test-autoscaling-action"
	testAutoscalingActionIntVar = int64(1)
	testAutoscalingActionErr    = errors.New("autoscaling action test error occurred")
)

func TestPreRunAutoscalingGroupActionIds(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingActionVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgActionId), testAutoscalingActionVar)
		err := PreRunAutoscalingGroupActionIds(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunAutoscalingGroupActionIdsRequiredFlagErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		err := PreRunAutoscalingGroupActionIds(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingActionList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingActionVar)
		rm.AutoscalingGroup.EXPECT().ListActions(testAutoscalingActionVar).Return(testAutoscalingActions, nil, nil)
		err := RunAutoscalingActionList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingActionListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingActionVar)
		rm.AutoscalingGroup.EXPECT().ListActions(testAutoscalingActionVar).Return(testAutoscalingActions, nil, testAutoscalingActionErr)
		err := RunAutoscalingActionList(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingActionGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingActionVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgActionId), testAutoscalingActionVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().GetAction(testAutoscalingActionVar, testAutoscalingActionVar).Return(&testAutoscalingActionGet, nil, nil)
		err := RunAutoscalingActionGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingActionGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingActionVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgActionId), testAutoscalingActionVar)
		rm.AutoscalingGroup.EXPECT().GetAction(testAutoscalingActionVar, testAutoscalingActionVar).Return(&testAutoscalingActionGet, nil, testAutoscalingActionErr)
		err := RunAutoscalingActionGet(cfg)
		assert.Error(t, err)
	})
}

func TestGetAutoscalingActionsCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling action", config.ArgCols), []string{"ActionStatus"})
	getAutoscalingActionCols(core.GetGlobalFlagName("autoscaling action", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetAutoscalingActionsColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling action", config.ArgCols), []string{"Unknown"})
	getAutoscalingActionCols(core.GetGlobalFlagName("autoscaling action", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestGetAutoscalingActionsIds(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	err := os.Setenv(ionoscloudautoscaling.IonosUsernameEnvVar, "user")
	assert.NoError(t, err)
	err = os.Setenv(ionoscloudautoscaling.IonosPasswordEnvVar, "pass")
	assert.NoError(t, err)
	viper.Set(config.ArgServerUrl, config.DefaultApiURL)
	getAutoscalingActionsIds(w, testAutoscalingActionVar)
	err = w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`401 Unauthorized`)
	assert.True(t, re.Match(b.Bytes()))
}
