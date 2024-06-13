package kubevirt

import (
	"github.com/spf13/cobra"

	"github.com/openshift/hypershift/cmd/nodepool/core"
	kubevirtnodepool "github.com/openshift/hypershift/cmd/nodepool/kubevirt"
)

func NewCreateCommand(coreOpts *core.CreateNodePoolOptions) *cobra.Command {
	platformOpts := kubevirtnodepool.DefaultOptions()
	cmd := &cobra.Command{
		Use:          "kubevirt",
		Short:        "Creates basic functional NodePool resources for KubeVirt platform",
		SilenceUsage: true,
	}
	kubevirtnodepool.BindOptions(platformOpts, cmd.Flags())
	cmd.RunE = coreOpts.CreateRunFunc(platformOpts)

	return cmd
}
