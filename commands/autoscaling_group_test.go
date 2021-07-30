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
	testAutoscalingGroupGet = sdkautoscaling.Group{
		Group: ionoscloudautoscaling.Group{
			Id:         &testAutoscalingGroupVar,
			Properties: &testAutoscalingGroupProperties.GroupProperties,
			Metadata: &ionoscloudautoscaling.Metadata{
				State: &testAutoscalingStateTestVar,
			},
		},
	}
	testAutoscalingGroup = sdkautoscaling.Group{
		Group: ionoscloudautoscaling.Group{
			Properties: &testAutoscalingGroupProperties.GroupProperties,
		},
	}
	testAutoscalingGroupProperties = sdkautoscaling.GroupProperties{
		GroupProperties: ionoscloudautoscaling.GroupProperties{
			Datacenter: &ionoscloudautoscaling.GroupPropertiesDatacenter{
				Id: &testAutoscalingGroupVar,
			},
			MaxReplicaCount: &testAutoscalingGroupIntVar,
			MinReplicaCount: &testAutoscalingGroupIntVar,
			Name:            &testAutoscalingGroupVar,
			Policy: &ionoscloudautoscaling.GroupPolicy{
				Metric: (*ionoscloudautoscaling.Metric)(&testAutoscalingGroupVar),
				Range:  &testAutoscalingGroupVar,
				ScaleInAction: &ionoscloudautoscaling.GroupPolicyAction{
					Amount:         &testAutoscalingGroupFloatVar,
					AmountType:     (*ionoscloudautoscaling.ActionAmount)(&testAutoscalingGroupVar),
					CooldownPeriod: &testAutoscalingGroupVar,
				},
				ScaleInThreshold: &testAutoscalingGroupFloatVar,
				ScaleOutAction: &ionoscloudautoscaling.GroupPolicyAction{
					Amount:         &testAutoscalingGroupFloatVar,
					AmountType:     (*ionoscloudautoscaling.ActionAmount)(&testAutoscalingGroupVar),
					CooldownPeriod: &testAutoscalingGroupVar,
				},
				ScaleOutThreshold: &testAutoscalingGroupFloatVar,
				Unit:              (*ionoscloudautoscaling.QueryUnit)(&testAutoscalingGroupVar),
			},
			TargetReplicaCount: &testAutoscalingGroupIntVar,
			Template: &ionoscloudautoscaling.GroupPropertiesTemplate{
				Id: &testAutoscalingGroupVar,
			},
		},
	}
	testAutoscalingGroupGetNew = sdkautoscaling.Group{
		Group: ionoscloudautoscaling.Group{
			Id:         &testAutoscalingGroupVar,
			Properties: testAutoscalingGroupNew.Properties,
			Metadata: &ionoscloudautoscaling.Metadata{
				State: &testAutoscalingStateTestVar,
			},
		},
	}
	testAutoscalingGroupNew = sdkautoscaling.Group{
		Group: ionoscloudautoscaling.Group{
			Properties: &ionoscloudautoscaling.GroupProperties{
				Datacenter:         testAutoscalingGroupGet.Properties.Datacenter,
				TargetReplicaCount: testAutoscalingGroupGet.Properties.TargetReplicaCount,
				MaxReplicaCount:    &testAutoscalingGroupIntNewVar,
				MinReplicaCount:    &testAutoscalingGroupIntNewVar,
				Name:               &testAutoscalingGroupNewVar,
				Policy: &ionoscloudautoscaling.GroupPolicy{
					Metric: (*ionoscloudautoscaling.Metric)(&testAutoscalingGroupNewVar),
					Range:  &testAutoscalingGroupNewVar,
					ScaleInAction: &ionoscloudautoscaling.GroupPolicyAction{
						Amount:         &testAutoscalingGroupFloatNewVar,
						AmountType:     (*ionoscloudautoscaling.ActionAmount)(&testAutoscalingGroupNewVar),
						CooldownPeriod: &testAutoscalingGroupNewVar,
					},
					ScaleInThreshold: &testAutoscalingGroupFloatNewVar,
					ScaleOutAction: &ionoscloudautoscaling.GroupPolicyAction{
						Amount:         &testAutoscalingGroupFloatNewVar,
						AmountType:     (*ionoscloudautoscaling.ActionAmount)(&testAutoscalingGroupNewVar),
						CooldownPeriod: &testAutoscalingGroupNewVar,
					},
					ScaleOutThreshold: &testAutoscalingGroupFloatNewVar,
					Unit:              (*ionoscloudautoscaling.QueryUnit)(&testAutoscalingGroupNewVar),
				},
				Template: &ionoscloudautoscaling.GroupPropertiesTemplate{
					Id: &testAutoscalingGroupNewVar,
				},
			},
		},
	}
	testAutoscalingGroups = sdkautoscaling.Groups{
		GroupCollection: ionoscloudautoscaling.GroupCollection{
			Id:    &testAutoscalingGroupVar,
			Items: &[]ionoscloudautoscaling.Group{testAutoscalingGroupGet.Group, testAutoscalingGroupGetNew.Group},
		},
	}
	testAutoscalingGroupVar         = "test-autoscaling-group"
	testAutoscalingGroupNewVar      = "test-autoscaling-group-new"
	testAutoscalingGroupIntVar      = int64(1)
	testAutoscalingGroupIntNewVar   = int64(2)
	testAutoscalingGroupFloatVar    = float32(3)
	testAutoscalingGroupFloatNewVar = float32(4)
	testAutoscalingGroupErr         = errors.New("autoscaling group test error occurred")
)

