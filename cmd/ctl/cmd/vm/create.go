package vm

import "github.com/spf13/cobra"

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <ImageID>",
		Short: "create vm from image",
		RunE:  runCreateCmd,
	}

	return cmd
}

func runCreateCmd(cmd *cobra.Command, args []string) error {
}
