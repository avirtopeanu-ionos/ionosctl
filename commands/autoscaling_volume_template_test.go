package commands

import (
	"bufio"
	"bytes"
	"regexp"
	"testing"

	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/ionos-cloud/ionosctl/pkg/utils/clierror"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRunAutoscalingVolumeTemplateList(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().Get(testAutoscalingTemplateVar).Return(&testAutoscalingTemplateGet, nil, nil)
		err := RunAutoscalingVolumeTemplateList(cfg)
		assert.NoError(t, err)
	})
}

func TestRunAutoscalingVolumeTemplateListErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocks) {
		viper.Reset()
		viper.Set(config.ArgOutput, config.DefaultOutputFormat)
		viper.Set(config.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, config.ArgTemplateId), testAutoscalingTemplateVar)
		viper.Set(config.ArgServerUrl, config.DefaultApiURL)
		rm.AutoscalingTemplate.EXPECT().Get(testAutoscalingTemplateVar).Return(&testAutoscalingTemplateGet, nil, testAutoscalingTemplateErr)
		err := RunAutoscalingVolumeTemplateList(cfg)
		assert.Error(t, err)
	})
}

func TestGetAutoscalingVolumeTemplatesCols(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling volume-template", config.ArgCols), []string{"Name"})
	getAutoscalingVolumeTemplateCols(core.GetGlobalFlagName("autoscaling volume-template", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
}

func TestGetAutoscalingVolumeTemplatesColsErr(t *testing.T) {
	defer func(a func()) { clierror.ErrAction = a }(clierror.ErrAction)
	var b bytes.Buffer
	clierror.ErrAction = func() { return }
	w := bufio.NewWriter(&b)
	viper.Set(core.GetGlobalFlagName("autoscaling volume-template", config.ArgCols), []string{"Unknown"})
	getAutoscalingVolumeTemplateCols(core.GetGlobalFlagName("autoscaling volume-template", config.ArgCols), w)
	err := w.Flush()
	assert.NoError(t, err)
	re := regexp.MustCompile(`unknown column Unknown`)
	assert.True(t, re.Match(b.Bytes()))
}
