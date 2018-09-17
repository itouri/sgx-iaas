package templete

import (
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(newTempleteCmd())
}

func newTempleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Templete",
		Short: "Control Templete resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newRegisterCmd(),
		newDeleteCmd(),
		newListCmd(),
	)

	return cmd
}
