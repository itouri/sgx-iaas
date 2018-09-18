package image

import (
	"github.com/spf13/cobra"
)

func NewImageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Control Image resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newCryptoCmd(),
		newDeleteCmd(),
		newListCmd(),
		newRegisterCmd(),
	)

	return cmd
}
