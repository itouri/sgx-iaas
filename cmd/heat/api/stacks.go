package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/heat/engine"
	"github.com/itouri/sgx-iaas/pkg/domain"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"

	yaml "gopkg.in/yaml.v2"
)

type Req struct {
	Template string `json:"template"`
}

func GetStacks(c domain.Context) error {

}

func PostStack(c domain.Context) error {
	req := &Req{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	template := &heat.Template{}

	// yamlファイルの解釈
	err = yaml.Unmarshal([]byte(req.Template), template)
	if err != nil {
		//TODO StatusOKではない
		return c.String(http.StatusOK, err.Error())
	}

	//TODO valitation
	engine.RegisterStack(template)

	return nil
}
