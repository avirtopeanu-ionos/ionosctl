package commands

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.uber.org/multierr"

	"github.com/fatih/structs"
	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/completer"
	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/query"
	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/waiter"
	"github.com/ionos-cloud/ionosctl/pkg/constants"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/ionos-cloud/ionosctl/pkg/printer"
	"github.com/ionos-cloud/ionosctl/pkg/utils"
	cloudapiv6 "github.com/ionos-cloud/ionosctl/services/cloudapi-v6"
	"github.com/ionos-cloud/ionosctl/services/cloudapi-v6/resources"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GroupCmd() *core.Command {
	ctx := context.TODO()
	groupCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "group",
			Aliases:          []string{"g"},
			Short:            "Group Operations",
			Long:             "The sub-commands of `ionosctl group` allow you to list, get, create, update, delete Groups, but also operations: add/remove/list/update User from the Group.",
			TraverseChildren: true,
		},
	}
	globalFlags := groupCmd.GlobalFlags()
	globalFlags.StringSliceP(constants.ArgCols, "", defaultGroupCols, printer.ColsMessage(allGroupCols))
	_ = viper.BindPFlag(core.GetFlagName(groupCmd.Name(), constants.ArgCols), globalFlags.Lookup(constants.ArgCols))
	_ = groupCmd.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return allGroupCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	list := core.NewCommand(ctx, groupCmd, core.CommandBuilder{
		Namespace:  "group",
		Resource:   "group",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Groups",
		LongDesc:   "Use this command to get a list of available Groups available on your account\n\nYou can filter the results using `--filters` option. Use the following format to set filters: `--filters KEY1=VALUE1,KEY2=VALUE2`.\n" + completer.GroupsFiltersUsage(),
		Example:    listGroupExample,
		PreCmdRun:  PreRunGroupList,
		CmdRun:     RunGroupList,
		InitClient: true,
	})
	list.AddInt32Flag(cloudapiv6.ArgMaxResults, cloudapiv6.ArgMaxResultsShort, cloudapiv6.DefaultMaxResults, cloudapiv6.ArgMaxResultsDescription)
	list.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultListDepth, cloudapiv6.ArgDepthDescription)
	list.AddStringFlag(cloudapiv6.ArgOrderBy, "", "", cloudapiv6.ArgOrderByDescription)
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgOrderBy, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.GroupsFilters(), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringSliceFlag(cloudapiv6.ArgFilters, cloudapiv6.ArgFiltersShort, []string{""}, cloudapiv6.ArgFiltersDescription)
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgFilters, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.GroupsFilters(), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, groupCmd, core.CommandBuilder{
		Namespace:  "group",
		Resource:   "group",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get a Group",
		LongDesc:   "Use this command to retrieve details about a specific Group.\n\nRequired values to run command:\n\n* Group Id",
		Example:    getGroupExample,
		PreCmdRun:  PreRunGroupId,
		CmdRun:     RunGroupGet,
		InitClient: true,
	})
	get.AddUUIDFlag(cloudapiv6.ArgGroupId, cloudapiv6.ArgIdShort, "", cloudapiv6.GroupId, core.RequiredFlagOption())
	_ = get.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.GroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)
	get.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultGetDepth, cloudapiv6.ArgDepthDescription)

	/*
		Create Command
	*/
	create := core.NewCommand(ctx, groupCmd, core.CommandBuilder{
		Namespace:  "group",
		Resource:   "group",
		Verb:       "create",
		Aliases:    []string{"c"},
		ShortDesc:  "Create a Group",
		LongDesc:   `Use this command to create a new Group and set Group privileges. You can specify the name for the new Group. By default, all privileges will be set to false. You need to use flags privileges to be set to true.`,
		Example:    createGroupExample,
		PreCmdRun:  core.NoPreRun,
		CmdRun:     RunGroupCreate,
		InitClient: true,
	})
	create.AddStringFlag(cloudapiv6.ArgName, cloudapiv6.ArgNameShort, "Unnamed Group", "Name for the Group")
	create.AddBoolFlag(cloudapiv6.ArgCreateDc, "", false, "The group will be allowed to create Data Centers. E.g.: --create-dc=true, --create-dc=false")
	create.AddBoolFlag(cloudapiv6.ArgCreateSnapshot, "", false, "The group will be allowed to create Snapshots. E.g.: --create-snapshot=true, --create-snapshot=false")
	create.AddBoolFlag(cloudapiv6.ArgReserveIp, "", false, "The group will be allowed to reserve IP addresses. E.g.: --reserve-ip=true, --reserve-ip=false")
	create.AddBoolFlag(cloudapiv6.ArgAccessLog, "", false, "The group will be allowed to access the activity log. E.g.: --access-logs=true, --access-logs=false")
	create.AddBoolFlag(cloudapiv6.ArgCreatePcc, "", false, "The group will be allowed to create PCCs. E.g.: --create-pcc=true, --create-pcc=false")
	create.AddBoolFlag(cloudapiv6.ArgS3Privilege, "", false, "The group will be allowed to manage S3. E.g.: --s3privilege=true, --s3privilege=false")
	create.AddBoolFlag(cloudapiv6.ArgCreateBackUpUnit, "", false, "The group will be able to manage Backup Units. E.g.: --create-backup=true, --create-backup=false")
	create.AddBoolFlag(cloudapiv6.ArgCreateNic, "", false, "The group will be allowed to create NICs. E.g.: --create-nic=true, --create-nic=false")
	create.AddBoolFlag(cloudapiv6.ArgCreateK8s, "", false, "The group will be allowed to create K8s Clusters. E.g.: --create-k8s=true, --create-k8s=false")
	create.AddBoolFlag(cloudapiv6.ArgCreateFlowLog, "", false, "The group will be allowed to create Flow Logs. E.g.: --create-flowlog=true, --create-flowlog=false")
	create.AddBoolFlag(cloudapiv6.ArgAccessMonitoring, "", false, "Privilege for a group to access and manage monitoring related functionality using Monotoring-as-a-Service. E.g.: --access-monitoring=true, --access-monitoring=false")
	create.AddBoolFlag(cloudapiv6.ArgAccessCerts, "", false, "Privilege for a group to access and manage certificates. E.g.: --access-certs=true, --access-certs=false")
	create.AddBoolFlag(constants.ArgWaitForRequest, constants.ArgWaitForRequestShort, constants.DefaultWait, "Wait for Request for Group creation to be executed")
	create.AddIntFlag(constants.ArgTimeout, constants.ArgTimeoutShort, constants.DefaultTimeoutSeconds, "Timeout option for Request for Group creation [seconds]")
	create.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultCreateDepth, cloudapiv6.ArgDepthDescription)

	/*
		Update Command
	*/
	update := core.NewCommand(ctx, groupCmd, core.CommandBuilder{
		Namespace: "group",
		Resource:  "group",
		Verb:      "update",
		Aliases:   []string{"u", "up"},
		ShortDesc: "Update a Group",
		LongDesc: `Use this command to update details about a specific Group.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option.

Required values to run command:

* Group Id`,
		Example:    updateGroupExample,
		PreCmdRun:  PreRunGroupId,
		CmdRun:     RunGroupUpdate,
		InitClient: true,
	})
	update.AddUUIDFlag(cloudapiv6.ArgGroupId, cloudapiv6.ArgIdShort, "", cloudapiv6.GroupId, core.RequiredFlagOption())
	_ = update.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.GroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	update.AddStringFlag(cloudapiv6.ArgName, cloudapiv6.ArgNameShort, "", "Name for the Group")
	update.AddBoolFlag(cloudapiv6.ArgCreateDc, "", false, "The group will be allowed to create Data Centers. E.g.: --create-dc=true, --create-dc=false")
	update.AddBoolFlag(cloudapiv6.ArgCreateSnapshot, "", false, "The group will be allowed to create Snapshots. E.g.: --create-snapshot=true, --create-snapshot=false")
	update.AddBoolFlag(cloudapiv6.ArgReserveIp, "", false, "The group will be allowed to reserve IP addresses. E.g.: --reserve-ip=true, --reserve-ip=false")
	update.AddBoolFlag(cloudapiv6.ArgAccessLog, "", false, "The group will be allowed to access the activity log. E.g.: --access-logs=true, --access-logs=false")
	update.AddBoolFlag(cloudapiv6.ArgCreatePcc, "", false, "The group will be allowed to create PCCs. E.g.: --create-pcc=true, --create-pcc=false")
	update.AddBoolFlag(cloudapiv6.ArgS3Privilege, "", false, "The group will be allowed to manage S3. E.g.: --s3privilege=true, --s3privilege=false")
	update.AddBoolFlag(cloudapiv6.ArgCreateBackUpUnit, "", false, "The group will be able to manage Backup Units. E.g.: --create-backup=true, --create-backup=false")
	update.AddBoolFlag(cloudapiv6.ArgCreateNic, "", false, "The group will be allowed to create NICs. E.g.: --create-nic=true, --create-nic=false")
	update.AddBoolFlag(cloudapiv6.ArgCreateK8s, "", false, "The group will be allowed to create K8s Clusters. E.g.: --create-k8s=true, --create-k8s=false")
	update.AddBoolFlag(cloudapiv6.ArgCreateFlowLog, "", false, "The group will be allowed to create Flow Logs. E.g.: --create-flowlog=true, --create-flowlog=false")
	update.AddBoolFlag(cloudapiv6.ArgAccessMonitoring, "", false, "Privilege for a group to access and manage monitoring related functionality using Monotoring-as-a-Service. E.g.: --access-monitoring=true, --access-monitoring=false")
	update.AddBoolFlag(cloudapiv6.ArgAccessCerts, "", false, "Privilege for a group to access and manage certificates. E.g.: --access-certs=true, --access-certs=false")
	update.AddBoolFlag(constants.ArgWaitForRequest, constants.ArgWaitForRequestShort, constants.DefaultWait, "Wait for Request for Group update to be executed")
	update.AddIntFlag(constants.ArgTimeout, constants.ArgTimeoutShort, constants.DefaultTimeoutSeconds, "Timeout option for Request for Group update [seconds]")
	update.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultUpdateDepth, cloudapiv6.ArgDepthDescription)

	/*
		Delete Command
	*/
	deleteCmd := core.NewCommand(ctx, groupCmd, core.CommandBuilder{
		Namespace: "group",
		Resource:  "group",
		Verb:      "delete",
		Aliases:   []string{"d"},
		ShortDesc: "Delete a Group",
		LongDesc: `Use this operation to delete a single Group. Resources that are assigned to the Group are NOT deleted, but are no longer accessible to the Group members unless the member is a Contract Owner, Admin, or Resource Owner.

Required values to run command:

* Group Id`,
		Example:    deleteGroupExample,
		PreCmdRun:  PreRunGroupDelete,
		CmdRun:     RunGroupDelete,
		InitClient: true,
	})
	deleteCmd.AddUUIDFlag(cloudapiv6.ArgGroupId, cloudapiv6.ArgIdShort, "", cloudapiv6.GroupId, core.RequiredFlagOption())
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgGroupId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.GroupsIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	deleteCmd.AddBoolFlag(constants.ArgWaitForRequest, constants.ArgWaitForRequestShort, constants.DefaultWait, "Wait for Request for Group deletion to be executed")
	deleteCmd.AddBoolFlag(cloudapiv6.ArgAll, cloudapiv6.ArgAllShort, false, "Delete all Groups.")
	deleteCmd.AddIntFlag(constants.ArgTimeout, constants.ArgTimeoutShort, constants.DefaultTimeoutSeconds, "Timeout option for Request for Group deletion [seconds]")
	deleteCmd.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultDeleteDepth, cloudapiv6.ArgDepthDescription)

	groupCmd.AddCommand(GroupResourceCmd())
	groupCmd.AddCommand(GroupUserCmd())
	return groupCmd
}

