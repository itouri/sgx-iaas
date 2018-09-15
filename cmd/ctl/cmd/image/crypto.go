package image

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/itouri/sgx-iaas/cmd/ctl/cmd"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/spf13/cobra"
)

func newCryptoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "crypto <FilePath>",
		Short: "crypto file",
		RunE:  runCryptoCmd,
	}

	return cmd
}

func runCryptoCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a File Path.")
	}

	// 公開鍵をとってくる
	raURL, err := cmd.GetEndPoint(keystone.RA)
	if err != nil {
		return err
	}

	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Post(raURL, "application/zip", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	fmt.Println(resp.StatusCode)
	return nil
}
