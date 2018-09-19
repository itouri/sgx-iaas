package image

import (
	"io"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spacemonkeygo/openssl"
	"github.com/itouri/sgx-iaas/cmd/ctl/cmd"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/spf13/cobra"
)

var (
	storePubKeyPath string
)

func init() {
	// TODO reading from configure file
	storePubKeyPath = "/etc/sgxiaas/"
}

func newCryptoCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "crypto <FilePath>",
		Short: "crypto file with RA server public key",
		RunE:  runCryptoCmd,
	}

	return command
}

func runCryptoCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a file path.")
	}
	cryptoFilePath := args[0]

	if !isFileExist(storePubKeyPath) {
		err := getKey()
		if err != nil {
			return err
		}
	}

	// TODO encrypt
	// めんどくさいから各自コマンドツールでやってほしいよ，，，，

	// err := encrypt()
	// if err := nil {
	// 	return err
	// }

	return nil
}

func isFileExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func getKey() error {
	raURL, err := util.GetEndPoint(keystone.RA)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", raURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}
	fmt.Println(resp.StatusCode)

	// store key
	file, err := os.Create(storePubKeyPath)
	if err := nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	if err := nil {
		return err
	}

	return nil
}
