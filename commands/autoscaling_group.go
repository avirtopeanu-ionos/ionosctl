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

func autoscalingGroup() *core.Command {
	ctx := context.TODO()
	autoscalingGroupCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "group",
			Aliases:          []string{"g"},
			Short:            "Autoscaling Group Operations",
			Long:             "The sub-commands of `ionosctl autoscaling group` allow you to create, list, get, update and delete Autoscaling Groups.",
			TraverseChildren: true,
		},
	}

	/*
		List Command
	*/
	list := core.NewCommand(ctx, autoscalingGroupCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "group",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Autoscaling Groups",
		LongDesc:   "Use this command to retrieve a complete list of Autoscaling Groups provisioned under your account.",
		Example:    listGroupAutoscalingExample,
		PreCmdRun:  noPreRun,
		CmdRun:     RunAutoscalingGroupList,
		InitClient: true,
	})
	list.AddStringSliceFlag(config.ArgCols, "", defaultAutoscalingGroupCols, utils.ColsMessage(allAutoscalingGroupCols))
	_ = list.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allAutoscalingGroupCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, autoscalingGroupCmd, core.CommandBuilder{
		Namespace:  "autoscaling",
		Resource:   "group",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get an Autoscaling Group",
		LongDesc:   "Use this command to retrieve details about an Autoscaling Group by using its ID.\n\nRequired values to run command:\n\n* Autoscaling Group Id",
		Example:    getGroupAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingGroupId,
		CmdRun:     RunAutoscalingGroupGet,
		InitClient: true,
	})
	get.AddStringFlag(config.ArgGroupId, config.ArgIdShort, "", config.RequiredFlagGroupId)
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddStringSliceFlag(config.ArgCols, "", defaultAutoscalingGroupCols, utils.ColsMessage(allAutoscalingGroupCols))
	_ = get.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allAutoscalingGroupCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Create Command
	*/
	create := core.NewCommand(ctx, autoscalingGroupCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "group",
		Verb:      "create",
		Aliases:   []string{"c"},
		ShortDesc: "Create an Autoscaling Group",
		LongDesc: `Create an Autoscaling Group. 

Regarding some of the Autoscaling Group Properties, please see more details:

* [Group][DatacenterId]The Datacenter Id property represents VMs for this Autoscaling Group will be created in this Virtual Datacenter. Please note, that it have the same location as the template.
* [Group][TargetReplicaCount] Depending on the scaling policy, the target number of VMs will be adjusted automatically. VMs will be created or destroyed automatically in order to adjust the actual number of VMs to this number. This value can be set only at Group creation time, subsequent change via update (PUT) request is not possible
* [Group Policy][ScaleInThreshold] Scale In Threshold is a lower threshold on the value of ` + "`" + `metric` + "`" + `. Will be used with ` + "`" + `less than` + "`" + ` (<) operator. Exceeding this will start a Scale-In Action as specified by the ` + "`" + `scaleInAction` + "`" + ` property. The value must have a higher minimum delta to the ` + "`" + `scaleOutThreshold` + "`" + ` depending on the ` + "`" + `metric` + "`" + ` to avoid competitive actions at the same time
* [Group Policy][ScaleOutThreshold] An upper threshold on the value of ` + "`" + `metric` + "`" + `.  Will be used with ` + "`" + `greater than` + "`" + ` (>) operator. Exceeding this will start a Scale-Out Action as specified by the ` + "`" + `scaleOutAction` + "`" + ` property. The value must have a lower minimum delta to the ` + "`" + `scaleInThreshold` + "`" + ` depending on the ` + "`" + `metric` + "`" + ` to avoid competitive actions at the same time
* [Group Policy Action][Amount] When ` + "`" + `amountType == ABSOLUTE` + "`" + `, amount parameter is the number of VMs added or removed in one step. When ` + "`" + `amountType == PERCENTAGE` + "`" + `, amount parameter is a percentage value, which will be applied to the Autoscaling Group's current ` + "`" + `targetReplicaCount` + "`" + ` in order to derive the number of VMs that will be added or removed in one step. There will always be at least one VM added or removed
* [Group Policy Action][CoolDownPeriod] Minimum time to pass after this Scaling Action has started, until the next Scaling Action will be started. Additionally, if a Scaling Action is currently in progress, no second Scaling Action will be started for the same Autoscaling Group. Instead, the Metric will be re-evaluated after the current Scaling Action completed (either successful or with failures)

Required values to run command:

* Autoscaling Template Id
* Datacenter Id`,
		Example:    createGroupAutoscalingExample,
		PreCmdRun:  PreRunDatacenterAutoscalingTemplateIds,
		CmdRun:     RunAutoscalingGroupCreate,
		InitClient: true,
	})
	create.AddStringSliceFlag(config.ArgCols, "", defaultAutoscalingGroupCols, utils.ColsMessage(allAutoscalingGroupCols))
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allAutoscalingGroupCols, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgName, config.ArgNameShort, "Unnamed Autoscaling Group", "User-defined name for the Autoscaling Group")
	create.AddStringFlag(config.ArgDataCenterId, "", "", config.RequiredFlagDatacenterId)
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getDataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	create.AddIntFlag(config.ArgMaxReplicaCount, "", 5, "Maximum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200")
	create.AddIntFlag(config.ArgMinReplicaCount, "", 1, "Minimum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200")
	create.AddIntFlag(config.ArgTargetReplicaCount, config.ArgTargetReplicaCountShort, 1, "The target number of VMs in this Group. Minimum: 0; Maximum: 200")
	create.AddStringFlag(config.ArgTemplateId, config.ArgIdShort, "", config.RequiredFlagTemplateId)
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	// Group Policy
	create.AddStringFlag(config.ArgMetric, config.ArgMetricShort, "INSTANCE_CPU_UTILIZATION_AVERAGE", "[Group Policy] The Metric that should trigger Scaling Actions. The values of the Metric are checked in fixed intervals")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"INSTANCE_CPU_UTILIZATION_AVERAGE", "INSTANCE_NETWORK_IN_BYTES", "INSTANCE_NETWORK_IN_PACKETS", "INSTANCE_NETWORK_OUT_BYTES", "INSTANCE_NETWORK_OUT_PACKETS"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgRange, config.ArgRangeShort, "PT24H", "[Group Policy] Defines the range of time from which samples will be aggregated")
	create.AddStringFlag(config.ArgUnit, "", "PER_HOUR", "[Group Policy] Unit of the applied Metric")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"PER_HOUR", "PER_MINUTE", "PER_SECOND", "TOTAL"}, cobra.ShellCompDirectiveNoFileComp
	})
	// Scale In Action Policy
	create.AddFloat32Flag(config.ArgScaleInThreshold, "", 33, "[Group Policy][Scale In Action] A lower threshold on the value of `metric`")
	create.AddFloat32Flag(config.ArgScaleInAmount, "", 1, "[Group Policy][Scale In Action] Amount of VMs (in percentage or absolute value) to be removed in a Scale In Action")
	create.AddStringFlag(config.ArgScaleInAmountType, "", "ABSOLUTE", "[Group Policy][Scale In Action] The type for the given amount")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ABSOLUTE", "PERCENTAGE"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgScaleInCoolDownPeriod, "", "5m", "[Group Policy][Scale In Action] Cool Down Period")
	// Scale Out Action Policy
	create.AddFloat32Flag(config.ArgScaleOutThreshold, "", 77, "[Group Policy][Scale Out Action] An upper threshold on the value of `metric`")
	create.AddFloat32Flag(config.ArgScaleOutAmount, "", 1, "[Group Policy][Scale Out Action] Amount of VMs (in percentage or absolute value) to be added in a Scale Out Action")
	create.AddStringFlag(config.ArgScaleOutAmountType, "", "ABSOLUTE", "[Group Policy][Scale Out Action] The type for the given amount")
	_ = create.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ABSOLUTE", "PERCENTAGE"}, cobra.ShellCompDirectiveNoFileComp
	})
	create.AddStringFlag(config.ArgScaleOutCoolDownPeriod, "", "5m", "[Group Policy][Scale Out Action] Cool Down Period")

	/*
		Update Command
	*/
	update := core.NewCommand(ctx, autoscalingGroupCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "group",
		Verb:      "update",
		Aliases:   []string{"up"},
		ShortDesc: "Update an Autoscaling Group",
		LongDesc: `Update an Autoscaling Group. 

Required values to run command:

* Autoscaling Group Id`,
		Example:    updateGroupAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingGroupId,
		CmdRun:     RunAutoscalingGroupUpdate,
		InitClient: true,
	})
	update.AddStringFlag(config.ArgGroupId, config.ArgIdShort, "", config.RequiredFlagGroupId)
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	update.AddStringFlag(config.ArgName, config.ArgNameShort, "", "User-defined name for the Autoscaling Group")
	update.AddIntFlag(config.ArgMaxReplicaCount, "", 0, "Maximum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200")
	update.AddIntFlag(config.ArgMinReplicaCount, "", 0, "Minimum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200")
	update.AddStringFlag(config.ArgTemplateId, "", "", "The unique Template Id")
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgTemplateId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingTemplatesIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	// Group Policy
	update.AddStringFlag(config.ArgMetric, config.ArgMetricShort, "", "[Group Policy] The Metric that should trigger Scaling Actions. The values of the Metric are checked in fixed intervals")
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"INSTANCE_CPU_UTILIZATION_AVERAGE", "INSTANCE_NETWORK_IN_BYTES", "INSTANCE_NETWORK_IN_PACKETS", "INSTANCE_NETWORK_OUT_BYTES", "INSTANCE_NETWORK_OUT_PACKETS"}, cobra.ShellCompDirectiveNoFileComp
	})
	update.AddStringFlag(config.ArgRange, config.ArgRangeShort, "", "[Group Policy] Defines the range of time from which samples will be aggregated")
	update.AddStringFlag(config.ArgUnit, "", "", "[Group Policy] Unit of the applied Metric")
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"PER_HOUR", "PER_MINUTE", "PER_SECOND", "TOTAL"}, cobra.ShellCompDirectiveNoFileComp
	})
	// Scale In Action Policy
	update.AddFloat32Flag(config.ArgScaleInThreshold, "", 0, "[Group Policy][Scale In Action] A lower threshold on the value of `metric`")
	update.AddFloat32Flag(config.ArgScaleInAmount, "", 0, "[Group Policy][Scale In Action] Amount of VMs (in percentage or absolute value) to be removed in a Scale In Action")
	update.AddStringFlag(config.ArgScaleInAmountType, "", "", "[Group Policy][Scale In Action] The type for the given amount")
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ABSOLUTE", "PERCENTAGE"}, cobra.ShellCompDirectiveNoFileComp
	})
	update.AddStringFlag(config.ArgScaleInCoolDownPeriod, "", "", "[Group Policy][Scale In Action] Cool Down Period")
	// Scale Out Action Policy
	update.AddFloat32Flag(config.ArgScaleOutThreshold, "", 40, "[Group Policy][Scale Out Action] An upper threshold on the value of `metric`")
	update.AddFloat32Flag(config.ArgScaleOutAmount, "", 0, "[Group Policy][Scale Out Action] Amount of VMs (in percentage or absolute value) to be added in a Scale Out Action")
	update.AddStringFlag(config.ArgScaleOutAmountType, "", "", "[Group Policy][Scale Out Action] The type for the given amount")
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgMetric, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ABSOLUTE", "PERCENTAGE"}, cobra.ShellCompDirectiveNoFileComp
	})
	update.AddStringFlag(config.ArgScaleOutCoolDownPeriod, "", "", "[Group Policy][Scale Out Action] Cool Down Period")
	update.AddStringSliceFlag(config.ArgCols, "", defaultAutoscalingGroupCols, utils.ColsMessage(allAutoscalingGroupCols))
	_ = update.Command.RegisterFlagCompletionFunc(config.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allAutoscalingGroupCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		Delete Command
	*/
	deleteCmd := core.NewCommand(ctx, autoscalingGroupCmd, core.CommandBuilder{
		Namespace: "autoscaling",
		Resource:  "group",
		Verb:      "delete",
		Aliases:   []string{"d"},
		ShortDesc: "Delete an Autoscaling Group",
		LongDesc: `Use this command to delete a specified Autoscaling Group from your account.

Required values to run command:

* Autoscaling Group Id`,
		Example:    deleteGroupAutoscalingExample,
		PreCmdRun:  PreRunAutoscalingGroupId,
		CmdRun:     RunAutoscalingGroupDelete,
		InitClient: true,
	})
	deleteCmd.AddStringFlag(config.ArgGroupId, config.ArgIdShort, "", config.RequiredFlagGroupId)
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(config.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getAutoscalingGroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})

	return autoscalingGroupCmd
}

func PreRunDatacenterAutoscalingTemplateIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgDataCenterId, config.ArgTemplateId)
}

func PreRunAutoscalingGroupId(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.NS, config.ArgGroupId)
}

func RunAutoscalingGroupList(c *core.CommandConfig) error {
	autoscalingGroups, _, err := c.AutoscalingGroups().List()
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingGroupPrint(nil, c, getAutoscalingGroups(autoscalingGroups)))
}

func RunAutoscalingGroupGet(c *core.CommandConfig) error {
	autoGroup, _, err := c.AutoscalingGroups().Get(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingGroupPrint(nil, c, []sdkAutoscaling.Group{*autoGroup}))
}

func RunAutoscalingGroupCreate(c *core.CommandConfig) error {
	groupProperties, err := getNewAutoscalingGroup(c)
	if err != nil {
		return err
	}
	dc, resp, err := c.AutoscalingGroups().Create(sdkAutoscaling.Group{
		Group: ionoscloudAutoscaling.Group{
			Properties: &groupProperties.GroupProperties,
		},
	})
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingGroupPrint(resp, c, []sdkAutoscaling.Group{*dc}))
}

func RunAutoscalingGroupUpdate(c *core.CommandConfig) error {
	autoGroup, _, err := c.AutoscalingGroups().Get(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	if err != nil {
		return err
	}
	groupProperties, err := getUpdateAutoscalingGroup(c, autoGroup)
	if err != nil {
		return err
	}
	dc, resp, err := c.AutoscalingGroups().Update(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)), sdkAutoscaling.Group{
		Group: ionoscloudAutoscaling.Group{
			Properties: &groupProperties.GroupProperties,
		},
	})
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingGroupPrint(resp, c, []sdkAutoscaling.Group{*dc}))
}

