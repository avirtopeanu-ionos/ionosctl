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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func K8sNodeCmd() *core.Command {
	ctx := context.TODO()
	k8sCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "node",
			Aliases:          []string{"n"},
			Short:            "Kubernetes Node Operations",
			Long:             "The sub-commands of `ionosctl k8s node` allow you to list, get, recreate, delete Kubernetes Nodes.",
			TraverseChildren: true,
		},
	}
	globalFlags := k8sCmd.GlobalFlags()
	globalFlags.StringSliceP(constants.ArgCols, "", defaultK8sNodeCols, printer.ColsMessage(defaultK8sNodeCols))
	_ = viper.BindPFlag(core.GetFlagName(k8sCmd.Name(), constants.ArgCols), globalFlags.Lookup(constants.ArgCols))
	_ = k8sCmd.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultK8sNodeCols, cobra.ShellCompDirectiveNoFileComp
	})

	/*
		List Command
	*/
	list := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace:  "k8s",
		Resource:   "node",
		Verb:       "list",
		Aliases:    []string{"l", "ls"},
		ShortDesc:  "List Kubernetes Nodes",
		LongDesc:   "Use this command to get a list of existing Kubernetes Nodes.\n\nYou can filter the results using `--filters` option. Use the following format to set filters: `--filters KEY1=VALUE1,KEY2=VALUE2`.\n" + completer.K8sNodesFiltersUsage() + "\n\nRequired values to run command:\n\n* K8s Cluster Id\n* K8s NodePool Id",
		Example:    listK8sNodesExample,
		PreCmdRun:  PreRunK8sNodesList,
		CmdRun:     RunK8sNodeList,
		InitClient: true,
	})
	list.AddUUIDFlag(cloudapiv6.ArgK8sClusterId, "", "", cloudapiv6.K8sClusterId, core.RequiredFlagOption())
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddUUIDFlag(cloudapiv6.ArgK8sNodePoolId, "", "", cloudapiv6.K8sNodePoolId, core.RequiredFlagOption())
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(list.NS, cloudapiv6.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddInt32Flag(cloudapiv6.ArgMaxResults, cloudapiv6.ArgMaxResultsShort, cloudapiv6.DefaultMaxResults, cloudapiv6.ArgMaxResultsDescription)
	list.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultListDepth, cloudapiv6.ArgDepthDescription)
	list.AddStringFlag(cloudapiv6.ArgOrderBy, "", "", cloudapiv6.ArgOrderByDescription)
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgOrderBy, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodesFilters(), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringSliceFlag(cloudapiv6.ArgFilters, cloudapiv6.ArgFiltersShort, []string{""}, cloudapiv6.ArgFiltersDescription)
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgFilters, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodesFilters(), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace:  "k8s",
		Resource:   "node",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get a Kubernetes Node",
		LongDesc:   "Use this command to retrieve details about a specific Kubernetes Node.You can wait for the Node to be in \"ACTIVE\" state using `--wait-for-state` flag together with `--timeout` option.\n\nRequired values to run command:\n\n* K8s Cluster Id\n* K8s NodePool Id\n* K8s Node Id",
		Example:    getK8sNodeExample,
		PreCmdRun:  PreRunK8sClusterNodesIds,
		CmdRun:     RunK8sNodeGet,
		InitClient: true,
	})
	get.AddUUIDFlag(cloudapiv6.ArgK8sClusterId, "", "", cloudapiv6.K8sClusterId, core.RequiredFlagOption())
	_ = get.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddUUIDFlag(cloudapiv6.ArgK8sNodePoolId, "", "", cloudapiv6.K8sNodePoolId, core.RequiredFlagOption())
	_ = get.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(get.NS, cloudapiv6.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddUUIDFlag(cloudapiv6.ArgK8sNodeId, cloudapiv6.ArgIdShort, "", cloudapiv6.K8sNodeId, core.RequiredFlagOption())
	_ = get.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodeId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodesIds(os.Stderr,
			viper.GetString(core.GetFlagName(get.NS, cloudapiv6.ArgK8sClusterId)),
			viper.GetString(core.GetFlagName(get.NS, cloudapiv6.ArgK8sNodePoolId)),
		), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddBoolFlag(constants.ArgWaitForState, constants.ArgWaitForStateShort, constants.DefaultWait, "Wait for specified Node to be in ACTIVE state")
	get.AddIntFlag(constants.ArgTimeout, constants.ArgTimeoutShort, cloudapiv6.K8sTimeoutSeconds, "Timeout option for waiting for Node to be in ACTIVE state [seconds]")
	get.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)
	get.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultGetDepth, cloudapiv6.ArgDepthDescription)

	/*
		Recreate Command
	*/
	recreate := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "recreate",
		Aliases:   []string{"r"},
		ShortDesc: "Recreate a Kubernetes Node",
		LongDesc: `You can recreate a single Kubernetes Node.

Managed Kubernetes starts a process which based on the NodePool's template creates & configures a new Node, waits for status "ACTIVE", and migrates all the Pods from the faulty Node, deleting it once empty. While this operation occurs, the NodePool will have an extra billable "ACTIVE" Node.

Required values to run command:

* K8s Cluster Id
* K8s NodePool Id
* K8s Node Id`,
		Example:    recreateK8sNodeExample,
		PreCmdRun:  PreRunK8sClusterNodesIds,
		CmdRun:     RunK8sNodeRecreate,
		InitClient: true,
	})
	recreate.AddUUIDFlag(cloudapiv6.ArgK8sClusterId, "", "", cloudapiv6.K8sClusterId, core.RequiredFlagOption())
	_ = recreate.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	recreate.AddUUIDFlag(cloudapiv6.ArgK8sNodePoolId, "", "", cloudapiv6.K8sNodePoolId, core.RequiredFlagOption())
	_ = recreate.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(recreate.NS, cloudapiv6.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	recreate.AddUUIDFlag(cloudapiv6.ArgK8sNodeId, cloudapiv6.ArgIdShort, "", cloudapiv6.K8sNodeId, core.RequiredFlagOption())
	_ = recreate.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodeId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodesIds(os.Stderr,
			viper.GetString(core.GetFlagName(recreate.NS, cloudapiv6.ArgK8sClusterId)),
			viper.GetString(core.GetFlagName(recreate.NS, cloudapiv6.ArgK8sNodePoolId)),
		), cobra.ShellCompDirectiveNoFileComp
	})
	recreate.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultCreateDepth, cloudapiv6.ArgDepthDescription)

	/*
		Delete Command
	*/
	deleteCmd := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "delete",
		Aliases:   []string{"d"},
		ShortDesc: "Delete a Kubernetes Node",
		LongDesc: `This command deletes a Kubernetes Node within an existing Kubernetes NodePool in a Cluster.

Required values to run command:

* K8s Cluster Id
* K8s NodePool Id
* K8s Node Id`,
		Example:    deleteK8sNodeExample,
		PreCmdRun:  PreRunK8sClusterNodesIdsAll,
		CmdRun:     RunK8sNodeDelete,
		InitClient: true,
	})
	deleteCmd.AddUUIDFlag(cloudapiv6.ArgK8sClusterId, "", "", cloudapiv6.K8sClusterId, core.RequiredFlagOption())
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	deleteCmd.AddUUIDFlag(cloudapiv6.ArgK8sNodePoolId, "", "", cloudapiv6.K8sNodePoolId, core.RequiredFlagOption())
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodePoolId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodePoolsIds(os.Stderr, viper.GetString(core.GetFlagName(deleteCmd.NS, cloudapiv6.ArgK8sClusterId))), cobra.ShellCompDirectiveNoFileComp
	})
	deleteCmd.AddUUIDFlag(cloudapiv6.ArgK8sNodeId, "", "", cloudapiv6.K8sNodeId, core.RequiredFlagOption())
	_ = deleteCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sNodeId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sNodesIds(os.Stderr,
			viper.GetString(core.GetFlagName(deleteCmd.NS, cloudapiv6.ArgK8sClusterId)),
			viper.GetString(core.GetFlagName(deleteCmd.NS, cloudapiv6.ArgK8sNodePoolId)),
		), cobra.ShellCompDirectiveNoFileComp
	})
	deleteCmd.AddBoolFlag(cloudapiv6.ArgAll, cloudapiv6.ArgAllShort, false, "Delete all the Kubernetes Nodes within an existing Kubernetes NodePool in a Cluster.")
	deleteCmd.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultDeleteDepth, cloudapiv6.ArgDepthDescription)

	return k8sCmd
}

