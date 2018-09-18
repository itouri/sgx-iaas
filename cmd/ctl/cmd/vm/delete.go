package vm

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
		Use:   "delete <VMID>",
		Short: "delete vm from image",
		RunE:  runDeleteCmd,
	}

	return command
}

func runDeleteCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a VM ID.")
	}

	novaURL, err := util.GetEndPoint(keystone.Nova)
	if err != nil {
		return err
	}

	resp, err := http.Post(novaURL+"/vm/"+args[0], "text/plain", nil)
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
