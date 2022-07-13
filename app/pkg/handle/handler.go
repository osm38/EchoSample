package handle

import (
	"net/http"
	"echosample/pkg/structs"
	"github.com/labstack/echo/v4"
)

func GetHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func GetUserID(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetUser(c echo.Context) error {
	u := new(structs.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}