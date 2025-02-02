package commands

import (
	"context"
	"os"

	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/query"

	"github.com/fatih/structs"
	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/completer"
	"github.com/ionos-cloud/ionosctl/pkg/constants"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/ionos-cloud/ionosctl/pkg/printer"
	cloudapiv6 "github.com/ionos-cloud/ionosctl/services/cloudapi-v6"
	"github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ResourceCmd() *core.Command {
	ctx := context.TODO()
	resourceCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "resource",
			Aliases:          []string{"res"},
			Short:            "Resource Operations",
			Long:             "The sub-commands of `ionosctl resource` allow you to list, get Resources.",
			TraverseChildren: true,
		},
	}
	globalFlags := resourceCmd.GlobalFlags()
	globalFlags.StringSliceP(constants.ArgCols, "", defaultResourceCols, printer.ColsMessage(defaultResourceCols))
	_ = viper.BindPFlag(core.GetFlagName(resourceCmd.Name(), constants.ArgCols), globalFlags.Lookup(constants.ArgCols))
	_ = resourceCmd.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultResourceCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	list := core.NewCommand(ctx, resourceCmd, core.CommandBuilder{
		Namespace:  "resource",
		Resource:   "resource",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Resources",
		LongDesc:   "Use this command to get a full list of existing Resources. To sort list by Resource Type, use `ionosctl resource get` command.",
		Example:    listResourcesExample,
		PreCmdRun:  core.NoPreRun,
		CmdRun:     RunResourceList,
		InitClient: true,
	})
	list.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)
	list.AddInt32Flag(cloudapiv6.ArgMaxResults, cloudapiv6.ArgMaxResultsShort, cloudapiv6.DefaultMaxResults, cloudapiv6.ArgMaxResultsDescription)
	list.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultListDepth, cloudapiv6.ArgDepthDescription)

	/*
		Get Command
	*/
	getRsc := core.NewCommand(ctx, resourceCmd, core.CommandBuilder{
		Namespace:  "resource",
		Resource:   "resource",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get all Resources of a Type or a specific Resource Type",
		LongDesc:   "Use this command to get all Resources of a Type or a specific Resource Type using its Type and ID.\n\nRequired values to run command:\n\n* Type",
		Example:    getResourceExample,
		PreCmdRun:  PreRunResourceType,
		CmdRun:     RunResourceGet,
		InitClient: true,
	})
	getRsc.AddStringFlag(cloudapiv6.ArgType, "", "", "The specific Type of Resources to retrieve information about", core.RequiredFlagOption())
	_ = getRsc.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgType, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"datacenter", "snapshot", "image", "ipblock", "pcc", "backupunit", "k8s"}, cobra.ShellCompDirectiveNoFileComp
	})
	getRsc.AddUUIDFlag(cloudapiv6.ArgResourceId, cloudapiv6.ArgIdShort, "", "The ID of the specific Resource to retrieve information about")
	_ = getRsc.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgResourceId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.ResourcesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	getRsc.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)

	return resourceCmd
}

func PreRunResourceType(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.Command, c.NS, cloudapiv6.ArgType)
}

func RunResourceList(c *core.CommandConfig) error {
	resourcesListed, resp, err := c.CloudApiV6Services.Users().ListResources()
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getResourcePrint(c, getResources(resourcesListed)))
}

func RunResourceGet(c *core.CommandConfig) error {
	c.Printer.Verbose("Resource with id: %v is getting...", viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgResourceId)))
	if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgResourceId)) {
		resourceListed, resp, err := c.CloudApiV6Services.Users().GetResourceByTypeAndId(
			viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgType)),
			viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgResourceId)),
		)
		if resp != nil {
			c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
		}
		if err != nil {
			return err
		}
		return c.Printer.Print(getResourcePrint(c, getResource(resourceListed)))
	} else {
		resourcesListed, resp, err := c.CloudApiV6Services.Users().GetResourcesByType(viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgType)))
		if resp != nil {
			c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
		}
		if err != nil {
			return err
		}
		return c.Printer.Print(getResourcePrint(c, getResources(resourcesListed)))
	}
}

// Group Resources Commands

