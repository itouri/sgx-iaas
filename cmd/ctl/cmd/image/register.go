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

func newRegisterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register <FilePath>",
		Short: "register stack",
		RunE:  runRegisterCmd,
	}

	return cmd
}

func runRegisterCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a File Path.")
	}

	glanceURL, err := cmd.GetEndPoint(keystone.Glance)
	if err != nil {
		return err
	}

	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Post(glanceURL, "application/zip", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	fmt.Println(resp.StatusCode)
	return nil
}
