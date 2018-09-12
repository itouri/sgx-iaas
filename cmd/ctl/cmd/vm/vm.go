package vm

import (
	"github.com/itouri/sgx-iaas/cmd/ctl/cmd"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(newVMCmd())
}

func newVMCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "VM",
		Short: "Control VM resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newCreateCmd(),
		newDeleteCmd(),
		// TODO newListCmd(),
	)

	return cmd
}