func PreRunK8sNodesList(c *core.PreCommandConfig) error {
	if err := core.CheckRequiredFlags(c.Command, c.NS, cloudapiv6.ArgK8sClusterId, cloudapiv6.ArgK8sNodePoolId); err != nil {
		return err
	}
	if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgFilters)) {
		return query.ValidateFilters(c, completer.K8sNodesFilters(), completer.K8sNodesFiltersUsage())
	}
	return nil
}

func PreRunK8sClusterNodesIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.Command, c.NS, cloudapiv6.ArgK8sClusterId, cloudapiv6.ArgK8sNodePoolId, cloudapiv6.ArgK8sNodeId)
}

func PreRunK8sClusterNodesIdsAll(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlagsSets(c.Command, c.NS,
		[]string{cloudapiv6.ArgK8sClusterId, cloudapiv6.ArgK8sNodePoolId, cloudapiv6.ArgAll},
		[]string{cloudapiv6.ArgK8sClusterId, cloudapiv6.ArgK8sNodePoolId, cloudapiv6.ArgK8sNodeId},
	)
}

func RunK8sNodeList(c *core.CommandConfig) error {
	c.Printer.Verbose("Listing Nodes from K8s NodePool ID: %v from K8s Cluster ID: %v",
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId)))
	// Add Query Parameters for GET Requests
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	k8ss, resp, err := c.CloudApiV6Services.K8s().ListNodes(
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId)),
		listQueryParams,
	)
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getK8sNodePrint(c, getK8sNodes(k8ss)))
}