func TestPreRunAutoscalingGroupId(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		err := PreRunAutoscalingGroupId(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunAutoscalingGroupIdRequiredFlagErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		err := PreRunAutoscalingGroupId(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().List().Return(testAutoscalingGroups, nil, nil)
		err := RunAutoscalingGroupList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingGroupListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().List().Return(testAutoscalingGroups, nil, testAutoscalingGroupErr)
		err := RunAutoscalingGroupList(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().Get(testAutoscalingGroupVar).Return(&testAutoscalingGroupGet, nil, nil)
		err := RunAutoscalingGroupGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingGroupGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		rm.AutoscalingGroup.EXPECT().Get(testAutoscalingGroupVar).Return(&testAutoscalingGroupGet, nil, testAutoscalingGroupErr)
		err := RunAutoscalingGroupGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupCreate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaxReplicaCount), testAutoscalingGroupIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMinReplicaCount), testAutoscalingGroupIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetReplicaCount), testAutoscalingGroupIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMetric), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRange), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInThreshold), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutThreshold), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUnit), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmount), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmountType), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInCoolDownPeriod), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmount), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmountType), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutCoolDownPeriod), testAutoscalingGroupVar)
		rm.AutoscalingGroup.EXPECT().Create(testAutoscalingGroup).Return(&testAutoscalingGroupGet, nil, nil)
		err := RunAutoscalingGroupCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingGroupCreateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgDataCenterId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaxReplicaCount), testAutoscalingGroupIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMinReplicaCount), testAutoscalingGroupIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetReplicaCount), testAutoscalingGroupIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMetric), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRange), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInThreshold), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutThreshold), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUnit), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmount), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmountType), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInCoolDownPeriod), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmount), testAutoscalingGroupFloatVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmountType), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutCoolDownPeriod), testAutoscalingGroupVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().Create(testAutoscalingGroup).Return(&testAutoscalingGroupGet, nil, testAutoscalingGroupErr)
		err := RunAutoscalingGroupCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupUpdate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaxReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMinReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMetric), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRange), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInThreshold), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutThreshold), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUnit), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmount), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmountType), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInCoolDownPeriod), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmount), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmountType), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutCoolDownPeriod), testAutoscalingGroupNewVar)
		rm.AutoscalingGroup.EXPECT().Get(testAutoscalingGroupVar).Return(&testAutoscalingGroupGet, nil, nil)
		rm.AutoscalingGroup.EXPECT().Update(testAutoscalingGroupVar, testAutoscalingGroupNew).Return(&testAutoscalingGroupGetNew, nil, nil)
		err := RunAutoscalingGroupUpdate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingGroupUpdateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaxReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMinReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMetric), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRange), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInThreshold), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutThreshold), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUnit), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmount), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmountType), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInCoolDownPeriod), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmount), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmountType), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutCoolDownPeriod), testAutoscalingGroupNewVar)
		rm.AutoscalingGroup.EXPECT().Get(testAutoscalingGroupVar).Return(&testAutoscalingGroupGet, nil, nil)
		rm.AutoscalingGroup.EXPECT().Update(testAutoscalingGroupVar, testAutoscalingGroupNew).Return(&testAutoscalingGroupGetNew, nil, testAutoscalingGroupErr)
		err := RunAutoscalingGroupUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupUpdateGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMaxReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMinReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTargetReplicaCount), testAutoscalingGroupIntNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgMetric), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRange), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInThreshold), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutThreshold), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUnit), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmount), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInAmountType), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleInCoolDownPeriod), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmount), testAutoscalingGroupFloatNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutAmountType), testAutoscalingGroupNewVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgScaleOutCoolDownPeriod), testAutoscalingGroupNewVar)
		rm.AutoscalingGroup.EXPECT().Get(testAutoscalingGroupVar).Return(&testAutoscalingGroupGet, nil, testAutoscalingGroupErr)
		err := RunAutoscalingGroupUpdate(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupDelete(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().Delete(testAutoscalingGroupVar).Return(nil, nil)
		err := RunAutoscalingGroupDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingGroupDeleteErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingGroup.EXPECT().Delete(testAutoscalingGroupVar).Return(nil, testAutoscalingGroupErr)
		err := RunAutoscalingGroupDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingGroupDeleteAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		rm.AutoscalingGroup.EXPECT().Delete(testAutoscalingGroupVar).Return(nil, nil)
		err := RunAutoscalingGroupDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingGroupDeleteAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgGroupId), testAutoscalingGroupVar)
		cfg.Stdin = os.Stdin
		err := RunAutoscalingGroupDelete(cfg)
		assert.Error(t, err)
	})
}

func TestGetAutoscalingGroupsCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling group", config.ArgCols), []string{"Name"})
	getAutoscalingGroupCols(core.GetGlobalFlagName("autoscaling group", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetAutoscalingGroupsColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling group", config.ArgCols), []string{"Unknown"})
	getAutoscalingGroupCols(core.GetGlobalFlagName("autoscaling group", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestGetAutoscalingGroupsIds(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	err := os.Setenv(ionoscloudautoscaling.IonosUsernameEnvVar, "user")
	assert.NoError(t, err)
	err = os.Setenv(ionoscloudautoscaling.IonosPasswordEnvVar, "pass")
	assert.NoError(t, err)
	viper.Set(config.ArgServerUrl, config.DefaultApiURL)
	getAutoscalingGroupsIds(w)
	err = w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`401 Unauthorized`)
	assert.True(t, re.Match(b.Bytes()))
}