func RunAutoscalingGroupDelete(c *core.CommandConfig) error {
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete autoscaling group"); err != nil {
		return err
	}
	resp, err := c.AutoscalingGroups().Delete(viper.GetString(core.GetFlagName(c.NS, config.ArgGroupId)))
	if err != nil {
		return err
	}
	return c.Printer.Print(getAutoscalingGroupPrint(resp, c, nil))
}

func getAutoscalingGroups(groups sdkAutoscaling.Groups) []sdkAutoscaling.Group {
	tpls := make([]sdkAutoscaling.Group, 0)
	for _, tpl := range *groups.Items {
		tpls = append(tpls, sdkAutoscaling.Group{Group: tpl})
	}
	return tpls
}

func getNewAutoscalingGroup(c *core.CommandConfig) (*sdkAutoscaling.GroupProperties, error) {
	input := ionoscloudAutoscaling.GroupProperties{}
	// Autoscaling Group - VM Properties
	input.SetName(viper.GetString(core.GetFlagName(c.NS, config.ArgName)))
	input.SetMaxReplicaCount(viper.GetInt64(core.GetFlagName(c.NS, config.ArgMaxReplicaCount)))
	input.SetMinReplicaCount(viper.GetInt64(core.GetFlagName(c.NS, config.ArgMinReplicaCount)))
	input.SetTargetReplicaCount(viper.GetInt64(core.GetFlagName(c.NS, config.ArgTargetReplicaCount)))

	// Set Group Datacenter where the VMs will be created
	// Note: MAKE SURE THE DATACENTER HAS THE SAME LOCATION AS TEMPLATE
	datacenterId := viper.GetString(core.GetFlagName(c.NS, config.ArgDataCenterId))
	inputDatacenter := ionoscloudAutoscaling.GroupPropertiesDatacenter{}
	inputDatacenter.SetId(datacenterId)
	input.SetDatacenter(inputDatacenter)

	// Set Group Template to be used for the VMs creation
	templateId := viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId))
	inputTemplate := ionoscloudAutoscaling.GroupPropertiesTemplate{}
	inputTemplate.SetId(templateId)
	input.SetTemplate(inputTemplate)

	// Set Group Policy
	inputPolicies := ionoscloudAutoscaling.GroupPolicy{}
	inputPolicies.SetMetric(ionoscloudAutoscaling.Metric(viper.GetString(core.GetFlagName(c.NS, config.ArgMetric))))
	inputPolicies.SetRange(viper.GetString(core.GetFlagName(c.NS, config.ArgRange)))
	inputPolicies.SetUnit(ionoscloudAutoscaling.QueryUnit(viper.GetString(core.GetFlagName(c.NS, config.ArgUnit))))
	inputPolicies.SetScaleInThreshold(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleInThreshold))))
	inputPolicies.SetScaleOutThreshold(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleOutThreshold))))

	// Set Group Action Policy
	// SCALE IN Action
	inputScaleInAction := ionoscloudAutoscaling.GroupPolicyAction{}
	inputScaleInAction.SetAmount(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleInAmount))))
	inputScaleInAction.SetAmountType(ionoscloudAutoscaling.ActionAmount(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleInAmountType))))
	inputScaleInAction.SetCooldownPeriod(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleInCoolDownPeriod)))
	inputPolicies.SetScaleInAction(inputScaleInAction)
	// SCALE OUT Action
	inputScaleOutAction := ionoscloudAutoscaling.GroupPolicyAction{}
	inputScaleOutAction.SetAmount(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleOutAmount))))
	inputScaleOutAction.SetAmountType(ionoscloudAutoscaling.ActionAmount(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleOutAmountType))))
	inputScaleOutAction.SetCooldownPeriod(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleOutCoolDownPeriod)))
	inputPolicies.SetScaleOutAction(inputScaleOutAction)

	// Set Group Policy (required)
	input.SetPolicy(inputPolicies)

	return &sdkAutoscaling.GroupProperties{
		GroupProperties: input,
	}, nil
}

