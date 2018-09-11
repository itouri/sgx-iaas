package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// TODO AddCommand to Root
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
		newListCmd(),
	)

	return cmd
}
