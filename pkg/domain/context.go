package domain

import (
	"net/http"

	"github.com/labstack/echo"
)

// type Context interface {
// 	Param(string) string
// 	Bind(interface{}) error
// 	JSON(int, interface{}) error
// 	String(int, string) error
// 	NoContent(int) error
// }

type Context struct {
	echo.Context
}

func (c *Context) BindValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}
	return nil
}
