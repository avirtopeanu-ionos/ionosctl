package commands

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

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
			Short:            "VM Autoscaling Template Operations",
			Long:             "The sub-commands of `ionosctl autoscaling template` allow you to create, list, get, update and delete VM Autoscaling Templates.",
			TraverseChildren: true,
		},
	}
	globalFlags := autoscalingTemplateCmd.GlobalFlags()
	globalFlags.StringSliceP(config.ArgCols, "", defaultAutoscalingTemplateCols, utils.ColsMessage(allAutoscalingTemplateCols))
	_ = viper.BindPFlag(core.GetGlobalFlagName(autoscalingTemplateCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = autoscalingTemplateCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allAutoscalingTemplateCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "template",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List VM Autoscaling Templates",
		LongDesc:   "Use this command to retrieve a complete list of VM Autoscaling Templates provisioned under your account.",
		Example:    listTemplateAutoscalingExample,
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
		ShortDesc:  "Get a VM Autoscaling Template",
		LongDesc:   "Use this command to retrieve details about a VM Autoscaling Template by using its ID.\n\nRequired values to run command:\n\n* VM Autoscaling Template Id",
		Example:    getTemplateAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingTemplateId,
		CmdRun:     RunAutoscalingTemplateGet,
		InitClient: true,
	})
	get.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Create Command
	*/
	create := core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "template",
		Verb:      "create",
		Aliases:   []string{"c"},
		ShortDesc: "Create a VM Autoscaling Template",
		LongDesc: `Use this command to create a VM Autoscaling Template. The VM Autoscaling Template contains information for the VMs. You can specify the name, location, availability zone, cores, cpu family for the VMs.

Regarding the Ram size, it must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB.

Right now, the VM Autoscaling Template supports only one Template Volume. Important: the volume created will NOT be deleted on SCALE IN type of Autoscaling Actions. If you want to create a Volume Template, you need to provide Image Id. If you want to see the Volume Template properties, use ` + "`" + `ionosctl autoscaling volume-template list` + "`" + ` command.

Also, the VM Autoscaling Template supports multiple NIC Templates. To create a VM Autoscaling Template with multiple NIC Templates use ` + "`" + `--lan-ids "LAN_ID1,LAN_ID2"` + "`" + ` and ` + "`" + `--template-nics "NAME1,NAME2"` + "`" + ` options. It is recommended to use both options. If you want to see the NIC Templates properties, use ` + "`" + `ionosctl autoscaling nic-template list` + "`" + ` command.`,
		Example:    createTemplateAutoscalingExample,
		PreCmdRun:  noPreRun,
		CmdRun:     RunAutoscalingTemplateCreate,
		InitClient: true,
	})
	create.AddStringFlag(config.ArgName, config.ArgNameShort, "Unnamed VM Autoscaling Template", "Name of the VM Autoscaling Template")
	create.AddStringFlag(config.ArgLocation, config.ArgLocationShort, "de/txl", "Location for the VM Autoscaling Template")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgLocation, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getLocationIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgAvailabilityZone, config.ArgAvailabilityZoneShort, "AUTO", "Zone where the VMs created using this VM Autoscaling Template")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgCpuFamily, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"AUTO", "ZONE_1", "ZONE_2"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddIntFlag(config.ArgCores, "", 1, "The total number of cores for the VMs. Minimum: 1")
	create.AddStringFlag(config.ArgRam, "", "2048", "The amount of memory for the VMs. Size must be specified in multiples of 256. e.g. --ram 2048 or --ram 2048MB")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgRam, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"256MB", "512MB", "1024MB", "2048MB", "2GB", "3GB", "4GB", "5GB", "10GB", "16GB"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgCPUFamily, "", "", "CPU family for the VMs created using the VM Autoscaling Template. If null, the VM will be created with the default CPU family from the assigned location")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgCPUFamily, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"AMD_OPTERON", "INTEL_XEON", "INTEL_SKYLAKE"}, cobra.ShellCompDirectiveNoFileComp
	})
	// Flags for NIC Templates
	create.AddIntSliceFlag(config.ArgLanIds, "", []int{1}, "Lan Ids for the NIC Templates. Minimum value for Lan Id: 1")
	create.AddStringSliceFlag(config.ArgTemplateNics, "", []string{"Unnamed VM Autoscaling NIC Template"}, "Names for the NIC Templates")
	// Flags for Volume Template
	create.AddStringFlag(config.ArgImageId, "", "", "Image installed on the Volume. Only Id of the Image is supported currently. Required flag when creating a Volume Template")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgImageId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getImageIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgPassword, config.ArgPasswordShort, "abcde1234", "Image password for the Volume Template")
	create.AddStringFlag(config.ArgUserData, "", "", "User-Data (Cloud Init) for the Volume Template")
	create.AddStringFlag(config.ArgTemplateVolume, "", "Unnamed VM Autoscaling Template Volume", "Name of the Volume Template")
	create.AddStringFlag(config.ArgSize, "", strconv.Itoa(config.DefaultVolumeSize), "User-defined size for this template volume in GB. e.g.: --size 10 or --size 10GB.")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgSize, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"10GB", "20GB", "50GB", "100GB", "1TB"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgType, "", "HDD", "Type of the Volume")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgLicenceType, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"HDD", "SSD", "SSD_PREMIUM", "SSD_STANDARD"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringSliceFlag(config.ArgSshKeys, "", []string{""}, "SSH Keys that have access to the Volume")

	/*
		Delete Command
	*/
	deleteCmd := core.NewCommand(ctx, autoscalingTemplateCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "template",
		Verb:      "delete",
		Aliases:   []string{"d"},
		ShortDesc: "Delete a VM Autoscaling Template",
		LongDesc: `Use this command to delete a specified VM Autoscaling Template from your account.

Required values to run command:

* VM Autoscaling Template Id`,
		Example:    deleteTemplateAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingTemplateId,
		CmdRun:     RunAutoscalingTemplateDelete,
		InitClient: true,
	})
	deleteCmd.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
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
	return c.Printer.Print(getTemplatePrint(nil, c, getAutoscalingTemplates(autoscalingTemplates)))
}

