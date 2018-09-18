package image

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/itouri/sgx-iaas/cmd/keystone/util"
	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
	"github.com/spf13/cobra"
)

func newRegisterCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "register <FilePath>",
		Short: "register stack",
		RunE:  runRegisterCmd,
	}

	return command
}

func runRegisterCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a File Path.")
	}

	glanceURL, err := util.GetEndPoint(keystone.Glance)
	if err != nil {
		return err
	}

	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()

	w := multipart.NewWriter(&buf)

	fw, err := w.CreateFormFile("file", file)
	if err != nil {
		return err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return err
	}
	w.Close()

	req, err := http.NewRequest(http.MethodPost, uri, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
}