func GroupResourceCmd() *core.Command {
	ctx := context.TODO()
	resourceCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "resource",
			Aliases:          []string{"res"},
			Short:            "Group Resource Operations",
			Long:             "The sub-command of `ionosctl group resource` allows you to list Resources from a Group.",
			TraverseChildren: true,
		},
	}

	/*
		List Resources Command
	*/
	listResources := core.NewCommand(ctx, resourceCmd, core.CommandBuilder{
		Namespace:  "group",
		Resource:   "resource",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Resources from a Group",
		LongDesc:   "Use this command to get a list of Resources assigned to a Group. To see more details about existing Resources, use `ionosctl resource` commands.\n\nRequired values to run command:\n\n* Group Id",
		Example:    listGroupResourcesExample,
		PreCmdRun:  PreRunGroupId,
		CmdRun:     RunGroupResourceList,
		InitClient: true,
	})
	listResources.AddInt32Flag(cloudapiv6.ArgMaxResults, cloudapiv6.ArgMaxResultsShort, cloudapiv6.DefaultMaxResults, cloudapiv6.ArgMaxResultsDescription)
	listResources.AddStringSliceFlag(constants.ArgCols, "", defaultResourceCols, printer.ColsMessage(defaultResourceCols))
	_ = listResources.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultResourceCols, cobra.ShellCompDirectiveNoFileComp
	})
	listResources.AddUUIDFlag(cloudapiv6.ArgGroupId, "", "", cloudapiv6.GroupId, core.RequiredFlagOption())
	_ = listResources.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.GroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	return resourceCmd
}

func RunGroupResourceList(c *core.CommandConfig) error {
	// Add Query Parameters for GET Requests
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	c.Printer.Verbose("Listing Resources from Group with ID: %v...", viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId)))
	resourcesListed, resp, err := c.CloudApiV6Services.Groups().ListResources(viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId)), listQueryParams)
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getResourcePrint(c, getResourceGroups(resourcesListed)))
}

// Output Printing

var defaultResourceCols = []string{"ResourceId", "Name", "SecAuthProtection", "Type", "State"}

type ResourcePrint struct {
	ResourceId        string `json:"ResourceId,omitempty"`
	Name              string `json:"Name,omitempty"`
	SecAuthProtection bool   `json:"SecAuthProtection,omitempty"`
	Type              string `json:"Type,omitempty"`
	State             string `json:"State,omitempty"`
}

func getResourcePrint(c *core.CommandConfig, res []resources.Resource) printer.Result {
	r := printer.Result{}
	if c != nil {
		if res != nil {
			r.OutputJSON = res
			r.KeyValue = getResourcesKVMaps(res)
			r.Columns = printer.GetHeadersAllDefault(defaultResourceCols, viper.GetStringSlice(core.GetFlagName(c.Resource, constants.ArgCols)))
		}
	}
	return r
}
func getResource(res *resources.Resource) []resources.Resource {
	ress := make([]resources.Resource, 0)
	if res != nil {
		ress = append(ress, resources.Resource{Resource: res.Resource})
	}
	return ress
}

func getResources(groups resources.Resources) []resources.Resource {
	u := make([]resources.Resource, 0)
	if items, ok := groups.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			u = append(u, resources.Resource{Resource: item})
		}
	}
	return u
}

func getResourceGroups(groups resources.ResourceGroups) []resources.Resource {
	u := make([]resources.Resource, 0)
	if items, ok := groups.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			u = append(u, resources.Resource{Resource: item})
		}
	}
	return u
}

func getResourcesKVMaps(rs []resources.Resource) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(rs))
	for _, r := range rs {
		var rPrint ResourcePrint
		if id, ok := r.GetIdOk(); ok && id != nil {
			rPrint.ResourceId = *id
		}
		if properties, ok := r.GetPropertiesOk(); ok && properties != nil {
			if name, ok := properties.GetNameOk(); ok && name != nil {
				rPrint.Name = *name
			}
			if sh, ok := properties.GetSecAuthProtectionOk(); ok && sh != nil {
				rPrint.SecAuthProtection = *sh
			}
		}
		if typeResource, ok := r.GetTypeOk(); ok && typeResource != nil {
			rPrint.Type = string(*typeResource)
		}
		if metadata, ok := r.GetMetadataOk(); ok && metadata != nil {
			if state, ok := metadata.GetStateOk(); ok && state != nil {
				rPrint.State = *state
			}
		}
		o := structs.Map(rPrint)
		out = append(out, o)
	}
	return out
}