func RunAutoscalingTemplateGet(c *core.CommandConfig) error {
	c.Printer.Verbose("VM Autoscaling Template with ID %v is getting...", viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	autoTemplate, _, err := c.AutoscalingTemplates().Get(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(nil, c, []sdkAutoscaling.Template{*autoTemplate}))
}

func RunAutoscalingTemplateCreate(c *core.CommandConfig) error {
	templateProperties, err := getNewAutoscalingTemplate(c)
	if err != nil {
		return err
	}
	c.Printer.Verbose("Creating VM Autoscaling Template...")
	dc, resp, err := c.AutoscalingTemplates().Create(sdkAutoscaling.Template{
		Template: ionoscloudAutoscaling.Template{
			Properties: &templateProperties.TemplateProperties,
		},
	})
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(resp, c, []sdkAutoscaling.Template{*dc}))
}

func RunAutoscalingTemplateDelete(c *core.CommandConfig) error {
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete VM autoscaling template"); err != nil {
		return err
	}
	c.Printer.Verbose("VM Autoscaling Template with ID %v is deleting...", viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	resp, err := c.AutoscalingTemplates().Delete(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getTemplatePrint(resp, c, nil))
}

func getAutoscalingTemplates(templates sdkAutoscaling.Templates) []sdkAutoscaling.Template {
	tpls := make([]sdkAutoscaling.Template, 0)
	for _, tpl := range *templates.Items {
		tpls = append(tpls, sdkAutoscaling.Template{Template: tpl})
	}
	return tpls
}

