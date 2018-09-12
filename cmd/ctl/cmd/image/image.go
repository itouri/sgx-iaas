package image

import (
	"github.com/itouri/sgx-iaas/cmd/ctl/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(newImageCmd())
}

func newImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Image",
		Short: "Control Image resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newCryptoCmd(),
		newDeleteCmd(),
		newListCmd(),
		newResiterCmd(),
	)

	return cmd
}