func getUpdateAutoscalingGroup(c *core.CommandConfig, autoGroup *sdkAutoscaling.Group) (*sdkAutoscaling.GroupProperties, error) {
	// Get actual group propertiesOk
	if propertiesOk, ok := autoGroup.GetPropertiesOk(); ok && propertiesOk != nil {
		// Autoscaling Group - Update VM Properties
		if viper.IsSet(core.GetFlagName(c.NS, config.ArgName)) {
			propertiesOk.SetName(viper.GetString(core.GetFlagName(c.NS, config.ArgName)))
		}
		if viper.IsSet(core.GetFlagName(c.NS, config.ArgMaxReplicaCount)) {
			propertiesOk.SetMaxReplicaCount(viper.GetInt64(core.GetFlagName(c.NS, config.ArgMaxReplicaCount)))
		}
		if viper.IsSet(core.GetFlagName(c.NS, config.ArgMinReplicaCount)) {
			propertiesOk.SetMinReplicaCount(viper.GetInt64(core.GetFlagName(c.NS, config.ArgMinReplicaCount)))
		}
		// Set Group Template to be used for the VMs creation
		if viper.IsSet(core.GetFlagName(c.NS, config.ArgTemplateId)) {
			templateId := viper.GetString(core.GetFlagName(c.NS, config.ArgTemplateId))
			inputTemplate := ionoscloudAutoscaling.GroupPropertiesTemplate{}
			inputTemplate.SetId(templateId)
			propertiesOk.SetTemplate(inputTemplate)
		}
		// Set Group Policy
		if policyOk, ok := propertiesOk.GetPolicyOk(); ok && policyOk != nil {
			if viper.IsSet(core.GetFlagName(c.NS, config.ArgMetric)) {
				policyOk.SetMetric(ionoscloudAutoscaling.Metric(viper.GetString(core.GetFlagName(c.NS, config.ArgMetric))))
			}
			if viper.IsSet(core.GetFlagName(c.NS, config.ArgRange)) {
				policyOk.SetRange(viper.GetString(core.GetFlagName(c.NS, config.ArgRange)))
			}
			if viper.IsSet(core.GetFlagName(c.NS, config.ArgUnit)) {
				policyOk.SetUnit(ionoscloudAutoscaling.QueryUnit(viper.GetString(core.GetFlagName(c.NS, config.ArgUnit))))
			}
			if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleInThreshold)) {
				policyOk.SetScaleInThreshold(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleInThreshold))))
			}
			if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleOutThreshold)) {
				policyOk.SetScaleOutThreshold(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleOutThreshold))))
			}

			// Set Group Action Policy
			// SCALE IN Action
			if scaleInActionOk, ok := policyOk.GetScaleInActionOk(); ok && scaleInActionOk != nil {
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleInAmount)) {
					scaleInActionOk.SetAmount(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleInAmount))))
				}
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleInAmountType)) {
					scaleInActionOk.SetAmountType(ionoscloudAutoscaling.ActionAmount(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleInAmountType))))
				}
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleInCoolDownPeriod)) {
					scaleInActionOk.SetCooldownPeriod(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleInCoolDownPeriod)))
				}
				policyOk.SetScaleInAction(*scaleInActionOk)
			}
			// SCALE OUT Action
			if scaleOutActionOk, ok := policyOk.GetScaleOutActionOk(); ok && scaleOutActionOk != nil {
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleOutAmount)) {
					scaleOutActionOk.SetAmount(float32(viper.GetFloat64(core.GetFlagName(c.NS, config.ArgScaleOutAmount))))
				}
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleOutAmountType)) {
					scaleOutActionOk.SetAmountType(ionoscloudAutoscaling.ActionAmount(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleOutAmountType))))
				}
				if viper.IsSet(core.GetFlagName(c.NS, config.ArgScaleOutCoolDownPeriod)) {
					scaleOutActionOk.SetCooldownPeriod(viper.GetString(core.GetFlagName(c.NS, config.ArgScaleOutCoolDownPeriod)))
				}
				policyOk.SetScaleOutAction(*scaleOutActionOk)
			}

			// Set Group Policy (required)
			propertiesOk.SetPolicy(*policyOk)
		}
		return &sdkAutoscaling.GroupProperties{
			GroupProperties: *propertiesOk,
		}, nil
	} else {
		return nil, errors.New("error")
	}
}

