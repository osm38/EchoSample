package handle

import (
	"net/http"
	"echosample/pkg/structs"
	"github.com/labstack/echo/v4"
	"echosample/pkg/db/repository"
	"echosample/pkg/authenticate"
	"echosample/pkg/session"
	"fmt"
)

func GetHelloWorld(c echo.Context) error {
	u := &structs.User{
		Name: "hoge",
		Age: 20,
	}
	// return c.String(http.StatusOK, "Hello World!")
	return c.JSON(http.StatusOK, u)
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
	return c.JSON(http.StatusOK, u)
}

func DoLogin(c echo.Context) error {
	name := c.FormValue("name")
	pw := c.FormValue("password")

	isSuccess := authenticate.Login(name, pw)

	if !isSuccess {
		fmt.Println("login fail!")
		return c.NoContent(http.StatusUnauthorized)
	}
	fmt.Println("login Success!!")

	sess := session.GetSession(c)
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func AjaxSample(c echo.Context) error {
	user := repository.FindUserByID(1)
	_ = user
	return c.NoContent(http.StatusOK)
}
