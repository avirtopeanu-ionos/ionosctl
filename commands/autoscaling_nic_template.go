package commands

import (
	"context"
	"errors"
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

func autoscalingNicTemplate() *core.Command {
	ctx := context.TODO()
	autoscalingNicTemplateCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "nic-template",
			Aliases:          []string{"n"},
			Short:            "VM Autoscaling NIC Template Operations",
			Long:             "The sub-command of `ionosctl autoscaling nic-template` allows you to list NIC Templates from a VM Autoscaling Template.",
			TraverseChildren: true,
		},
	}
	globalFlags := autoscalingNicTemplateCmd.GlobalFlags()
	globalFlags.StringSliceP(config.ArgCols, "", defaultNicTemplateCols, utils.ColsMessage(defaultNicTemplateCols))
	_ = viper.BindPFlag(core.GetGlobalFlagName(autoscalingNicTemplateCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = autoscalingNicTemplateCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultNicTemplateCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	list := core.NewCommand(ctx, autoscalingNicTemplateCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "nic-template",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List NIC Templates from a VM Autoscaling Template",
		LongDesc:   "Use this command to retrieve a complete list of NIC Templates from a specific VM Autoscaling Template provisioned under your account.\n\nRequired values to run command:\n\n* VM Autoscaling Template Id",
		Example:    listNicTemplateAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingTemplateId,
		CmdRun:     RunAutoscalingNicTemplateList,
		InitClient: true,
	})
	list.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	return autoscalingNicTemplateCmd
}

func RunAutoscalingNicTemplateList(c *core.CommandConfig) error {
	c.Printer.Verbose("Getting NIC Templates for Template with ID: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	autoscalingTpl, _, err := c.AutoscalingTemplates().Get(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	if err != nil {
		return err
	}
	if propertiesOk, ok := autoscalingTpl.GetPropertiesOk(); ok && propertiesOk != nil {
		if nicsOk, ok := propertiesOk.GetNicsOk(); ok && nicsOk != nil {
			return c.Printer.Print(getAutoscalingNicTemplatePrint(c, getAutoscalingNicTemplates(nicsOk)))
		} else {
			return errors.New("error getting NICs from autoscaling template")
		}
	} else {
		return errors.New("error getting properties from autoscaling template")
	}
}

func getAutoscalingNicTemplates(templates *[]ionoscloudautoscaling.TemplateNic) []sdkAutoscaling.TemplateNic {
	tpls := make([]sdkAutoscaling.TemplateNic, 0)
	if templates != nil {
		for _, tpl := range *templates {
			tpls = append(tpls, sdkAutoscaling.TemplateNic{TemplateNic: tpl})
		}
	}
	return tpls
}

// Output Printing

var defaultNicTemplateCols = []string{"Name", "LanId"}

type NicTemplatePrint struct {
	LanId int32  `json:"LanId,omitempty"`
	Name  string `json:"Name,omitempty"`
}

func getAutoscalingNicTemplatePrint(c *core.CommandConfig, dcs []sdkAutoscaling.TemplateNic) printer.Result {
	r := printer.Result{}
	if c != nil {
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getAutoscalingNicTemplatesKVMaps(dcs)
			r.Columns = getAutoscalingNicTemplateCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getAutoscalingNicTemplateCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultNicTemplateCols
	}
	columnsMap := map[string]string{
		"Name":  "Name",
		"LanId": "LanId",
	}
	var autoscalingNicTemplateCols []string
	for _, k := range cols {
		col := columnsMap[k]
		if col != "" {
			autoscalingNicTemplateCols = append(autoscalingNicTemplateCols, col)
		} else {
			clierror.CheckError(errors.New("unknown column "+k), outErr)
		}
	}
	return autoscalingNicTemplateCols
}

func getAutoscalingNicTemplatesKVMaps(templates []sdkAutoscaling.TemplateNic) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(templates))
	for _, template := range templates {
		var templatePrint NicTemplatePrint
		if nameOk, ok := template.GetNameOk(); ok && nameOk != nil {
			templatePrint.Name = *nameOk
		}
		if lanOk, ok := template.GetLanOk(); ok && lanOk != nil {
			templatePrint.LanId = *lanOk
		}
		o := structs.Map(templatePrint)
		out = append(out, o)
	}
	return out
}
