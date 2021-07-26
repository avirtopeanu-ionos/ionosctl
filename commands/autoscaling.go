package commands

import (
	"github.com/ionos-cloud/ionosctl/pkg/core"
	"github.com/spf13/cobra"
)

func autoscaling() *core.Command {
	autoscalingCmd := &core.Command{
		Command: &cobra.Command{
			Use:              "autoscaling",
			Aliases:          []string{"auto"},
			Short:            "Autoscaling Resources Operations",
			Long:             "The sub-commands of `ionosctl autoscaling` allow you to use Autoscaling Resources.",
			TraverseChildren: true,
		},
	}
	autoscalingCmd.AddCommand(k8sVersion())

	return autoscalingCmd
}
