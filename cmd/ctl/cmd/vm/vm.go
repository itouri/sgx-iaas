package vm

import (
	"github.com/spf13/cobra"
)

func NewVMCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vm",
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
