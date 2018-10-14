package interactor

import (
	"fmt"
	"os"
	"os/exec"
)

type VmInteractor struct {
}

// do vm on graphene
func createVM() error {
	// TODO どうする？
	// shをexecするのが一番早い気がする
	out, err := exec.Command("").Output()
	fmt.Println(string(out))
	if err != nil {
		return err
	}

	return nil
}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
