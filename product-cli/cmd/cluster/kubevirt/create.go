package kubevirt

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/openshift/hypershift/cmd/cluster/core"
	"github.com/openshift/hypershift/cmd/cluster/kubevirt"
)

func NewCreateCommand(opts *core.CreateOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "kubevirt",
		Short:        "Creates basic functional HostedCluster resources on KubeVirt platform",
		SilenceUsage: true,
	}

	kubevirtOpts := kubevirt.DefaultOptions()
	kubevirt.BindOptions(kubevirtOpts, cmd.Flags())
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if opts.Timeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, opts.Timeout)
			defer cancel()
		}

		if err := CreateCluster(ctx, opts, kubevirtOpts); err != nil {
			opts.Log.Error(err, "Failed to create cluster")
			return err
		}
		return nil
	}

	return cmd
}

func CreateCluster(ctx context.Context, opts *core.CreateOptions, kubevirtOpts *kubevirt.CreateOptions) error {
	return core.CreateCluster(ctx, opts, kubevirtOpts)
}
