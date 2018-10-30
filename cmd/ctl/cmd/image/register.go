package image

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

// copy from https://astaxie.gitbooks.io/build-web-application-with-golang/ja/04.5.html
func runRegisterCmd(command *cobra.Command, args []string) error {
	if len(args) != 1 {
		log.Fatalf("Please provide a File Path.")
	}

	glanceURL, err := util.GetEndPoint(keystone.Glance)
	if err != nil {
		return fmt.Errorf("GetEndPoint: " + err.Error())
	}

	filename := args[0]

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//キーとなる操作
	fileWriter, err := bodyWriter.CreateFormFile("image", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//ファイルハンドル操作をオープンする
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(glanceURL+"/images", contentType, bodyBuf)
	if err != nil {
		fmt.Println("Posting feiled: URL:" + glanceURL + "/images")
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