func PreRunGroupList(c *core.PreCommandConfig) error {
	if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgFilters)) {
		return query.ValidateFilters(c, completer.GroupsFilters(), completer.GroupsFiltersUsage())
	}
	return nil
}

func PreRunGroupId(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.Command, c.NS, cloudapiv6.ArgGroupId)
}

func PreRunGroupDelete(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlagsSets(c.Command, c.NS,
		[]string{cloudapiv6.ArgGroupId},
		[]string{cloudapiv6.ArgAll},
	)
}

func PreRunGroupUserIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.Command, c.NS, cloudapiv6.ArgGroupId, cloudapiv6.ArgUserId)
}

func RunGroupList(c *core.CommandConfig) error {
	// Add Query Parameters for GET Requests
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	groups, resp, err := c.CloudApiV6Services.Groups().List(listQueryParams)
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getGroupPrint(nil, c, getGroups(groups)))
}

func RunGroupGet(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	c.Printer.Verbose("Group with id: %v is getting...", viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId)))
	u, resp, err := c.CloudApiV6Services.Groups().Get(viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId)), queryParams)
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getGroupPrint(nil, c, getGroup(u)))
}

func RunGroupCreate(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	properties := getGroupCreateInfo(c)
	newGroup := resources.Group{
		Group: ionoscloud.Group{
			Properties: &properties.GroupProperties,
		},
	}
	u, resp, err := c.CloudApiV6Services.Groups().Create(newGroup, queryParams)
	if resp != nil && printer.GetId(resp) != "" {
		c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
	}
	if err != nil {
		return err
	}
	if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getGroupPrint(nil, c, getGroup(u)))
}

