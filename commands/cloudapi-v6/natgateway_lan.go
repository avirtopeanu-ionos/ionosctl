package commands

import (
	"context"
	"errors"
	"os"

	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/query"

	"github.com/fatih/structs"
	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/completer"
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

func NatgatewayLanCmd() *core.Command {
	ctx := context.TODO()
	natgatewayLanCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "lan",
			Short:            "NAT Gateway Lan Operations",
			Long:             "The sub-commands of `ionosctl natgateway lan` allow you to add, list, remove NAT Gateway Lans.",
			TraverseChildren: true,
		},
	}

	/*
		List Command
	*/
	list := core.NewCommand(ctx, natgatewayLanCmd, core.CommandBuilder{
		Namespace: "natgateway",
		Resource:  "lan",
		Verb:      "list",
		Aliases:   []string{"l", "ls"},
		ShortDesc: "List NAT Gateway Lans",
		LongDesc: `Use this command to list NAT Gateway Lans from a specified NAT Gateway.

Required values to run command:

* Data Center Id
* NAT Gateway Id`,
		Example:    listNatGatewayLanExample,
		PreCmdRun:  PreRunDcNatGatewayIds,
		CmdRun:     RunNatGatewayLanList,
		InitClient: true,
	})
	list.AddUUIDFlag(cloudapiv6.ArgDataCenterId, "", "", cloudapiv6.DatacenterId, core.RequiredFlagOption())
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.DataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddUUIDFlag(cloudapiv6.ArgNatGatewayId, "", "", cloudapiv6.NatGatewayId, core.RequiredFlagOption())
	_ = list.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgNatGatewayId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.NatGatewaysIds(os.Stderr, viper.GetString(core.GetFlagName(list.NS, cloudapiv6.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	list.AddStringSliceFlag(constants.ArgCols, "", defaultNatGatewayLanCols, printer.ColsMessage(defaultNatGatewayLanCols))
	_ = list.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultNatGatewayLanCols, cobra.ShellCompDirectiveNoFileComp
	})
	list.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)
	list.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultDeleteDepth, cloudapiv6.ArgDepthDescription)

	/*
		Add Command
	*/
	add := core.NewCommand(ctx, natgatewayLanCmd, core.CommandBuilder{
		Namespace: "natgateway",
		Resource:  "lan",
		Verb:      "add",
		Aliases:   []string{"a"},
		ShortDesc: "Add a NAT Gateway Lan",
		LongDesc: `Use this command to add a NAT Gateway Lan in a specified NAT Gateway.

If IPs are not set manually, using ` + "`" + `--ips` + "`" + ` option, an IP will be automatically assigned. IPs must contain valid subnet mask. If user will not provide any IP then system will generate an IP with /24 subnet.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option.

Required values to run command:

* Data Center Id
* NAT Gateway Id
* Lan Id`,
		Example:    addNatGatewayLanExample,
		PreCmdRun:  PreRunDcNatGatewayLanIds,
		CmdRun:     RunNatGatewayLanAdd,
		InitClient: true,
	})
	add.AddUUIDFlag(cloudapiv6.ArgDataCenterId, "", "", cloudapiv6.DatacenterId, core.RequiredFlagOption())
	_ = add.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.DataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	add.AddUUIDFlag(cloudapiv6.ArgNatGatewayId, "", "", cloudapiv6.NatGatewayId, core.RequiredFlagOption())
	_ = add.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgNatGatewayId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.NatGatewaysIds(os.Stderr, viper.GetString(core.GetFlagName(add.NS, cloudapiv6.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	add.AddIntFlag(cloudapiv6.ArgLanId, cloudapiv6.ArgIdShort, 1, cloudapiv6.LanId, core.RequiredFlagOption())
	_ = add.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgLanId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.LansIds(os.Stderr, viper.GetString(core.GetFlagName(add.NS, cloudapiv6.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	add.AddStringSliceFlag(cloudapiv6.ArgIps, "", nil, "Collection of Gateway IPs. If not set, it will automatically reserve public IPs")
	add.AddBoolFlag(constants.ArgWaitForRequest, constants.ArgWaitForRequestShort, constants.DefaultWait, "Wait for the Request for NAT Gateway Lan addition to be executed")
	add.AddIntFlag(constants.ArgTimeout, constants.ArgTimeoutShort, constants.DefaultTimeoutSeconds, "Timeout option for Request for NAT Gateway Lan addition [seconds]")
	add.AddStringSliceFlag(constants.ArgCols, "", defaultNatGatewayLanCols, printer.ColsMessage(defaultNatGatewayLanCols))
	_ = add.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultNatGatewayLanCols, cobra.ShellCompDirectiveNoFileComp
	})
	add.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultMiscDepth, cloudapiv6.ArgDepthDescription)

	/*
		Remove Command
	*/
	removeCmd := core.NewCommand(ctx, natgatewayLanCmd, core.CommandBuilder{
		Namespace: "natgateway",
		Resource:  "lan",
		Verb:      "remove",
		Aliases:   []string{"r"},
		ShortDesc: "Remove a NAT Gateway Lan",
		LongDesc: `Use this command to remove a specified NAT Gateway Lan from a NAT Gateway.

You can wait for the Request to be executed using ` + "`" + `--wait-for-request` + "`" + ` option. You can force the command to execute without user input using ` + "`" + `--force` + "`" + ` option.

Required values to run command:

* Data Center Id
* NAT Gateway Id
* Lan Id`,
		Example:    removeNatGatewayLanExample,
		PreCmdRun:  PreRunDcNatGatewayLanRemove,
		CmdRun:     RunNatGatewayLanRemove,
		InitClient: true,
	})
	removeCmd.AddUUIDFlag(cloudapiv6.ArgDataCenterId, "", "", cloudapiv6.DatacenterId, core.RequiredFlagOption())
	_ = removeCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgDataCenterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.DataCentersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	removeCmd.AddUUIDFlag(cloudapiv6.ArgNatGatewayId, "", "", cloudapiv6.NatGatewayId, core.RequiredFlagOption())
	_ = removeCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgNatGatewayId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.NatGatewaysIds(os.Stderr, viper.GetString(core.GetFlagName(removeCmd.NS, cloudapiv6.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	removeCmd.AddIntFlag(cloudapiv6.ArgLanId, cloudapiv6.ArgIdShort, 1, cloudapiv6.LanId, core.RequiredFlagOption())
	_ = removeCmd.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgLanId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.LansIds(os.Stderr, viper.GetString(core.GetFlagName(removeCmd.NS, cloudapiv6.ArgDataCenterId))), cobra.ShellCompDirectiveNoFileComp
	})
	removeCmd.AddBoolFlag(constants.ArgWaitForRequest, constants.ArgWaitForRequestShort, constants.DefaultWait, "Wait for the Request for NAT Gateway Lan deletion to be executed")
	removeCmd.AddIntFlag(constants.ArgTimeout, constants.ArgTimeoutShort, constants.DefaultTimeoutSeconds, "Timeout option for Request for NAT Gateway Lan deletion [seconds]")
	removeCmd.AddStringSliceFlag(constants.ArgCols, "", defaultNatGatewayLanCols, printer.ColsMessage(defaultNatGatewayLanCols))
	_ = removeCmd.Command.RegisterFlagCompletionFunc(constants.ArgCols, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return defaultNatGatewayLanCols, cobra.ShellCompDirectiveNoFileComp
	})
	removeCmd.AddBoolFlag(cloudapiv6.ArgAll, cloudapiv6.ArgAllShort, false, "Remove all NAT Gateway Lans.")
	removeCmd.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultDeleteDepth, cloudapiv6.ArgDepthDescription)

	return natgatewayLanCmd
}

func PreRunDcNatGatewayLanIds(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlags(c.Command, c.NS, cloudapiv6.ArgDataCenterId, cloudapiv6.ArgNatGatewayId, cloudapiv6.ArgLanId)
}

func PreRunDcNatGatewayLanRemove(c *core.PreCommandConfig) error {
	return core.CheckRequiredFlagsSets(c.Command, c.NS,
		[]string{cloudapiv6.ArgDataCenterId, cloudapiv6.ArgNatGatewayId, cloudapiv6.ArgLanId},
		[]string{cloudapiv6.ArgDataCenterId, cloudapiv6.ArgNatGatewayId, cloudapiv6.ArgAll},
	)
}

func RunNatGatewayLanList(c *core.CommandConfig) error {
	ng, resp, err := c.CloudApiV6Services.NatGateways().Get(
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgDataCenterId)),
		viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgNatGatewayId)),
		resources.QueryParams{},
	)
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(getNatGatewayLanPrint(nil, c, getNatGatewayLans(ng)))
}