// Output Printing

var (
	defaultAutoscalingGroupCols = []string{"GroupId", "Name", "DatacenterId", "Location", "TemplateId", "TargetReplicaCount", "State"}
	allAutoscalingGroupCols     = []string{"GroupId", "Name", "DatacenterId", "Location", "TemplateId", "MaxReplicaCount", "MinReplicaCount", "TargetReplicaCount", "Metric", "Range", "Unit",
		"ScaleInThreshold", "ScaleInAmount", "ScaleInAmountType", "ScaleInCoolDownPeriod", "ScaleOutThreshold", "ScaleOutAmount", "ScaleOutAmountType", "ScaleOutCoolDownPeriod", "State"}
)

type AutoscalingGroupPrint struct {
	GroupId                string  `json:"GroupId,omitempty"`
	Name                   string  `json:"Name,omitempty"`
	DatacenterId           string  `json:"DatacenterId,omitempty"`
	Location               string  `json:"Location,omitempty"`
	TemplateId             string  `json:"TemplateId,omitempty"`
	MaxReplicaCount        int64   `json:"MaxReplicaCount,omitempty"`
	MinReplicaCount        int64   `json:"MinReplicaCount,omitempty"`
	TargetReplicaCount     int64   `json:"TargetReplicaCount,omitempty"`
	Metric                 string  `json:"Metric,omitempty"`
	Range                  string  `json:"Range,omitempty"`
	Unit                   string  `json:"Unit,omitempty"`
	ScaleInThreshold       float32 `json:"ScaleInThreshold,omitempty"`
	ScaleInAmount          float32 `json:"ScaleInAmount,omitempty"`
	ScaleInAmountType      string  `json:"ScaleInAmountType,omitempty"`
	ScaleInCoolDownPeriod  string  `json:"ScaleInCoolDownPeriod,omitempty"`
	ScaleOutThreshold      float32 `json:"ScaleOutThreshold,omitempty"`
	ScaleOutAmount         float32 `json:"ScaleOutAmount,omitempty"`
	ScaleOutAmountType     string  `json:"ScaleOutAmountType,omitempty"`
	ScaleOutCoolDownPeriod string  `json:"ScaleOutCoolDownPeriod,omitempty"`
	State                  string  `json:"State,omitempty"`
}

