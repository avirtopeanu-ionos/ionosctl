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
			Long:             "The sub-commands of `ionosctl autoscaling` allow you to manage Autoscaling Resources.",
			TraverseChildren: true,
		},
	}

	autoscalingCmd.AddCommand(autoscalingTemplate())
	autoscalingCmd.AddCommand(autoscalingNicTemplate())
	autoscalingCmd.AddCommand(autoscalingVolumeTemplate())

	return autoscalingCmd
}
