package query

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/ionos-cloud/ionosctl/pkg/constants"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	cloudapiv6 "github.com/ionos-cloud/ionosctl/services/cloudapi-v6"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	testFiltersVar    = []string{"test", "testing", "filter", "validate"}
	testFilterVar     = "name"
	testMaxResultsVar = int32(2)
)

func TestValidateFilters(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{"test=test"})
		err := ValidateFilters(cfg, testFiltersVar, "")
		assert.NoError(t, err)
	})
}

func TestValidateFiltersLengthErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		err := ValidateFilters(cfg, testFiltersVar, "")
		assert.Error(t, err)
	})
}

func TestValidateFiltersFormatErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{"test"})
		err := ValidateFilters(cfg, testFiltersVar, "")
		assert.Error(t, err)
	})
}

func TestValidateFiltersInvalidsErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{"name=test", "location=test"})
		err := ValidateFilters(cfg, testFiltersVar, "")
		assert.Error(t, err)
	})
}

func TestValidateFiltersInvalidErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.PreCmdConfigTest(t, w, func(cfg *core.PreCommandConfig) {
		viper.Reset()
		viper.Set(constants.ArgQuiet, false)
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{"name=test"})
		err := ValidateFilters(cfg, testFiltersVar, "")
		assert.Error(t, err)
	})
}

func TestGetListQueryParams(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(constants.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{"name=test", "location=test"})
		//cfg.Command.Command.Flags().Set(cloudapiv6.ArgOrderBy, testFilterVar)
		//cfg.Command.Command.Flags().(cloudapiv6.ArgMaxResults, testMaxResultsVar)
		result, err := GetListQueryParams(cfg)
		assert.NoError(t, err)
		assert.True(t, result.Filters != nil)
		filtersKV := *result.Filters
		assert.True(t, filtersKV["name"] == "test")
		assert.True(t, filtersKV["location"] == "test")
		//assert.True(t, *result.OrderBy == testFilterVar) Muted temporarily due to viper pflag mapping removal
		//assert.True(t, *result.MaxResults == testMaxResultsVar)
	})
}

func TestGetListQueryParamsErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	core.CmdConfigTest(t, w, func(cfg *core.CommandConfig, rm *core.ResourcesMocksTest) {
		viper.Reset()
		viper.Set(constants.ArgOutput, constants.DefaultOutputFormat)
		viper.Set(constants.ArgQuiet, false)
		viper.Set(core.GetFlagName(cfg.NS, cloudapiv6.ArgFilters), []string{"name"})
		_, err := GetListQueryParams(cfg)
		assert.Error(t, err)
	})
}
