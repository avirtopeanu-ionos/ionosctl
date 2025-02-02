package commands

import (
	"context"
	"os"

	"github.com/ionos-cloud/ionosctl/commands/cloudapi-v6/completer"
	"github.com/ionos-cloud/ionosctl/pkg/constants"
	"github.com/ionos-cloud/ionosctl/pkg/core"
	cloudapiv6 "github.com/ionos-cloud/ionosctl/services/cloudapi-v6"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func K8sKubeconfigCmd() *core.Command {
	ctx := context.TODO()
	k8sCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "kubeconfig",
			Aliases:          []string{"cfg", "config"},
			Short:            "Kubernetes Kubeconfig Operations",
			Long:             "The sub-command of `ionosctl k8s kubeconfig` allows you to get the configuration file of a Kubernetes Cluster.",
			TraverseChildren: true,
		},
	}

	/*
		Get Command
	*/
	get := core.NewCommand(ctx, k8sCmd, core.CommandBuilder{
		Namespace:  "k8s",
		Resource:   "kubeconfig",
		Verb:       "get",
		Aliases:    []string{"g"},
		ShortDesc:  "Get the kubeconfig file for a Kubernetes Cluster",
		LongDesc:   "Use this command to retrieve the kubeconfig file for a given Kubernetes Cluster.\n\nRequired values to run command:\n\n* K8s Cluster Id",
		Example:    getK8sKubeconfigExample,
		PreCmdRun:  PreRunK8sClusterId,
		CmdRun:     RunK8sKubeconfigGet,
		InitClient: true,
	})
	get.AddUUIDFlag(cloudapiv6.ArgK8sClusterId, "", "", cloudapiv6.K8sClusterId, core.RequiredFlagOption())
	_ = get.Command.RegisterFlagCompletionFunc(cloudapiv6.ArgK8sClusterId, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return completer.K8sClustersIds(os.Stderr), cobra.ShellCompDirectiveNoFileComp
	})
	get.AddBoolFlag(constants.ArgNoHeaders, "", false, cloudapiv6.ArgNoHeadersDescription)
	get.AddInt32Flag(cloudapiv6.ArgDepth, cloudapiv6.ArgDepthShort, cloudapiv6.DefaultGetDepth, cloudapiv6.ArgDepthDescription)
	return k8sCmd
}

func RunK8sKubeconfigGet(c *core.CommandConfig) error {
	c.Printer.Verbose("K8s kube config with id: %v is getting...", viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId)))
	u, resp, err := c.CloudApiV6Services.K8s().ReadKubeConfig(viper.GetString(core.GetFlagName(c.NS, cloudapiv6.ArgK8sClusterId)))
	if resp != nil {
		c.Printer.Verbose(constants.MessageRequestTime, resp.RequestTime)
	}
	if err != nil {
		return err
	}
	return c.Printer.Print(u)
}
