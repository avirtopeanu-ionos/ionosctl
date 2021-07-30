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

func autoscalingAction() *core.Command {
	ctx := context.TODO()
	autoscalingActionCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "action",
			Aliases:          []string{"a"},
			Short:            "Autoscaling Action Operations",
			Long:             "The sub-commands of `ionosctl autoscaling action` allow you to list, get Actions from an Autoscaling Group.",
			TraverseChildren: true,
		},
	}
	globalFlags := autoscalingActionCmd.GlobalFlags()
	globalFlags.StringSliceP(config.ArgCols, "", defaultAutoscalingActionCols, utils.ColsMessage(defaultAutoscalingActionCols))
	_ = viper.BindPFlag(core.GetGlobalFlagName(autoscalingActionCmd.Name(), config.ArgCols), globalFlags.Lookup(config.ArgCols))
	_ = autoscalingActionCmd.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultAutoscalingActionCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	list := core.NewCommand(ctx, autoscalingActionCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "action",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Actions from an Autoscaling Group",
		LongDesc:   "Use this command to retrieve a complete list of Actions from an Autoscaling Group provisioned under your account.\n\nRequired values to run command:\n\n* Autoscaling Group Id",
		Example:    "",
		PreCmdRun:  PreRunGroupId,
		CmdRun:     RunAutoscalingActionList,
		InitClient: true,
	})
	list.AddStringFlag(config.ArgGroupId, "", "", config.RequiredFlagGroupId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, autoscalingActionCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "action",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get an Action from an Autoscaling Group",
		LongDesc:   "Use this command to retrieve details about an Action from an Autoscaling Group by using its ID.\n\nRequired values to run command:\n\n* Autoscaling Group Id\n* Action Id",
		Example:    "",
		PreCmdRun:  PreRunAutoscalingGroupActionIds,
		CmdRun:     RunAutoscalingActionGet,
		InitClient: true,
	})
	get.AddStringFlag(config.ArgGroupId, "", "", config.RequiredFlagGroupId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddStringFlag(config.ArgActionId, config.ArgIdShort, "", config.RequiredFlagActionId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgActionId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingActionsIds(os.Stderr, viper.GetString(core.GetFlagName(get.NS, config.ArgGroupId))), cobra.ShellCompDirectiveNoFileComp
	})

	return autoscalingActionCmd
}

func PreRunAutoscalingGroupActionIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgGroupId, config.ArgActionId)
}

func RunAutoscalingActionList(c *core.CommandConfig) error {
	autoscalingActions, _, err := c.AutoscalingGroups().ListActions(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingActionPrint(nil, c, getAutoscalingActions(autoscalingActions)))
}

func RunAutoscalingActionGet(c *core.CommandConfig) error {
	autoAction, _, err := c.AutoscalingGroups().GetAction(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)), viper.GetString(core.GetFlagName(c.NS, config.ArgActionId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingActionPrint(nil, c, []sdkAutoscaling.Action{*autoAction}))
}

func getAutoscalingActions(actions sdkAutoscaling.Actions) []sdkAutoscaling.Action {
	autoscalingActions := make([]sdkAutoscaling.Action, 0)
	for _, a := range *actions.Items {
		autoscalingActions = append(autoscalingActions, sdkAutoscaling.Action{Action: a})
	}
	return autoscalingActions
}

// Output Printing

var defaultAutoscalingActionCols = []string{"ActionId", "ActionStatus", "ActionType", "TargetReplicaCount"}

type AutoscalingActionPrint struct {
	ActionId           string `json:"ActionId,omitempty"`
	ActionStatus       string `json:"ActionStatus,omitempty"`
	ActionType         string `json:"ActionType,omitempty"`
	TargetReplicaCount int64  `json:"TargetReplicaCount,omitempty"`
}

func getAutoscalingActionPrint(resp *sdkAutoscaling.Response, c *core.CommandConfig, dcs []sdkAutoscaling.Action) printer.Result {
	r := printer.Result{}
	if c != nil {
		if resp != nil {
			r.Resource = c.Resource
			r.Verb = c.Verb
		}
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getAutoscalingActionsKVMaps(dcs)
			r.Columns = getAutoscalingActionCols(core.GetGlobalFlagName(c.Resource, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getAutoscalingActionCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultAutoscalingActionCols
	}
	columnsMap := map[string]string{
		"ActionId":           "ActionId",
		"ActionStatus":       "ActionStatus",
		"ActionType":         "ActionType",
		"TargetReplicaCount": "TargetReplicaCount",
	}
	var autoscalingActionCols []string
	for _, k := range cols {
		col := columnsMap[k]
		if col != "" {
			autoscalingActionCols = append(autoscalingActionCols, col)
		} else {
			clierror.CheckError(errors.New("unknown column "+k), outErr)
		}
	}
	return autoscalingActionCols
}

func getAutoscalingActionsKVMaps(actions []sdkAutoscaling.Action) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(actions))
	for _, g := range actions {
		var autoscalingActionPrint AutoscalingActionPrint
		if idOk, ok := g.GetIdOk(); ok && idOk != nil {
			autoscalingActionPrint.ActionId = *idOk
		}
		if properties, ok := g.GetPropertiesOk(); ok && properties != nil {
			if actionTypeOk, ok := properties.GetActionTypeOk(); ok && actionTypeOk != nil {
				autoscalingActionPrint.ActionType = string(*actionTypeOk)
			}
			if actionStatusOk, ok := properties.GetActionStatusOk(); ok && actionStatusOk != nil {
				autoscalingActionPrint.ActionStatus = string(*actionStatusOk)
			}
			if targetReplicaCountOk, ok := properties.GetTargetReplicaCountOk(); ok && targetReplicaCountOk != nil {
				autoscalingActionPrint.TargetReplicaCount = *targetReplicaCountOk
			}
		}
		o := structs.Map(autoscalingActionPrint)
		out = append(out, o)
	}
	return out
}

func getAutoscalingActionsIds(outErr io.Writer, groupId string) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := sdkAutoscaling.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	autoscalingActionSvc := sdkAutoscaling.NewGroupService(clientSvc.Get(), context.TODO())
	autoscalingActions, _, err := autoscalingActionSvc.ListActions(groupId)
	clierror.CheckError(err, outErr)
	actionIds := make([]string, 0)
	if items, ok := autoscalingActions.ActionCollection.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				actionIds = append(actionIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return actionIds
}