func RunGroupUpdate(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	u, resp, err := c.CloudApiV6Services.Groups().Get(viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId)), queryParams)
	if err != nil {
		return err
	}
	properties := getGroupUpdateInfo(u, c)
	newGroup := resources.Group{
		Group: ionoscloud.Group{
			Properties: &properties.GroupProperties,
		},
	}
	groupUpd, resp, err := c.CloudApiV6Services.Groups().Update(viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId)), newGroup, queryParams)
	if resp != nil && printer.GetId(resp) != "" {
		c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
	}
	if err != nil {
		return err
	}
	if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getGroupPrint(resp, c, getGroup(groupUpd)))
}

func RunGroupDelete(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	groupId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgGroupId))
	if viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAll)) {
		if err := DeleteAllGroups(c); err != nil {
			return err
		}
		return c.Printer.Print(printer.Result{Resource: c.Resource, Verb: c.Verb})
	} else {
		if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete group"); err != nil {
			return err
		}
		c.Printer.Verbose("Starting deleting Group with id: %v...", groupId)
		resp, err := c.CloudApiV6Services.Groups().Delete(groupId, queryParams)
		if resp != nil && printer.GetId(resp) != "" {
			c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
		}
		if err != nil {
			return err
		}
		if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
			return err
		}
		return c.Printer.Print(getGroupPrint(resp, c, nil))
	}
}

