package cmd

import (
	"fmt"
	"os"

	"github.com/itouri/sgx-iaas/cmd/keystone/util"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/spf13/cobra"
)

var (
	endpoints        map[keystone.EnumServiceType]string
	keystoneEndpoint string
)

var RootCmd = &cobra.Command{
	Use:   "sgxiaas",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func GetEndPoint(est keystone.EnumServiceType) (string, error) {
	if endpoints == nil {
		endpoints = map[keystone.EnumServiceType]string{}
	}

	if ep, ok := endpoints[est]; ok {
		return ep, nil
	}

	// TODO clientの通信の部分を作り込む
	return util.ResolveServiceEndpoint(keystoneEndpoint, est)
}