func getNewAutoscalingTemplate(c *core.CommandConfig) (*sdkAutoscaling.TemplateProperties, error) {
	input := ionoscloudAutoscaling.TemplateProperties{}
	// Autoscaling Template - VM Properties
	input.SetName(viper.GetString(core.GetFlagName(c.NS, config.ArgName)))
	c.Printer.Verbose("Property Name set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgName)))
	input.SetLocation(viper.GetString(core.GetFlagName(c.NS, config.ArgLocation)))
	c.Printer.Verbose("Property Location set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgLocation)))
	input.SetAvailabilityZone(ionoscloudAutoscaling.AvailabilityZone(viper.GetString(core.GetFlagName(c.NS, config.ArgAvailabilityZone))))
	c.Printer.Verbose("Property Availability Zone set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgAvailabilityZone)))
	input.SetCores(viper.GetInt32(core.GetFlagName(c.NS, config.ArgCores)))
	c.Printer.Verbose("Property Cores set: %v", viper.GetInt32(core.GetFlagName(c.NS, config.ArgCores)))
	size, err := utils.ConvertSize(viper.GetString(core.GetFlagName(c.NS, config.ArgRam)), utils.MegaBytes)
	if err != nil {
		return nil, err
	}
	input.SetRam(int32(size))
	c.Printer.Verbose("Property RAM set: %vMB", viper.GetInt32(core.GetFlagName(c.NS, config.ArgRam)))
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgCPUFamily)) {
		input.SetCpuFamily(ionoscloudAutoscaling.CpuFamily(viper.GetString(core.GetFlagName(c.NS, config.ArgCPUFamily))))
		c.Printer.Verbose("Property CPU Family set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgCpuFamily)))
	}

	// Autoscaling NIC Template
	inputNics := make([]ionoscloudAutoscaling.TemplateNic, 0)
	nicNames := viper.GetStringSlice(core.GetFlagName(c.NS, config.ArgTemplateNics))
	c.Printer.Verbose("Property Name for NIC Templates set: %v", viper.GetStringSlice(core.GetFlagName(c.NS, config.ArgTemplateNics)))
	lanIds := viper.GetIntSlice(core.GetFlagName(c.NS, config.ArgLanIds))
	c.Printer.Verbose("Property Lan for NIC Templates set: %v", viper.GetStringSlice(core.GetFlagName(c.NS, config.ArgLanIds)))
	if len(nicNames) != len(lanIds) {
		return nil, errors.New("error creating NIC Templates. Hint: please use the `--lan-ids` and the `--template-nics` options with the same amount of values")
	} else {
		for i := 0; i < len(nicNames); i++ {
			lanId := int32(lanIds[i])
			inputNics = append(inputNics, ionoscloudAutoscaling.TemplateNic{
				Lan:  &lanId,
				Name: &nicNames[i],
			})
		}
	}
	input.SetNics(inputNics)

	// Autoscaling Volume Template
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgImageId)) {
		inputVolumes := make([]ionoscloudAutoscaling.TemplateVolume, 0)
		inputVolume := ionoscloudAutoscaling.TemplateVolume{}
		// Set Properties for Autoscaling Volume Template
		inputVolume.SetName(viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateVolume)))
		c.Printer.Verbose("Property Name for Volume Template set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateVolume)))
		inputVolume.SetImage(viper.GetString(core.GetFlagName(c.NS, config.ArgImageId)))
		c.Printer.Verbose("Property Image for Volume Template set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgImageId)))
		inputVolume.SetImagePassword(viper.GetString(core.GetFlagName(c.NS, config.ArgPassword)))
		c.Printer.Verbose("Property Password for Volume Template set")
		inputVolume.SetType(ionoscloudAutoscaling.VolumeHwType(viper.GetString(core.GetFlagName(c.NS, config.ArgType))))
		c.Printer.Verbose("Property Type for Volume Template set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgType)))
		inputVolume.SetUserData(viper.GetString(core.GetFlagName(c.NS, config.ArgUserData)))
		c.Printer.Verbose("Property User Data for Volume Template set: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgUserData)))
		inputVolume.SetSshKeys(viper.GetStringSlice(core.GetFlagName(c.NS, config.ArgSshKeys)))
		c.Printer.Verbose("Property SSH Keys for Volume Template set")
		volumeSize, err := utils.ConvertSize(viper.GetString(core.GetFlagName(c.NS, config.ArgSize)), utils.GigaBytes)
		if err != nil {
			return nil, err
		}
		inputVolume.SetSize(int32(volumeSize))
		c.Printer.Verbose("Property Size for Volume Template set: %v GB", viper.GetString(core.GetFlagName(c.NS, config.ArgSize)))
		inputVolumes = append(inputVolumes, inputVolume)
		input.SetVolumes(inputVolumes)
	}

	return &sdkAutoscaling.TemplateProperties{
		TemplateProperties: input,
	}, nil
}

// Output Printing

var (
	defaultAutoscalingTemplateCols = []string{"TemplateId", "Name", "Location", "CpuFamily", "AvailabilityZone", "Ram", "State"}
	allAutoscalingTemplateCols     = []string{"TemplateId", "Name", "Location", "CpuFamily", "AvailabilityZone", "Cores", "Ram", "State"}
)

type AutoscalingTemplatePrint struct {
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
		}
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getAutoscalingTemplatesKVMaps(dcs)
			r.Columns = getAutoscalingTemplateCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getAutoscalingTemplateCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultAutoscalingTemplateCols
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

func getAutoscalingTemplatesKVMaps(templates []sdkAutoscaling.Template) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(templates))
	for _, template := range templates {
		var templatePrint AutoscalingTemplatePrint
		if idOk, ok := template.GetIdOk(); ok && idOk != nil {
			templatePrint.TemplateId = *idOk
		}
		if properties, ok := template.GetPropertiesOk(); ok && properties != nil {
			if nameOk, ok := properties.GetNameOk(); ok && nameOk != nil {
				templatePrint.Name = *nameOk
			}
			if locationOk, ok := properties.GetLocationOk(); ok && locationOk != nil {
				templatePrint.Location = *locationOk
			}
			if cpuFamilyOk, ok := properties.GetCpuFamilyOk(); ok && cpuFamilyOk != nil {
				templatePrint.CpuFamily = string(*cpuFamilyOk)
			}
			if availabilityZoneOk, ok := properties.GetAvailabilityZoneOk(); ok && availabilityZoneOk != nil {
				templatePrint.AvailabilityZone = string(*availabilityZoneOk)
			}
			if ramOk, ok := properties.GetRamOk(); ok && ramOk != nil {
				templatePrint.Ram = fmt.Sprintf("%vMB", *ramOk)
			}
			if coresOk, ok := properties.GetCoresOk(); ok && coresOk != nil {
				templatePrint.Cores = *coresOk
			}
		}
		if metadataOk, ok := template.GetMetadataOk(); ok && metadataOk != nil {
			if stateOk, ok := metadataOk.GetStateOk(); ok && stateOk != nil {
				templatePrint.State = string(*stateOk)
			}
		}
		o := structs.Map(templatePrint)
		out = append(out, o)
	}
	return out
}

func getAutoscalingTemplatesIds(outErr io.Writer) []string {
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
