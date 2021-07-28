package commands

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fatih/structs"
	"github.com/ionos-cloud/ionosctl/pkg/config"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	sdkAutoscaling "github.com/ionos-cloud/ionosctl/pkg/resources/autoscaling"
	"github.com/ionos-cloud/ionosctl/pkg/utils"
	"github.com/ionos-cloud/ionosctl/pkg/utils/clierror"
	"github.com/ionos-cloud/ionosctl/pkg/utils/printer"
	ionoscloudautoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func autoscalingVolumeTemplate() *core.Command {
	ctx := context.TODO()
	autoscalingVolumeTemplateCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "volume-template",
			Aliases:          []string{"v"},
			Short:            "Autoscaling Volume Template Operations",
			Long:             "The sub-command of `ionosctl autoscaling volume-template` allows you to list Volume Templates from an Autoscaling Template.",
			TraverseChildren: true,
		},
	}
	globalFlags := autoscalingVolumeTemplateCmd.GlobalFlags()
	globalFlags.StringSliceP(config.ArgCols, "", defaultVolumeTemplateCols, utils.ColsMessage(defaultVolumeTemplateCols))
	_ = viper.BindPFlag(core.GetGlobalFlagName(autoscalingVolumeTemplateCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = autoscalingVolumeTemplateCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultVolumeTemplateCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	list := core.NewCommand(ctx, autoscalingVolumeTemplateCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "volume-template",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Volume Templates from an Autoscaling Template",
		LongDesc:   "Use this command to retrieve a complete list of Volume Templates from a specific Autoscaling Template provisioned under your account.\n\nRequired values to run command:\n\n* Autoscaling Template Id",
		Example:    listVolumeTemplateAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingTemplateId,
		CmdRun:     RunAutoscalingVolumeTemplateList,
		InitClient: true,
	})
	list.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	return autoscalingVolumeTemplateCmd
}

func RunAutoscalingVolumeTemplateList(c *core.CommandConfig) error {
	autoscalingTpl, _, err := c.AutoscalingTemplates().Get(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	if err != nil {
		return err
	}
	if propertiesOk, ok := autoscalingTpl.GetPropertiesOk(); ok && propertiesOk != nil {
		if volumesOk, ok := propertiesOk.GetVolumesOk(); ok && volumesOk != nil {
			return c.Printer.Print(getAutoscalingVolumeTemplatePrint(c, getAutoscalingVolumeTemplates(volumesOk)))
		} else {
			return errors.New("error getting volumes from autoscaling template")
		}
	} else {
		return errors.New("error getting properties from autoscaling template")
	}
}

func getAutoscalingVolumeTemplates(templates *[]ionoscloudautoscaling.TemplateVolume) []sdkAutoscaling.TemplateVolume {
	tpls := make([]sdkAutoscaling.TemplateVolume, 0)
	if templates != nil {
		for _, tpl := range *templates {
			tpls = append(tpls, sdkAutoscaling.TemplateVolume{TemplateVolume: tpl})
		}
	}
	return tpls
}

// Output Printing

var defaultVolumeTemplateCols = []string{"Name", "Size", "Type", "Image", "UserData"}

type VolumeTemplatePrint struct {
	Image    string `json:"Image,omitempty"`
	Name     string `json:"Name,omitempty"`
	Size     string `json:"Size,omitempty"`
	Type     string `json:"Type,omitempty"`
	UserData string `json:"UserData,omitempty"`
}

func getAutoscalingVolumeTemplatePrint(c *core.CommandConfig, dcs []sdkAutoscaling.TemplateVolume) printer.Result {
	r := printer.Result{}
	if c != nil {
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getAutoscalingVolumeTemplatesKVMaps(dcs)
			r.Columns = getAutoscalingVolumeTemplateCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getAutoscalingVolumeTemplateCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultVolumeTemplateCols
	}
	columnsMap := map[string]string{
		"Name":     "Name",
		"Size":     "Size",
		"Type":     "Type",
		"Image":    "Image",
		"UserData": "UserData",
	}
	var autoscalingVolumeTemplateCols []string
	for _, k := range cols {
		col := columnsMap[k]
		if col != "" {
			autoscalingVolumeTemplateCols = append(autoscalingVolumeTemplateCols, col)
		} else {
			clierror.CheckError(errors.New("unknown column "+k), outErr)
		}
	}
	return autoscalingVolumeTemplateCols
}

func getAutoscalingVolumeTemplatesKVMaps(templates []sdkAutoscaling.TemplateVolume) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(templates))
	for _, template := range templates {
		var templatePrint VolumeTemplatePrint
		if nameOk, ok := template.GetNameOk(); ok && nameOk != nil {
			templatePrint.Name = *nameOk
		}
		if sizeOk, ok := template.GetSizeOk(); ok && sizeOk != nil {
			templatePrint.Size = fmt.Sprintf("%vGB", *sizeOk)
		}
		if typeOk, ok := template.GetTypeOk(); ok && typeOk != nil {
			templatePrint.Type = string(*typeOk)
		}
		if imageOk, ok := template.GetImageOk(); ok && imageOk != nil {
			templatePrint.Image = *imageOk
		}
		if userDataOk, ok := template.GetUserDataOk(); ok && userDataOk != nil {
			templatePrint.UserData = *userDataOk
		}
		o := structs.Map(templatePrint)
		out = append(out, o)
	}
	return out
}