func getAutoscalingGroupPrint(resp *sdkAutoscaling.Response, c *core.CommandConfig, dcs []sdkAutoscaling.Group) printer.Result {
	r := printer.Result{}
	if c != nil {
		if resp != nil {
			r.Resource = c.Resource
			r.Verb = c.Verb
		}
		if dcs != nil {
			r.OutputJSON = dcs
			r.KeyValue = getAutoscalingGroupsKVMaps(dcs)
			r.Columns = getAutoscalingGroupCols(core.GetFlagName(c.NS, config.ArgCols), c.Printer.GetStderr())
		}
	}
	return r
}

func getAutoscalingGroupCols(flagName string, outErr io.Writer) []string {
	var cols []string
	if viper.IsSet(flagName) {
		cols = viper.GetStringSlice(flagName)
	} else {
		return defaultAutoscalingGroupCols
	}
	columnsMap := map[string]string{
		"GroupId":                "GroupId",
		"Name":                   "Name",
		"DatacenterId":           "DatacenterId",
		"Location":               "Location",
		"TemplateId":             "TemplateId",
		"MaxReplicaCount":        "MaxReplicaCount",
		"MinReplicaCount":        "MinReplicaCount",
		"TargetReplicaCount":     "TargetReplicaCount",
		"Metric":                 "Metric",
		"Range":                  "Range",
		"Unit":                   "Unit",
		"ScaleInThreshold":       "ScaleInThreshold",
		"ScaleInAmount":          "ScaleInAmount",
		"ScaleInAmountType":      "ScaleInAmountType",
		"ScaleInCoolDownPeriod":  "ScaleInCoolDownPeriod",
		"ScaleOutThreshold":      "ScaleOutThreshold",
		"ScaleOutAmount":         "ScaleOutAmount",
		"ScaleOutAmountType":     "ScaleOutAmountType",
		"ScaleOutCoolDownPeriod": "ScaleOutCoolDownPeriod",
		"State":                  "State",
	}
	var autoscalingGroupCols []string
	for _, k := range cols {
		col := columnsMap[k]
		if col != "" {
			autoscalingGroupCols = append(autoscalingGroupCols, col)
		} else {
			clierror.CheckError(errors.New("unknown column "+k), outErr)
		}
	}
	return autoscalingGroupCols
}