func RunNatGatewayLanAdd(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	dcId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgDataCenterId))
	natGatewayId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgNatGatewayId))
	ng, _, err := c.CloudApiV6Services.NatGateways().Get(dcId, natGatewayId, queryParams)
	if err != nil {
		return err
	}
	c.Printer.Verbose("Adding NatGateway with id %v to Datacenter with id: %v", natGatewayId, dcId)
	input := getNewNatGatewayLanInfo(c, ng)
	ng, resp, err := c.CloudApiV6Services.NatGateways().Update(dcId, natGatewayId, *input, queryParams)
	if resp != nil && printer.GetId(resp) != "" {
		c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
	}
	if err != nil {
		return err
	}
	if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
		return err
	}
	return c.Printer.Print(getNatGatewayLanPrint(resp, c, getNatGatewayLans(ng)))
}

func RunNatGatewayLanRemove(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	if viper.GetBool(core.GetFlagName(c.NS, cloudapiv6.ArgAll)) {
		if err := RemoveAllNatGatewayLans(c); err != nil {
			return err
		}
		return c.Printer.Print(printer.Result{Resource: c.Resource, Verb: c.Verb})
	} else {
		if err := utils.AskForConfirm(c.Stdin, c.Printer, "remove nat gateway lan"); err != nil {
			return err
		}
		dcId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgDataCenterId))
		natGatewayId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgNatGatewayId))
		ng, _, err := c.CloudApiV6Services.NatGateways().Get(dcId, natGatewayId, queryParams)
		if err != nil {
			return err
		}
		c.Printer.Verbose("Removing NatGateway with id %v to Datacenter with id: %v", natGatewayId, dcId)
		input := removeNatGatewayLanInfo(c, ng)
		ng, resp, err := c.CloudApiV6Services.NatGateways().Update(dcId, natGatewayId, *input, queryParams)
		if resp != nil && printer.GetId(resp) != "" {
			c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
		}
		if err != nil {
			return err
		}
		if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
			return err
		}
		return c.Printer.Print(getNatGatewayLanPrint(resp, c, nil))
	}
}