func getGroupCreateInfo(c *core.CommandConfig) *resources.GroupProperties {
	name := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgName))
	createDc := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateDc))
	createSnap := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateSnapshot))
	reserveIp := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgReserveIp))
	accessLog := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAccessLog))
	createBackUp := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateBackUpUnit))
	createPcc := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreatePcc))
	createNic := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateNic))
	createK8s := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateK8s))
	s3 := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgS3Privilege))
	createFlowLog := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateFlowLog))
	monitoring := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAccessMonitoring))
	certs := viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAccessCerts))
	c.Printer.Verbose("Properties set for creating the group: Name: %v, CreateDatacenter: %v, CreateSnapshot: %v, "+
		"ReserveIp: %v, AccessActivityLog: %v, CreateBackupUnit: %v, CreatePcc: %v, CreateInternetAccess: %v, CreateK8sCluster: %v, "+
		"S3Privilege: %v, CreateFlowLog: %v, AccessAndManageMonitoring: %v, AccessAndManageCertificates: %v",
		name, createDc, createSnap, reserveIp, accessLog, createBackUp, createPcc, createNic, createK8s, s3, createFlowLog, monitoring, certs)
	return &resources.GroupProperties{
		GroupProperties: ionoscloud.GroupProperties{
			Name:                        &name,
			CreateDataCenter:            &createDc,
			CreateSnapshot:              &createSnap,
			ReserveIp:                   &reserveIp,
			AccessActivityLog:           &accessLog,
			CreatePcc:                   &createPcc,
			S3Privilege:                 &s3,
			CreateBackupUnit:            &createBackUp,
			CreateInternetAccess:        &createNic,
			CreateK8sCluster:            &createK8s,
			CreateFlowLog:               &createFlowLog,
			AccessAndManageMonitoring:   &monitoring,
			AccessAndManageCertificates: &certs,
		},
	}
}