func RunK8sNodeGet(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	if err := utils.WaitForState(c, waiter.K8sNodeStateInterrogator, viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodeId))); err != nil {
		return err
	}
	c.Printer.Verbose("Getting K8s Node with ID: %v from K8s NodePool ID: %v from K8s Cluster ID: %v......",
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodeId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId)))
	u, resp, err := c.CloudApiV6Services.K8s().GetNode(
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodeId)),
		queryParams,
	)
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getK8sNodePrint(c, getK8sNode(u)))
}

func RunK8sNodeRecreate(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	if err := utils.AskForConfirm(c.Stdin, c.Printer, "recreate k8s node"); err != nil {
		return err
	}
	k8sClusterId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId))
	k8sNodePoolId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId))
	k8sNodeId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodeId))
	c.Printer.Verbose("K8sClusterId: %v, K8sNodePoolId: %v, K8sNodeId: %v",
		k8sClusterId, k8sNodePoolId, k8sNodeId)
	c.Printer.Verbose("Recreating Node...")
	resp, err := c.CloudApiV6Services.K8s().RecreateNode(k8sClusterId, k8sNodePoolId, k8sNodeId, queryParams)
	if resp != nil && printer.GetId(resp) != "" {
		c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Warn("Status: Command node recreate has been successfully executed")
}

func RunK8sNodeDelete(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	clusterId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId))
	nodepoolId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId))
	nodeId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodeId))
	if viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAll)) {
		if err := DeleteAllK8sNodes(c); err != nil {
			return err
		}
		return c.Printer.Print(printer.Result{Resource: c.Resource, Verb: c.Verb})
	} else {
		if err := utils.AskForConfirm(c.Stdin, c.Printer, "delete k8s node"); err != nil {
			return err
		}
		c.Printer.Verbose("Starting deleting Node with ID: %v from K8s NodePool ID: %v from K8s Cluster ID: %v...", nodeId, nodepoolId, clusterId)
		resp, err := c.CloudApiV6Services.K8s().DeleteNode(clusterId, nodepoolId, nodeId, queryParams)
		if resp != nil && printer.GetId(resp) != "" {
			c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
		}
		if err != nil {
			return err
		}
		return c.Printer.Print(printer.Result{Resource: c.Resource, Verb: c.Verb})
	}
}

