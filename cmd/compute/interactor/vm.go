package interactor

import (
	"os"
)

type VmInteractor struct {
}

// do vm on graphene
func (vc *VmInteractor) VMCreate(imageID string, createReqMetadata string) error {
	// TODO どうする？
	// unixドメインソケットを使ってvmの作成を依頼
	// unix-domain-socketのclient.goを移植

	return nil
}

func (vc *VmInteractor) VMDelete() error {
	// TODO
	// sgx_destroy_enclaveを呼ぶ

	return nil
}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