func getGroupUpdateInfo(oldGroup *resources.Group, c *core.CommandConfig) *resources.GroupProperties {
	var (
		groupName                                                           string
		createDc, createSnap, createPcc, createBackUp, createNic, createK8s bool
		reserveIp, accessLog, s3, createFlowLog, monitoring, certs          bool
	)
	if properties, ok := oldGroup.GetPropertiesOk(); ok && properties != nil {
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgName)) {
			groupName = viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgName))
			c.Printer.Verbose("Property Name set: %v", groupName)
		} else {
			if name, ok := properties.GetNameOk(); ok && name != nil {
				groupName = *name
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreateDc)) {
			createDc = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateDc))
			c.Printer.Verbose("Property CreateDataCenter set: %v", createDc)
		} else {
			if dc, ok := properties.GetCreateDataCenterOk(); ok && dc != nil {
				createDc = *dc
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreateSnapshot)) {
			createSnap = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateSnapshot))
			c.Printer.Verbose("Property CreateSnapshot set: %v", createSnap)
		} else {
			if s, ok := properties.GetCreateSnapshotOk(); ok && s != nil {
				createSnap = *s
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreatePcc)) {
			createPcc = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreatePcc))
			c.Printer.Verbose("Property CreatePcc set: %v", createPcc)
		} else {
			if s, ok := properties.GetCreatePccOk(); ok && s != nil {
				createPcc = *s
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreateK8s)) {
			createK8s = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateK8s))
			c.Printer.Verbose("Property CreateK8sCluster set: %v", createK8s)
		} else {
			if s, ok := properties.GetCreateK8sClusterOk(); ok && s != nil {
				createK8s = *s
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreateNic)) {
			createNic = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateNic))
			c.Printer.Verbose("Property CreateInternetAccess set: %v", createNic)
		} else {
			if s, ok := properties.GetCreateInternetAccessOk(); ok && s != nil {
				createNic = *s
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreateBackUpUnit)) {
			createBackUp = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateBackUpUnit))
			c.Printer.Verbose("Property CreateBackupUnit set: %v", createBackUp)
		} else {
			if s, ok := properties.GetCreateBackupUnitOk(); ok && s != nil {
				createBackUp = *s
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgReserveIp)) {
			reserveIp = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgReserveIp))
			c.Printer.Verbose("Property ReserveIp set: %v", reserveIp)
		} else {
			if ip, ok := properties.GetReserveIpOk(); ok && ip != nil {
				reserveIp = *ip
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgAccessLog)) {
			accessLog = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAccessLog))
			c.Printer.Verbose("Property AccessActivityLog set: %v", accessLog)
		} else {
			if log, ok := properties.GetAccessActivityLogOk(); ok && log != nil {
				accessLog = *log
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgS3Privilege)) {
			s3 = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgS3Privilege))
			c.Printer.Verbose("Property S3Privilege set: %v", s3)
		} else {
			if s, ok := properties.GetS3PrivilegeOk(); ok && s != nil {
				s3 = *s
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgCreateFlowLog)) {
			createFlowLog = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgCreateFlowLog))
			c.Printer.Verbose("Property CreateFlowLog set: %v", createFlowLog)
		} else {
			if f, ok := properties.GetCreateFlowLogOk(); ok && f != nil {
				createFlowLog = *f
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgAccessMonitoring)) {
			monitoring = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAccessMonitoring))
			c.Printer.Verbose("Property AccessAndManageMonitoring set: %v", monitoring)
		} else {
			if m, ok := properties.GetAccessAndManageMonitoringOk(); ok && m != nil {
				monitoring = *m
			}
		}
		if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgAccessCerts)) {
			certs = viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAccessCerts))
			c.Printer.Verbose("Property AccessAndManageCertificates set: %v", certs)
		} else {
			if accessCerts, ok := properties.GetAccessAndManageCertificatesOk(); ok && accessCerts != nil {
				certs = *accessCerts
			}
		}
	}
	return &resources.GroupProperties{
		GroupProperties: ionoscloud.GroupProperties{
			Name:                        &groupName,
			CreateDataCenter:            &createDc,
			CreateSnapshot:              &createSnap,
			ReserveIp:                   &reserveIp,
			AccessActivityLog:           &accessLog,
			CreatePcc:                   &createPcc,
			S3Privilege:                 &s3,
			CreateBackupUnit:            &createBackUp,
			CreateInternetAccess:        &createNic,
			CreateK8sCluster:            &createK8s,
			CreateFlowLog:               &createFlowLog,
			AccessAndManageMonitoring:   &monitoring,
			AccessAndManageCertificates: &certs,
		},
	}
}

