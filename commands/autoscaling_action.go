package commands

import (
	"context"
	"errors"
	ionoscloudAutoscaling "github.com/ionos-cloud/sdk-go-autoscaling"
	"io"
	"os"
	"strings"

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
			Short:            "VM Autoscaling Action Operations",
			Long:             "The sub-commands of `ionosctl autoscaling action` allow you to list, get Actions from a VM Autoscaling Group.",
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
		Namespace: "autoscaling",
		Resource:  "action",
		Verb:      "list",
		Aliases:   []string{"l", "ls"},
		ShortDesc: "List Actions from a VM Autoscaling Group",
		LongDesc: `Use this command to retrieve a complete list of Actions from a VM Autoscaling Group provisioned under your account.

Use flags to retrieve a list of Actions:

* sorting by type, using ` + "`" + `ionosctl autoscaling action list --group-id GROUP_ID --type ACTION_TYPE` + "`" + `
* sorting by status, using ` + "`" + `ionosctl autoscaling action list --group-id GROUP_ID --status ACTION_STATUS` + "`" + `

Required values to run command:

* VM Autoscaling Group Id`,
		Example:    listActionAutoscalingExample,
		PreCmdRun:  PreRunGroupId,
		CmdRun:     RunAutoscalingActionList,
		InitClient: true,
	})
	list.AddStringFlag(config.ArgGroupId, "", "", config.RequiredFlagGroupId)
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringFlag(config.ArgType, config.ArgTypeShort, "", "Sort Actions based on VM Autoscaling Action Type")
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgType, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"SCALE_IN", "SCALE_OUT"}, cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringFlag(config.ArgStatus, config.ArgStatusShort, "", "Sort Actions based on VM Autoscaling Action Status")
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgStatus, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"IN_PROGRESS", "SUCCESSFUL", "FAILED"}, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, autoscalingActionCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "action",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get an Action from a VM Autoscaling Group",
		LongDesc:   "Use this command to retrieve details about an Action from a VM Autoscaling Group by using its ID. You can wait for the Action to be in Successful state using `--wait-for-state` or `-W` option.\n\nRequired values to run command:\n\n* VM Autoscaling Group Id\n* Action Id",
		Example:    getActionAutoscalingExample,
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
	get.AddBoolFlag(config.ArgWaitForState, config.ArgWaitForStateShort, config.DefaultWait, "Wait for the Autoscaling Action to be SUCCESSFUL")
	get.AddIntFlag(config.ArgTimeout, config.ArgTimeoutShort, 600, "Timeout option for waiting for VM Autoscaling Action to be SUCCESSFUL [seconds]")

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
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgType)) {
		autoscalingActions = sortAutoscalingActionsByType(autoscalingActions, viper.GetString(core.GetFlagName(c.NS, config.ArgType)))
	}
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgStatus)) {
		autoscalingActions = sortAutoscalingActionsByStatus(autoscalingActions, viper.GetString(core.GetFlagName(c.NS, config.ArgStatus)))
	}
	return c.Printer.Print(getAutoscalingActionPrint(c, getAutoscalingActions(autoscalingActions)))
}

func RunAutoscalingActionGet(c *core.CommandConfig) error {
	if viper.GetBool(core.GetFlagName(c.NS, config.ArgWaitForState)) {
		if err := utils.WaitForState(c, GetStateAutoscalingAction, viper.GetString(core.GetFlagName(c.NS, config.ArgActionId))); err != nil {
			return err
		}
	}
	autoAction, _, err := c.AutoscalingGroups().GetAction(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)), viper.GetString(core.GetFlagName(c.NS, config.ArgActionId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingActionPrint(c, []sdkAutoscaling.Action{*autoAction}))
}

func GetStateAutoscalingAction(c *core.CommandConfig, objId string) (*string, error) {
	autoAction, _, err := c.AutoscalingGroups().GetAction(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)), objId)
	if err != nil {
		return nil, err
	}
	if metadata, ok := autoAction.GetPropertiesOk(); ok && metadata != nil {
		if state, ok := metadata.GetActionStatusOk(); ok && state != nil {
			return (*string)(state), nil
		}
	}
	return nil, nil
}

func getAutoscalingActions(actions sdkAutoscaling.Actions) []sdkAutoscaling.Action {
	autoscalingActions := make([]sdkAutoscaling.Action, 0)
	for _, a := range *actions.Items {
		autoscalingActions = append(autoscalingActions, sdkAutoscaling.Action{Action: a})
	}
	return autoscalingActions
}

// Sort VM Autoscaling Actions Based on Action Type
func sortAutoscalingActionsByType(actions sdkAutoscaling.Actions, actionType string) sdkAutoscaling.Actions {
	items := make([]ionoscloudAutoscaling.Action, 0)
	if itemsOk, ok := actions.GetItemsOk(); ok && itemsOk != nil {
		for _, item := range *itemsOk {
			if propertiesOk, ok := item.GetPropertiesOk(); ok && propertiesOk != nil {
				if actionTypeOk, ok := propertiesOk.GetActionTypeOk(); ok && actionTypeOk != nil {
					if string(*actionTypeOk) == strings.ToUpper(actionType) {
						items = append(items, item)
					}
				}
			}
		}
	}
	actions.Items = &items
	return actions
}

// Sort VM Autoscaling Actions Based on Action Status
func sortAutoscalingActionsByStatus(actions sdkAutoscaling.Actions, actionStatus string) sdkAutoscaling.Actions {
	items := make([]ionoscloudAutoscaling.Action, 0)
	if itemsOk, ok := actions.GetItemsOk(); ok && itemsOk != nil {
		for _, item := range *itemsOk {
			if propertiesOk, ok := item.GetPropertiesOk(); ok && propertiesOk != nil {
				if actionStatusOk, ok := propertiesOk.GetActionStatusOk(); ok && actionStatusOk != nil {
					if string(*actionStatusOk) == strings.ToUpper(actionStatus) {
						items = append(items, item)
					}
				}
			}
		}
	}
	actions.Items = &items
	return actions
}

// Output Printing

var defaultAutoscalingActionCols = []string{"ActionId", "ActionStatus", "ActionType", "TargetReplicaCount"}

type AutoscalingActionPrint struct {
	ActionId           string `json:"ActionId,omitempty"`
	ActionStatus       string `json:"ActionStatus,omitempty"`
	ActionType         string `json:"ActionType,omitempty"`
	TargetReplicaCount int64  `json:"TargetReplicaCount,omitempty"`
}

func getAutoscalingActionPrint(c *core.CommandConfig, dcs []sdkAutoscaling.Action) printer.Result {
	r := printer.Result{}
	if c != nil {
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