func getAutoscalingGroupsKVMaps(groups []sdkAutoscaling.Group) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(groups))
	for _, g := range groups {
		var autoscalingGroupPrint AutoscalingGroupPrint
		if idOk, ok := g.GetIdOk(); ok && idOk != nil {
			autoscalingGroupPrint.GroupId = *idOk
		}
		if properties, ok := g.GetPropertiesOk(); ok && properties != nil {
			if nameOk, ok := properties.GetNameOk(); ok && nameOk != nil {
				autoscalingGroupPrint.Name = *nameOk
			}
			if datacenterOk, ok := properties.GetDatacenterOk(); ok && datacenterOk != nil {
				if idOk, ok := datacenterOk.GetIdOk(); ok && idOk != nil {
					autoscalingGroupPrint.DatacenterId = *idOk
				}
			}
			if locationOk, ok := properties.GetLocationOk(); ok && locationOk != nil {
				autoscalingGroupPrint.Location = *locationOk
			}
			if templateOk, ok := properties.GetTemplateOk(); ok && templateOk != nil {
				if idOk, ok := templateOk.GetIdOk(); ok && idOk != nil {
					autoscalingGroupPrint.TemplateId = *idOk
				}
			}
			if maxReplicaCountOk, ok := properties.GetMaxReplicaCountOk(); ok && maxReplicaCountOk != nil {
				autoscalingGroupPrint.MaxReplicaCount = *maxReplicaCountOk
			}
			if minReplicaCountOk, ok := properties.GetMinReplicaCountOk(); ok && minReplicaCountOk != nil {
				autoscalingGroupPrint.MinReplicaCount = *minReplicaCountOk
			}
			if targetReplicaCountOk, ok := properties.GetTargetReplicaCountOk(); ok && targetReplicaCountOk != nil {
				autoscalingGroupPrint.TargetReplicaCount = *targetReplicaCountOk
			}
			if policyOk, ok := properties.GetPolicyOk(); ok && policyOk != nil {
				if metricOk, ok := policyOk.GetMetricOk(); ok && metricOk != nil {
					autoscalingGroupPrint.Metric = string(*metricOk)
				}
				if rangeOk, ok := policyOk.GetRangeOk(); ok && rangeOk != nil {
					autoscalingGroupPrint.Range = *rangeOk
				}
				if unitOk, ok := policyOk.GetUnitOk(); ok && unitOk != nil {
					autoscalingGroupPrint.Unit = string(*unitOk)
				}
				if scaleInThresholdOk, ok := policyOk.GetScaleInThresholdOk(); ok && scaleInThresholdOk != nil {
					autoscalingGroupPrint.ScaleInThreshold = *scaleInThresholdOk
				}
				if scaleOutThresholdOk, ok := policyOk.GetScaleOutThresholdOk(); ok && scaleOutThresholdOk != nil {
					autoscalingGroupPrint.ScaleOutThreshold = *scaleOutThresholdOk
				}
				if scaleInActionOk, ok := policyOk.GetScaleInActionOk(); ok && scaleInActionOk != nil {
					if amountOk, ok := scaleInActionOk.GetAmountOk(); ok && amountOk != nil {
						autoscalingGroupPrint.ScaleInAmount = *amountOk
					}
					if amountTypeOk, ok := scaleInActionOk.GetAmountTypeOk(); ok && amountTypeOk != nil {
						autoscalingGroupPrint.ScaleInAmountType = string(*amountTypeOk)
					}
					if cooldownPeriodOk, ok := scaleInActionOk.GetCooldownPeriodOk(); ok && cooldownPeriodOk != nil {
						autoscalingGroupPrint.ScaleInCoolDownPeriod = *cooldownPeriodOk
					}
				}
				if scaleOutActionOk, ok := policyOk.GetScaleOutActionOk(); ok && scaleOutActionOk != nil {
					if amountOk, ok := scaleOutActionOk.GetAmountOk(); ok && amountOk != nil {
						autoscalingGroupPrint.ScaleOutAmount = *amountOk
					}
					if amountTypeOk, ok := scaleOutActionOk.GetAmountTypeOk(); ok && amountTypeOk != nil {
						autoscalingGroupPrint.ScaleOutAmountType = string(*amountTypeOk)
					}
					if cooldownPeriodOk, ok := scaleOutActionOk.GetCooldownPeriodOk(); ok && cooldownPeriodOk != nil {
						autoscalingGroupPrint.ScaleOutCoolDownPeriod = *cooldownPeriodOk
					}
				}
			}
		}
		if metadataOk, ok := g.GetMetadataOk(); ok && metadataOk != nil {
			if stateOk, ok := metadataOk.GetStateOk(); ok && stateOk != nil {
				autoscalingGroupPrint.State = string(*stateOk)
			}
		}
		o := structs.Map(autoscalingGroupPrint)
		out = append(out, o)
	}
	return out
}

func getAutoscalingGroupsIds(outErr io.Writer) []string {
	err := config.Load()
	clierror.CheckError(err, outErr)
	clientSvc, err := sdkAutoscaling.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		viper.GetString(config.ArgServerUrl),
	)
	clierror.CheckError(err, outErr)
	autoscalingGroupSvc := sdkAutoscaling.NewGroupService(clientSvc.Get(), context.TODO())
	autoscalingGroups, _, err := autoscalingGroupSvc.List()
	clierror.CheckError(err, outErr)
	groupIds := make([]string, 0)
	if items, ok := autoscalingGroups.GroupCollection.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			if itemId, ok := item.GetIdOk(); ok && itemId != nil {
				groupIds = append(groupIds, *itemId)
			}
		}
	} else {
		return nil
	}
	return groupIds
}