func DeleteAllK8sNodes(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	clusterId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId))
	nodepoolId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sNodePoolId))
	c.Printer.Verbose("K8sCluster ID: %v", clusterId)
	c.Printer.Verbose("K8sNodePool ID: %v", nodepoolId)
	c.Printer.Verbose("Getting K8sNodes...")
	k8sNodes, resp, err := c.CloudApiV6Services.K8s().ListNodes(clusterId, nodepoolId, cloudapiv6.ParentResourceListQueryParams)
	if err != nil {
		return err
	}
	if k8sNodesItems, ok := k8sNodes.GetItemsOk(); ok && k8sNodesItems != nil {
		if len(*k8sNodesItems) > 0 {
			_ = c.Printer.Warn("K8sNodes to be deleted:")
			for _, dc := range *k8sNodesItems {
				delIdAndName := ""
				if id, ok := dc.GetIdOk(); ok && id != nil {
					delIdAndName += "K8sNodes Id: " + *id
				}
				if properties, ok := dc.GetPropertiesOk(); ok && properties != nil {
					if name, ok := properties.GetNameOk(); ok && name != nil {
						delIdAndName += " K8sNodes Name: " + *name
					}
				}
				_ = c.Printer.Warn(delIdAndName)
			}
			if err = utils.AskForConfirm(c.Stdin, c.Printer, "delete all the K8sNodes"); err != nil {
				return err
			}
			c.Printer.Verbose("Deleting all the K8sNodes...")
			var multiErr error
			for _, dc := range *k8sNodesItems {
				if id, ok := dc.GetIdOk(); ok && id != nil {
					c.Printer.Verbose("Staring deleting Node with ID: %v from K8s NodePool ID: %v from K8s Cluster ID: %v...",
						*id, nodepoolId, clusterId)
					resp, err = c.CloudApiV6Services.K8s().DeleteNode(clusterId, nodepoolId, *id, queryParams)
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
			return errors.New("no K8sNodes found")
		}
	} else {
		return errors.New("could not get items of K8sNodes")
	}
}

// Output Printing

var defaultK8sNodeCols = []string{"NodeId", "Name", "K8sVersion", "PublicIP", "PrivateIP", "State"}

type K8sNodePrint struct {
	NodeId     string `json:"NodeId,omitempty"`
	Name       string `json:"Name,omitempty"`
	K8sVersion string `json:"K8sVersion,omitempty"`
	PublicIP   string `json:"PublicIP,omitempty"`
	PrivateIP  string `json:"PrivateIP,omitempty"`
	State      string `json:"State,omitempty"`
}

func getK8sNodePrint(c *core.CommandConfig, k8ss []resources.K8sNode) printer.Result {
	r := printer.Result{}
	if c != nil {
		if k8ss != nil {
			r.OutputJSON = k8ss
			r.KeyValue = getK8sNodesKVMaps(k8ss)
			r.Columns = printer.GetHeadersAllDefault(defaultK8sNodeCols, viper.GetStringSlice(core.GetFlagName(c.Resource, constants.ArgCols)))
		}
	}
	return r
}

func getK8sNodes(k8ss resources.K8sNodes) []resources.K8sNode {
	u := make([]resources.K8sNode, 0)
	if items, ok := k8ss.GetItemsOk(); ok && items != nil {
		for _, item := range *items {
			u = append(u, resources.K8sNode{KubernetesNode: item})
		}
	}
	return u
}

func getK8sNode(u *resources.K8sNode) []resources.K8sNode {
	k8ss := make([]resources.K8sNode, 0)
	if u != nil {
		k8ss = append(k8ss, resources.K8sNode{KubernetesNode: u.KubernetesNode})
	}
	return k8ss
}

func getK8sNodesKVMaps(us []resources.K8sNode) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(us))
	for _, u := range us {
		var uPrint K8sNodePrint
		if id, ok := u.GetIdOk(); ok && id != nil {
			uPrint.NodeId = *id
		}
		if properties, ok := u.GetPropertiesOk(); ok && properties != nil {
			if name, ok := properties.GetNameOk(); ok && name != nil {
				uPrint.Name = *name
			}
			if v, ok := properties.GetK8sVersionOk(); ok && v != nil {
				uPrint.K8sVersion = *v
			}
			if v, ok := properties.GetPublicIPOk(); ok && v != nil {
				uPrint.PublicIP = *v
			}
			if priv, ok := properties.GetPrivateIPOk(); ok && priv != nil {
				uPrint.PrivateIP = *priv
			}
		}
		if meta, ok := u.GetMetadataOk(); ok && meta != nil {
			if state, ok := meta.GetStateOk(); ok && state != nil {
				uPrint.State = *state
			}
		}
		o := structs.Map(uPrint)
		out = append(out, o)
	}
	return out
}
