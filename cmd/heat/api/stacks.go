package api

import (
	"net/http"

	"github.com/itouri/sgx-iaas/cmd/heat/interactor"
	"github.com/itouri/sgx-iaas/pkg/domain"
	"github.com/itouri/sgx-iaas/pkg/domain/heat"

	yaml "gopkg.in/yaml.v2"
)

var tempInteractor *interactor.TemplateInteractor

func init() {
	tempInteractor = &interactor.TemplateInteractor{}
}

type Req struct {
	Template string `json:"template"`
}

func GetStacks(c echo.Context) error {

}

func PostStack(c echo.Context) error {
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
	tempInteractor.InsertTemplate(template)
	if err != nil {
		return err
	}

	return nil
}