func DeleteAllGroups(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	c.Printer.Verbose("Getting Groups...")
	groups, resp, err := c.CloudApiV6Services.Groups().List(cloudapiv6.ParentResourceListQueryParams)
	if err != nil {
		return err
	}
	if groupsItems, ok := groups.GetItemsOk(); ok && groupsItems != nil {
		if len(*groupsItems) > 0 {
			_ = c.Printer.Warn("Groups to be deleted:")
			for _, group := range *groupsItems {
				delIdAndName := ""
				if id, ok := group.GetIdOk(); ok && id != nil {
					delIdAndName += "Group Id: " + *id
				}
				if properties, ok := group.GetPropertiesOk(); ok && properties != nil {
					if name, ok := properties.GetNameOk(); ok && name != nil {
						delIdAndName += " Group Name: " + *name
					}
				}
				_ = c.Printer.Warn(delIdAndName)
			}
			if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete all the Groups"); err != nil {
				return err
			}
			c.Printer.Verbose("Deleting all the Groups...")
			var multiErr error
			for _, group := range *groupsItems {
				if id, ok := group.GetIdOk(); ok && id != nil {
					c.Printer.Verbose("Starting deleting Group with id: %v...", *id)
					resp, err = c.CloudApiV6Services.Groups().Delete(*id, queryParams)
					if resp != nil && printer.GetId(resp) != "" {
						c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
					}
					if err != nil {
						multiErr = multierr.Append(multiErr, fmt.Errorf(constants.ErrDeleteAll, c.Resource, *id, err))
						continue
					} else {
						_ = c.Printer.Warn(fmt.Sprintf(constants.MessageDeletingAll, c.Resource, *id))
					}
					if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
						multiErr = multierr.Append(multiErr, fmt.Errorf(constants.ErrWaitDeleteAll, c.Resource, *id, err))
						continue
					}
				}
			}
			if multiErr != nil {
				return multiErr
			}
			return nil
		} else {
			return errors.New("no Groups found")
		}
	} else {
		return errors.New("could not get items of Groups")
	}
}

// Output Printing

var (
	defaultGroupCols = []string{"GroupId", "Name", "CreateDataCenter", "CreateSnapshot", "CreatePcc", "CreateBackupUnit", "CreateInternetAccess", "CreateK8s", "ReserveIp"}
	allGroupCols     = []string{"GroupId", "Name", "CreateDataCenter", "CreateSnapshot", "ReserveIp", "AccessActivityLog", "CreatePcc", "S3Privilege", "CreateBackupUnit",
		"CreateInternetAccess", "CreateK8s", "CreateFlowLog", "AccessAndManageMonitoring", "AccessAndManageCertificates"}
)

type groupPrint struct {
	GroupId                     string `json:"GroupId,omitempty"`
	Name                        string `json:"Name,omitempty"`
	CreateDataCenter            bool   `json:"CreateDataCenter,omitempty"`
	CreateSnapshot              bool   `json:"CreateSnapshot,omitempty"`
	ReserveIp                   bool   `json:"ReserveIp,omitempty"`
	AccessActivityLog           bool   `json:"AccessActivityLog,omitempty"`
	CreatePcc                   bool   `json:"CreatePcc,omitempty"`
	S3Privilege                 bool   `json:"S3Privilege,omitempty"`
	CreateBackupUnit            bool   `json:"CreateBackupUnit,omitempty"`
	CreateInternetAccess        bool   `json:"CreateInternetAccess,omitempty"`
	CreateK8s                   bool   `json:"CreateK8s,omitempty"`
	CreateFlowLog               bool   `json:"CreateFlowLog,omitempty"`
	AccessAndManageMonitoring   bool   `json:"AccessAndManageMonitoring,omitempty"`
	AccessAndManageCertificates bool   `json:"AccessAndManageCertificates,omitempty"`
}

