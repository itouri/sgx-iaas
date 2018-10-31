package interactor

import (
	"log"
	"net"

	"github.com/google/uuid"
)

type VmInteractor struct {
	FilePath string
}

// do vm on graphene
func (vc *VmInteractor) VMCreate(imageID uuid.UUID, createReqMetadata []byte) error {
	conn, err := net.Dial("unix", vc.FilePath)
	if err != nil {
		log.Printf("error: % \n", err)
		return err
	}
	defer conn.Close()

	hex, err := imageID.MarshalBinary()

	message := string(hex)
	message += string(createReqMetadata)

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Printf("error: %v\n", err)
		return err
	}
	log.Printf("send: %s\n", message)
	// これをしないとEOFが通知されずにレスポンスの処理まで進んでくれない
	err = conn.(*net.UnixConn).CloseWrite()
	if err != nil {
		log.Printf("error: %v\n", err)
		return err
	}
	return nil
}

func (vc *VmInteractor) VMDelete() error {
	// TODO
	// sgx_destroy_enclaveを呼ぶ

	return nil
}

func callMasterEnclave() {

}
