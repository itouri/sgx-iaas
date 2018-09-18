package image

import (
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "show image list",
		Short: "show image list",
		RunE:  runListCmd,
	}

	return command
}

func runListCmd(command *cobra.Command, args []string) error {

}