func getGroupPrint(resp *resources.Response, c *core.CommandConfig, groups []resources.Group) printer.Result {
	r := printer.Result{}
	if c != nil {
		if resp != nil {
			r.ApiResponse = resp
			r.Resource = c.Resource
			r.Verb = c.Verb
			r.WaitForRequest = viper.GetBool(core.GetFlagName(c.NS, constants.ArgWaitForRequest))
		}
		if groups != nil {
			r.OutputJSON = groups
			r.KeyValue = getGroupsKVMaps(groups)
			r.Columns = printer.GetHeaders(allGroupCols, defaultGroupCols, viper.GetStringSlice(core.GetFlagName(c.Resource, constants.ArgCols)))
		}
	}
	return r
}

func getGroups(groups resources.Groups) []resources.Group {
	u := make([]resources.Group, 0)
	if items, ok := groups.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			u = append(u, resources.Group{Group: item})
		}
	}
	return u
}

func getGroup(u *resources.Group) []resources.Group {
	groups := make([]resources.Group, 0)
	if u != nil {
		groups = append(groups, resources.Group{Group: u.Group})
	}
	return groups
}

func getGroupsKVMaps(gs []resources.Group) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(gs))
	for _, g := range gs {
		var gPrint groupPrint
		if id, ok := g.GetIdOk(); ok && id != nil {
			gPrint.GroupId = *id
		}
		if properties, ok := g.GetPropertiesOk(); ok && properties != nil {
			if name, ok := properties.GetNameOk(); ok && name != nil {
				gPrint.Name = *name
			}
			if createDc, ok := properties.GetCreateDataCenterOk(); ok && createDc != nil {
				gPrint.CreateDataCenter = *createDc
			}
			if createSnapshot, ok := properties.GetCreateSnapshotOk(); ok && createSnapshot != nil {
				gPrint.CreateSnapshot = *createSnapshot
			}
			if reserveIp, ok := properties.GetReserveIpOk(); ok && reserveIp != nil {
				gPrint.ReserveIp = *reserveIp
			}
			if accessLog, ok := properties.GetAccessActivityLogOk(); ok && accessLog != nil {
				gPrint.AccessActivityLog = *accessLog
			}
			if createPcc, ok := properties.GetCreatePccOk(); ok && createPcc != nil {
				gPrint.CreatePcc = *createPcc
			}
			if s3, ok := properties.GetS3PrivilegeOk(); ok && s3 != nil {
				gPrint.S3Privilege = *s3
			}
			if createBackup, ok := properties.GetCreateBackupUnitOk(); ok && createBackup != nil {
				gPrint.CreateBackupUnit = *createBackup
			}
			if createNic, ok := properties.GetCreateInternetAccessOk(); ok && createNic != nil {
				gPrint.CreateInternetAccess = *createNic
			}
			if createK8s, ok := properties.GetCreateK8sClusterOk(); ok && createK8s != nil {
				gPrint.CreateK8s = *createK8s
			}
			if createFlowLogs, ok := properties.GetCreateFlowLogOk(); ok && createFlowLogs != nil {
				gPrint.CreateFlowLog = *createFlowLogs
			}
			if accessMonitoring, ok := properties.GetAccessAndManageMonitoringOk(); ok && accessMonitoring != nil {
				gPrint.AccessAndManageMonitoring = *accessMonitoring
			}
			if accessCerts, ok := properties.GetAccessAndManageCertificatesOk(); ok && accessCerts != nil {
				gPrint.AccessAndManageCertificates = *accessCerts
			}
		}
		o := structs.Map(gPrint)
		out = append(out, o)
	}
	return out
}
