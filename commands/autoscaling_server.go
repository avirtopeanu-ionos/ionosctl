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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func autoscalingServer() *core.Command {
	ctx := context.TODO()
	autoscalingServerCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "server",
			Aliases:          []string{"svr"},
			Short:            "VM Autoscaling Server Operations",
			Long:             "The sub-commands of `ionosctl autoscaling server` allow you to list, get Servers from a VM Autoscaling Group.",
			TraverseChildren: true,
		},
	}

	/*
		List Command
	*/
	list := core.NewCommand(ctx, autoscalingServerCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "server",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Servers from a VM Autoscaling Group",
		LongDesc:   "Use this command to retrieve a complete list of Servers from a specified VM Autoscaling Group provisioned under your account.\n\nRequired values to run command:\n\n* VM Autoscaling Group Id",
		Example:    listServerAutoscalingExample,
		PreCmdRun:  PreRunGroupId,
		CmdRun:     RunAutoscalingServerList,
		InitClient: true,
	})
	list.AddStringSliceFlag(config.ArgCols, "", defaultAutoscalingServerCols, utils.ColsMessage(defaultAutoscalingServerCols))
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultAutoscalingServerCols, cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringFlag(config.ArgGroupId, "", "", config.RequiredFlagGroupId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, autoscalingServerCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "server",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get a Server from a VM Autoscaling Group",
		LongDesc:   "Use this command to retrieve details about an Server from a specific VM Autoscaling Group by using its ID.\n\nRequired values to run command:\n\n* VM Autoscaling Group Id\n\n* Server Id",
		Example:    getServerAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingGroupServerIds,
		CmdRun:     RunAutoscalingServerGet,
		InitClient: true,
	})
	get.AddStringSliceFlag(config.ArgCols, "", defaultAutoscalingServerCols, utils.ColsMessage(defaultAutoscalingServerCols))
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultAutoscalingServerCols, cobra.ShellCompDirectiveNoFileComp
	})
	get.AddStringFlag(config.ArgGroupId, "", "", config.RequiredFlagGroupId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddStringFlag(config.ArgServerId, config.ArgIdShort, "", config.RequiredFlagServerId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgServerId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingServersIds(os.Stderr, viper.GetString(core.GetFlagName(get.NS, config.ArgGroupId))), cobra.ShellCompDirectiveNoFileComp
	})

	return autoscalingServerCmd
}

func PreRunAutoscalingGroupServerIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgGroupId, config.ArgServerId)
}

func RunAutoscalingServerList(c *core.CommandConfig) error {
	c.Printer.Verbose("Getting Servers for VM Autoscaling Group with ID: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	autoscalingServers, _, err := c.AutoscalingGroups().ListServers(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingServerPrint(c, getAutoscalingServers(autoscalingServers)))
}

func RunAutoscalingServerGet(c *core.CommandConfig) error {
	c.Printer.Verbose("Getting Server with ID: %v for VM Autoscaling Group with ID: %v", viper.GetString(core.GetFlagName(c.NS, config.ArgServerId)), viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	autoServer, _, err := c.AutoscalingGroups().GetServer(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)), viper.GetString(core.GetFlagName(c.NS, config.ArgServerId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingServerPrint(c, []sdkAutoscaling.Server{*autoServer}))
}

func getAutoscalingServers(servers sdkAutoscaling.Servers) []sdkAutoscaling.Server {
	autoscalingServers := make([]sdkAutoscaling.Server, 0)
	for _, a := range *servers.Items {
		autoscalingServers = append(autoscalingServers, sdkAutoscaling.Server{Server: a})
	}
	return autoscalingServers
}

// Output Printing

var defaultAutoscalingServerCols = []string{"ServerId", "DatacenterId", "Name"}

type AutoscalingServerPrint struct {
	ServerId     string `json:"ServerId,omitempty"`
	DatacenterId string `json:"DatacenterId,omitempty"`
	Name         string `json:"Name,omitempty"`
}

func getAutoscalingServerPrint(c *core.CommandConfig, dcs []sdkAutoscaling.Server) printer.Result {
	r := printer.Result{}
	if c != nil {
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getAutoscalingServersKVMaps(dcs)
			r.Columns = getAutoscalingServerCols(core.GetFlagName(c.NS, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getAutoscalingServerCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultAutoscalingServerCols
	}
	columnsMap := map[string]string{
		"ServerId":     "ServerId",
		"DatacenterId": "DatacenterId",
		"Name":         "Name",
	}
	var autoscalingServerCols []string
	for _, k := range cols {
		col := columnsMap[k]
		if col != "" {
			autoscalingServerCols = append(autoscalingServerCols, col)
		} else {
			clierror.CheckError(errors.New("unknown column "+k), outErr)
		}
	}
	return autoscalingServerCols
}

func getAutoscalingServersKVMaps(servers []sdkAutoscaling.Server) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(servers))
	for _, g := range servers {
		var autoscalingServerPrint AutoscalingServerPrint
		if idOk, ok := g.GetIdOk(); ok && idOk != nil {
			autoscalingServerPrint.ServerId = *idOk
		}
		if propertiesOk, ok := g.GetPropertiesOk(); ok && propertiesOk != nil {
			if datacenterServerOk, ok := propertiesOk.GetDatacenterServerOk(); ok && datacenterServerOk != nil {
				if idOk, ok := datacenterServerOk.GetIdOk(); ok && idOk != nil {
					autoscalingServerPrint.DatacenterId = *idOk
				}
			}
			if nameOk, ok := propertiesOk.GetNameOk(); ok && nameOk != nil {
				autoscalingServerPrint.Name = *nameOk
			}
		}
		o := structs.Map(autoscalingServerPrint)
		out = append(out, o)
	}
	return out
}

func getAutoscalingServersIds(outErr io.Writer, groupId string) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := sdkAutoscaling.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	autoscalingServerSvc := sdkAutoscaling.NewGroupService(clientSvc.Get(), context.TODO())
	autoscalingServers, _, err := autoscalingServerSvc.ListServers(groupId)
	clierror.CheckError(err, outErr)
	serverIds := make([]string, 0)
	if items, ok := autoscalingServers.ServerCollection.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				serverIds = append(serverIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return serverIds
}
