package commands

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"regexp"
	"strconv"
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
	testAutoscalingTemplateGet = sdkautoscaling.Template{
		Template: ionoscloudautoscaling.Template{
			Id:         &testAutoscalingTemplateVar,
			Properties: &testAutoscalingTemplateProperties.TemplateProperties,
			Metadata: &ionoscloudautoscaling.Metadata{
				State: &testAutoscalingStateTestVar,
			},
		},
	}
	testAutoscalingTemplate = sdkautoscaling.Template{
		Template: ionoscloudautoscaling.Template{
			Properties: &testAutoscalingTemplateProperties.TemplateProperties,
		},
	}
	testAutoscalingTemplateProperties = sdkautoscaling.TemplateProperties{
		TemplateProperties: ionoscloudautoscaling.TemplateProperties{
			Name:             &testAutoscalingTemplateVar,
			Location:         &testAutoscalingTemplateVar,
			AvailabilityZone: (*ionoscloudautoscaling.AvailabilityZone)(&testAutoscalingTemplateVar),
			CpuFamily:        (*ionoscloudautoscaling.CpuFamily)(&testAutoscalingTemplateVar),
			Cores:            &testAutoscalingTemplateIntVar,
			Ram:              &testAutoscalingTemplateSizeIntVar,
			Nics: &[]ionoscloudautoscaling.TemplateNic{
				{
					Lan:  &testAutoscalingTemplateIntVar,
					Name: &testAutoscalingTemplateVar,
				},
			},
			Volumes: &[]ionoscloudautoscaling.TemplateVolume{
				{
					Image:         &testAutoscalingTemplateVar,
					ImagePassword: &testAutoscalingTemplateVar,
					Name:          &testAutoscalingTemplateVar,
					Size:          &testAutoscalingTemplateSizeIntVar,
					SshKeys:       &[]string{testAutoscalingTemplateVar},
					Type:          (*ionoscloudautoscaling.VolumeHwType)(&testAutoscalingTemplateVar),
					UserData:      &testAutoscalingTemplateVar,
				},
			},
		},
	}
	templates = sdkautoscaling.Templates{
		TemplateCollection: ionoscloudautoscaling.TemplateCollection{
			Id:    &testAutoscalingTemplateVar,
			Items: &[]ionoscloudautoscaling.Template{testAutoscalingTemplateGet.Template, testAutoscalingTemplateGet.Template},
		},
	}
	testAutoscalingTemplateVar           = "test-autoscaling template"
	testAutoscalingTemplateIntVar        = int32(1)
	testAutoscalingTemplateSizeStringVar = strconv.Itoa(int(testAutoscalingTemplateSizeIntVar))
	testAutoscalingTemplateSizeIntVar    = int32(256)
	testAutoscalingStateTestVar          = ionoscloudautoscaling.MetadataState("ACTIVE")
	testAutoscalingTemplateErr           = errors.New("autoscaling template test error occurred")
)

func TestPreRunAutoscalingTemplateId(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		err := PreRunAutoscalingTemplateId(cfg)
		assert.NoError(t, err)
	})
}

func TestPreRunAutoscalingTemplateIdRequiredFlagErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		err := PreRunAutoscalingTemplateId(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingTemplateList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().List().Return(templates, nil, nil)
		err := RunAutoscalingTemplateList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingTemplateListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().List().Return(templates, nil, testAutoscalingTemplateErr)
		err := RunAutoscalingTemplateList(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingTemplateGet(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().Get(testAutoscalingTemplateVar).Return(&testAutoscalingTemplateGet, nil, nil)
		err := RunAutoscalingTemplateGet(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingTemplateGetErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		rm.AutoscalingTemplate.EXPECT().Get(testAutoscalingTemplateVar).Return(&testAutoscalingTemplateGet, nil, testAutoscalingTemplateErr)
		err := RunAutoscalingTemplateGet(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingTemplateCreate(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgLocation), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCPUFamily), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgAvailabilityZone), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCores), testAutoscalingTemplateIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRam), testAutoscalingTemplateSizeStringVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateNics), []string{testAutoscalingTemplateVar})
		viper.Set(core.GetFlagName(cfg.NS, config.ArgLanIds), []int{int(testAutoscalingTemplateIntVar)})
		viper.Set(core.GetFlagName(cfg.NS, config.ArgImageId), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgPassword), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateVolume), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgType), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUserData), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSize), testAutoscalingTemplateSizeStringVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSshKeys), []string{testAutoscalingTemplateVar})
		rm.AutoscalingTemplate.EXPECT().Create(testAutoscalingTemplate).Return(&testAutoscalingTemplateGet, nil, nil)
		err := RunAutoscalingTemplateCreate(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingTemplateCreateErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgName), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgLocation), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCPUFamily), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgAvailabilityZone), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgCores), testAutoscalingTemplateIntVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgRam), testAutoscalingTemplateSizeStringVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateNics), []string{testAutoscalingTemplateVar})
		viper.Set(core.GetFlagName(cfg.NS, config.ArgLanIds), []int{int(testAutoscalingTemplateIntVar)})
		viper.Set(core.GetFlagName(cfg.NS, config.ArgImageId), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgPassword), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateVolume), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgType), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgUserData), testAutoscalingTemplateVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSize), testAutoscalingTemplateSizeStringVar)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgSshKeys), []string{testAutoscalingTemplateVar})
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().Create(testAutoscalingTemplate).Return(&testAutoscalingTemplateGet, nil, testAutoscalingTemplateErr)
		err := RunAutoscalingTemplateCreate(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingTemplateDelete(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().Delete(testAutoscalingTemplateVar).Return(nil, nil)
		err := RunAutoscalingTemplateDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingTemplateDeleteErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, true)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().Delete(testAutoscalingTemplateVar).Return(nil, testAutoscalingTemplateErr)
		err := RunAutoscalingTemplateDelete(cfg)
		assert.Error(t, err)
	})
}

func TestRunAutoscalingTemplateDeleteAskForConfirm(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		cfg.Stdin = bytes.NewReader([]byte("YES\n"))
		rm.AutoscalingTemplate.EXPECT().Delete(testAutoscalingTemplateVar).Return(nil, nil)
		err := RunAutoscalingTemplateDelete(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingTemplateDeleteAskForConfirmErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(config.ArgForce, false)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		cfg.Stdin = os.Stdin
		err := RunAutoscalingTemplateDelete(cfg)
		assert.Error(t, err)
	})
}

func TestGetAutoscalingTemplatesCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling template", config.ArgCols), []string{"Name"})
	getAutoscalingTemplateCols(core.GetGlobalFlagName("autoscaling template", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetAutoscalingTemplatesColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling template", config.ArgCols), []string{"Unknown"})
	getAutoscalingTemplateCols(core.GetGlobalFlagName("autoscaling template", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}

func TestGetAutoscalingTemplatesIds(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	err := os.Setenv(ionoscloudautoscaling.IonosUsernameEnvVar, "user")
	assert.NoError(t, err)
	err = os.Setenv(ionoscloudautoscaling.IonosPasswordEnvVar, "pass")
	assert.NoError(t, err)
	viper.Set(config.ArgServerUrl, config.DefaultApiURL)
	getAutoscalingTemplatesIds(w)
	err = w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`401 Unauthorized`)
	assert.True(t, re.Match(b.Bytes()))
}
