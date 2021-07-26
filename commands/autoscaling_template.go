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
	ionoscloudAutoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func autoscalingTemplate() *core.Command {
	ctx := context.TODO()
	autoscalingTemplateCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "template",
			Aliases:          []string{"t"},
			Short:            "Autoscaling Template Operations",
			Long:             "The sub-commands of `ionosctl autoscaling template` allow you to create, list, get, update and delete Autoscaling Templates.",
			TraverseChildren: true,
		},
	}
	globalFlags := autoscalingTemplateCmd.GlobalFlags()
	globalFlags.StringSliceP(config.ArgCols, "", defaultTemplateCols, utils.ColsMessage(allTemplateCols))
	_ = viper.BindPFlag(core.GetGlobalFlagName(autoscalingTemplateCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = autoscalingTemplateCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allTemplateCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "template",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Autoscaling Templates",
		LongDesc:   "Use this command to retrieve a complete list of Autoscaling Templates provisioned under your account.",
		Example:    "",
		PreCmdRun:  noPreRun,
		CmdRun:     RunAutoscalingTemplateList,
		InitClient: true,
	})

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "template",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get an Autoscaling Template",
		LongDesc:   "Use this command to retrieve details about an Autoscaling Template by using its ID.\n\nRequired values to run command:\n\n* Autoscaling Template Id",
		Example:    "",
		PreCmdRun:  PreRunAutoscalingTemplateId,
		CmdRun:     RunAutoscalingTemplateGet,
		InitClient: true,
	})
	get.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Create Command
	*/
	create := core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "template",
		Verb:      "create",
		Aliases:   []string{"c"},
		ShortDesc: "Create an Autoscaling Template",
		LongDesc: `Use this command to create an Autoscaling Template. The Autoscaling Template contains information for the VMs. You can specify the name, location, availability zone, cores, cpu family for the VMs.

Regarding the ram size, it must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option.`,
		Example:    "",
		PreCmdRun:  noPreRun,
		CmdRun:     RunAutoscalingTemplateCreate,
		InitClient: true,
	})
	create.AddStringFlag(config.ArgName, config.ArgNameShort, "Unnamed Autoscaling Template", "Name of the Autoscaling Template")
	create.AddStringFlag(config.ArgLocation, config.ArgLocationShort, "de/txl", "Location for the Autoscaling Template")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgLocation, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getLocationIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgAvailabilityZone, config.ArgAvailabilityZoneShort, "AUTO", "Zone where the VMs created using this Autoscaling Template")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgCpuFamily, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"AUTO", "ZONE_1", "ZONE_2"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddIntFlag(config.ArgCores, "", 1, "The total number of cores for the VMs. Minimum: 1")
	create.AddStringFlag(config.ArgRam, "", "", "The amount of memory for the VMs. Size must be specified in multiples of 256. e.g. --ram 2048 or --ram 2048MB")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgRam, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"256MB", "512MB", "1024MB", "2048MB", "2GB", "3GB", "4GB", "5GB", "10GB", "16GB"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgCPUFamily, "", "", "CPU family for the VMs created using the Autoscaling Template. If null, the VM will be created with the default CPU family from the assigned location")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgCPUFamily, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"AMD_OPTERON", "INTEL_XEON", "INTEL_SKYLAKE"}, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Delete Command
	*/
	deleteCmd := core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "template",
		Verb:      "delete",
		Aliases:   []string{"d"},
		ShortDesc: "Delete an Autoscaling Template",
		LongDesc: `Use this command to delete a specified Autoscaling Template from your account.

Required values to run command:

* Autoscaling Template Id`,
		Example:    "",
		PreCmdRun:  PreRunAutoscalingTemplateId,
		CmdRun:     RunAutoscalingTemplateDelete,
		InitClient: true,
	})
	deleteCmd.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	return autoscalingTemplateCmd
}

func PreRunAutoscalingTemplateId(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgTemplateId)
}

func RunAutoscalingTemplateList(c *core.CommandConfig) error {
	autoscalingTemplates, _, err := c.AutoscalingTemplates().List()
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(nil, c, getAutoscalinTemplates(autoscalingTemplates)))
}

func RunAutoscalingTemplateGet(c *core.CommandConfig) error {
	autoTemplate, _, err := c.AutoscalingTemplates().Get(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(nil, c, []sdkAutoscaling.Template{*autoTemplate}))
}

