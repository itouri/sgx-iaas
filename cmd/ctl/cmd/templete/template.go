package templete

import (
	"github.com/spf13/cobra"
)

func NewTempleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "templete",
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
