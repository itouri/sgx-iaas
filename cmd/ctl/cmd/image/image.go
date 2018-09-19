package image

import (
	"github.com/spf13/cobra"
)

func NewImageCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "image",
		Short: "Control Image resources",
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}

	command.AddCommand(
		//newCryptoCmd(),
		newDeleteCmd(),
		newListCmd(),
		newRegisterCmd(),
	)

	return command
}
