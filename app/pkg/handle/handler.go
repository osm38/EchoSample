package handle

import (
	"net/http"
	"echosample/pkg/structs"
	"github.com/labstack/echo/v4"
	"echosample/pkg/db/repository"
	"echosample/pkg/authenticate"
	"echosample/pkg/session"
	"fmt"
	"echosample/pkg/bind"
	ub "echosample/pkg/util/bind"
	"echosample/pkg/log"
	"time"
)

var logger = log.GetInstance("handler")

func GetHelloWorld(c echo.Context) error {
	fmt.Println("GetHelloWorld execute!!")
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
	return c.JSON(http.StatusOK, u)
}

func DoLogin(c echo.Context) error {
	l := &bind.Login{}
	if err := ub.Bind(c, l); err != nil {
		logger.Error(err, "DoLogin is Fail!!", nil)
		return err
	}

	isSuccess := authenticate.Login(l.Name, l.Password)

	if !isSuccess {
		fmt.Println("login fail!")
		return c.NoContent(http.StatusUnauthorized)
	}
	fmt.Println("login Success!!")

	sess := session.GetSession(c)
	fmt.Println("GetSession Success!!")
	cookie := &http.Cookie{}
	cookie.Name = "sessID"
	cookie.Value = sess.ID
	cookie.Expires = time.Now().Add(86400 * 7)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func AjaxSample(c echo.Context) error {
	user := repository.FindUserByID(1)
	_ = user
	return c.NoContent(http.StatusOK)
}

func PostRoot(c echo.Context) error {
	fmt.Println("PostRoot Start!!")
	req := c.Request()
	url := req.URL.Path
	fmt.Println("url: " + url)
	return c.NoContent(http.StatusOK)
}