func RunAutoscalingTemplateCreate(c *core.CommandConfig) error {
	dc, resp, err := c.AutoscalingTemplates().Create(sdkAutoscaling.Template{
		Template: ionoscloudAutoscaling.Template{
			Properties: &ionoscloudAutoscaling.TemplateProperties{
				AvailabilityZone: nil,
				Cores:            nil,
				CpuFamily:        nil,
				Location:         nil,
				Name:             nil,
				Nics:             nil,
				Ram:              nil,
				Volumes:          nil,
			},
		},
	})
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(resp, c, []sdkAutoscaling.Template{*dc}))
}

func RunAutoscalingTemplateDelete(c *core.CommandConfig) error {
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete autoscaling template"); err != nil {
		return err
	}
	resp, err := c.AutoscalingTemplates().Delete(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(resp, c, nil))
}

func getAutoscalinTemplates(templates sdkAutoscaling.Templates) []sdkAutoscaling.Template {
	tpls := make([]sdkAutoscaling.Template, 0)
	for _, tpl := range *templates.Items {
		tpls = append(tpls, sdkAutoscaling.Template{Template: tpl})
	}
	return tpls
}

// Output Printing

var (
	defaultTemplateCols = []string{"TemplateId", "Name", "Location", "CpuFamily", "AvailabilityZone", "State"}
	allTemplateCols     = []string{"TemplateId", "Name", "Location", "CpuFamily", "AvailabilityZone", "Cores", "Ram", "State"}
)

type TemplatePrint struct {
	TemplateId       string `json:"TemplateId,omitempty"`
	AvailabilityZone string `json:"AvailabilityZone,omitempty"`
	Cores            int32  `json:"Cores,omitempty"`
	CpuFamily        string `json:"CpuFamily,omitempty"`
	Location         string `json:"Location,omitempty"`
	Name             string `json:"Name,omitempty"`
	Ram              string `json:"Ram,omitempty"`
	State            string `json:"State,omitempty"`
}

func getTemplatePrint(resp *sdkAutoscaling.Response, c *core.CommandConfig, dcs []sdkAutoscaling.Template) printer.Result {
	r := printer.Result{}
	if c != nil {
		if resp != nil {
			r.Resource = c.Resource
			r.Verb = c.Verb
			r.WaitForRequest = viper.GetBool(core.GetFlagName(c.NS, config.ArgWaitForRequest))
		}
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getTemplatesKVMaps(dcs)
			r.Columns = getTemplateCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getTemplateCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultTemplateCols
	}

	columnsMap := map[string]string{
		"TemplateId":       "TemplateId",
		"Name":             "Name",
		"Location":         "Location",
		"CpuFamily":        "CpuFamily",
		"AvailabilityZone": "AvailabilityZone",
		"Cores":            "Cores",
		"Ram":              "Ram",
		"State":            "State",
	}
	var autoscalingTemplateCols []string
	for _, k := range cols {
		col := columnsMap[k]
		if col != "" {
			autoscalingTemplateCols = append(autoscalingTemplateCols, col)
		} else {
			clierror.CheckError(errors.New("unknown column "+k), outErr)
		}
	}
	return autoscalingTemplateCols
}

func getTemplatesKVMaps(dcs []sdkAutoscaling.Template) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(dcs))
	for _, dc := range dcs {
		var dcPrint TemplatePrint
		if dcid, ok := dc.GetIdOk(); ok && dcid != nil {
			dcPrint.TemplateId = *dcid
		}
		if properties, ok := dc.GetPropertiesOk(); ok && properties != nil {
			if name, ok := properties.GetNameOk(); ok && name != nil {
				dcPrint.Name = *name
			}
			if loc, ok := properties.GetLocationOk(); ok && loc != nil {
				dcPrint.Location = *loc
			}
		}
		if metadata, ok := dc.GetMetadataOk(); ok && metadata != nil {
			if state, ok := metadata.GetStateOk(); ok && state != nil {
				dcPrint.State = string(*state)
			}
		}
		o := structs.Map(dcPrint)
		out = append(out, o)
	}
	return out
}

func getTemplatesIds(outErr io.Writer) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := sdkAutoscaling.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	autoscalingTemplateSvc := sdkAutoscaling.NewTemplateService(clientSvc.Get(), context.TODO())
	autoscalingTemplates, _, err := autoscalingTemplateSvc.List()
	clierror.CheckError(err, outErr)
	templateIds := make([]string, 0)
	if items, ok := autoscalingTemplates.TemplateCollection.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				templateIds = append(templateIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return templateIds
}
