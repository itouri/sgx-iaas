package cmd

import (
	"fmt"
	"os"

	"github.com/itouri/sgx-iaas/cmd/ctl/cmd/image"
	"github.com/itouri/sgx-iaas/cmd/ctl/cmd/templete"
	"github.com/itouri/sgx-iaas/cmd/ctl/cmd/vm"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sgxiaas",
	Short: "",
	Long:  ``,
}

func init() {
	RootCmd.AddCommand(image.NewImageCmd())
	RootCmd.AddCommand(templete.NewTempleteCmd())
	RootCmd.AddCommand(vm.NewVMCmd())
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