func RemoveAllNatGatewayLans(c *core.CommandConfig) error {
	listQueryParams, err := query.GetListQueryParams(c)
	if err != nil {
		return err
	}
	queryParams := listQueryParams.QueryParams
	dcId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgDataCenterId))
	natGatewayId := viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgNatGatewayId))
	c.Printer.Verbose("Datacenter ID: %v", dcId)
	c.Printer.Verbose("NatGateway ID: %v", natGatewayId)
	c.Printer.Verbose("Getting NatGateway...")
	natGateway, resp, err := c.CloudApiV6Services.NatGateways().Get(dcId, natGatewayId, cloudapiv6.ParentResourceQueryParams)
	if err != nil {
		return err
	}
	if natGatewayProperties, ok := natGateway.GetPropertiesOk(); ok && natGatewayProperties != nil {
		if lansOk, ok := natGatewayProperties.GetLansOk(); ok && lansOk != nil {
			if len(*lansOk) > 0 {
				_ = c.Printer.Warn("NAT Gateways Lan to be removed:")
				for _, lan := range *lansOk {
					if id, ok := lan.GetIdOk(); ok && id != nil {
						_ = c.Printer.Warn("NAT Gateways Lan Id: " + string(*id))
					}
				}
			} else {
				return errors.New("no NAT Gateways Lans found")
			}
		} else {
			return errors.New("could not get items of NAT Gateways Lans")
		}
	}
	if err = utils.AskForConfirm(c.Stdin, c.Printer, "remove all the NAT Gateways Lans"); err != nil {
		return err
	}
	c.Printer.Verbose("Removing all the NAT Gateways Lans...")
	proper := make([]ionoscloud.NatGatewayLanProperties, 0)
	if natGateway != nil {
		if properties, ok := natGateway.GetPropertiesOk(); ok && properties != nil {
			natGatewaysProps := &resources.NatGatewayProperties{
				NatGatewayProperties: ionoscloud.NatGatewayProperties{
					Lans: &proper,
				},
			}
			natGateway, resp, err = c.CloudApiV6Services.NatGateways().Update(dcId, natGatewayId, *natGatewaysProps, queryParams)
			if resp != nil && printer.GetId(resp) != "" {
				c.Printer.Verbose(constants.MessageRequestInfo, printer.GetId(resp), resp.RequestTime)
			}
			if err != nil {
				return err
			}
			if err = utils.WaitForRequest(c, waiter.RequestInterrogator, printer.GetId(resp)); err != nil {
				return err
			}
		}
	}
	return nil
}

