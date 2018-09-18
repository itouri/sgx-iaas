package image

import (
	"fmt"
	"log"
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/keystone/util"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete <ImageID>",
		Short: "delete image",
		RunE:  runDeleteCmd,
	}

	return command
}

func runDeleteCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a Image ID.")
	}

	glanceURL, err := util.GetEndPoint(keystone.Glance)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", glanceURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	fmt.Println(resp.StatusCode)
	return nil
}