func getNewNatGatewayLanInfo(c *core.CommandConfig, oldNg *resources.NatGateway) *resources.NatGatewayProperties {
	var proper []ionoscloud.NatGatewayLanProperties
	if oldNg != nil {
		if properties, ok := oldNg.GetPropertiesOk(); ok && properties != nil {
			if lans, ok := properties.GetLansOk(); ok && lans != nil {
				proper = *lans
			}
		}
	}
	input := ionoscloud.NatGatewayLanProperties{}
	if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgLanId)) {
		lanId := viper.GetInt32(core.GetFlagName(c.NS, cloudapiv6.ArgLanId))
		input.SetId(lanId)
		c.Printer.Verbose("Property Id set: %v", lanId)
	}
	if viper.IsSet(core.GetFlagName(c.NS, cloudapiv6.ArgIps)) {
		gatewayIps := viper.GetStringSlice(core.GetFlagName(c.NS, cloudapiv6.ArgIps))
		input.SetGatewayIps(gatewayIps)
		c.Printer.Verbose("Property GatewayIps set: %v", gatewayIps)
	}
	proper = append(proper, input)
	return &resources.NatGatewayProperties{
		NatGatewayProperties: ionoscloud.NatGatewayProperties{
			Lans: &proper,
		},
	}
}

func removeNatGatewayLanInfo(c *core.CommandConfig, oldNg *resources.NatGateway) *resources.NatGatewayProperties {
	proper := make([]ionoscloud.NatGatewayLanProperties, 0)
	if oldNg != nil {
		if properties, ok := oldNg.GetPropertiesOk(); ok && properties != nil {
			if lans, ok := properties.GetLansOk(); ok && lans != nil {
				for _, lanItem := range *lans {
					if id, ok := lanItem.GetIdOk(); ok && id != nil {
						if *id != viper.GetInt32(core.GetFlagName(c.NS, cloudapiv6.ArgLanId)) {
							proper = append(proper, lanItem)
						}
					}
				}
			}
		}
	}
	return &resources.NatGatewayProperties{
		NatGatewayProperties: ionoscloud.NatGatewayProperties{
			Lans: &proper,
		},
	}
}

// Output Printing

var defaultNatGatewayLanCols = []string{"NatGatewayLanId", "GatewayIps"}

type NatGatewayLanPrint struct {
	NatGatewayLanId int32    `json:"NatGatewayLanId,omitempty"`
	GatewayIps      []string `json:"GatewayIps,omitempty"`
}

func getNatGatewayLanPrint(resp *resources.Response, c *core.CommandConfig, ss []resources.NatGatewayLanProperties) printer.Result {
	r := printer.Result{}
	if c != nil {
		if resp != nil {
			r.ApiResponse = resp
			r.Resource = c.Resource
			r.Verb = c.Verb
			r.WaitForRequest = viper.GetBool(core.GetFlagName(c.NS, constants.ArgWaitForRequest))
			r.WaitForState = viper.GetBool(core.GetFlagName(c.NS, constants.ArgWaitForState))
		}
		if ss != nil {
			r.OutputJSON = ss
			r.KeyValue = getNatGatewayLansKVMaps(ss)
			r.Columns = printer.GetHeadersAllDefault(defaultNatGatewayLanCols, viper.GetStringSlice(core.GetFlagName(c.Resource, constants.ArgCols)))
		}
	}
	return r
}

func getNatGatewayLans(ng *resources.NatGateway) []resources.NatGatewayLanProperties {
	ss := make([]resources.NatGatewayLanProperties, 0)
	if ng != nil {
		if properties, ok := ng.GetPropertiesOk(); ok && properties != nil {
			if lans, ok := properties.GetLansOk(); ok && lans != nil {
				for _, lanItem := range *lans {
					ss = append(ss, resources.NatGatewayLanProperties{
						NatGatewayLanProperties: lanItem,
					})
				}
			}
		}
	}
	return ss
}

func getNatGatewayLansKVMaps(ss []resources.NatGatewayLanProperties) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(ss))
	for _, s := range ss {
		var natgatewayPrint NatGatewayLanPrint
		if id, ok := s.GetIdOk(); ok && id != nil {
			natgatewayPrint.NatGatewayLanId = *id
		}
		if ips, ok := s.GetGatewayIpsOk(); ok && ips != nil {
			natgatewayPrint.GatewayIps = *ips
		}
		o := structs.Map(natgatewayPrint)
		out = append(out, o)
	}
	return out
}
