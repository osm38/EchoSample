package authenticate

import(
	"echosample/pkg/db/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"fmt"
	"net/http"
	"echosample/pkg/session"
	"echosample/pkg/log"
)

var logger = log.GetInstance("authenticate")

type Result bool
const (
	Success Result = true
	Fail Result = false
)

func Login(name string, pw string) Result {
	user := repository.FindUserByName(name)

	if user.Password == pw {
		return Success
	}

	return Fail
}

type SimpleAuthConfig struct {
	Skipper middleware.Skipper
	Validator SimpleAuthValidator
}

var DefaultSimpleAuthConfig = SimpleAuthConfig {
	Skipper: DefaultSimpleSkipper,
	Validator: nil,
}

func DefaultSimpleSkipper(c echo.Context) bool {
	req := c.Request()
	if url := req.URL.Path; url == "/" {
		return true
	}
	return false
}

type SimpleAuthValidator func(string, string, echo.Context) (bool, error)

func DefaultSimpleAuthValidator(id string, pw string, c echo.Context) (bool, error) {
	return true, nil
}

func SimpleAuthWithConfig(config SimpleAuthConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultSimpleAuthConfig.Skipper
	}
	if config.Validator == nil {
		config.Validator = DefaultSimpleAuthConfig.Validator
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				next(c)
			}
			cookie, err := c.Cookie("sessID")
			if err != nil {
				logger.Error(err, "Cookie is None!!", nil)
				return c.NoContent(http.StatusInternalServerError)
			}
			session.Auth(cookie.Value)
			return next(c)
		}
	}
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessID")
		if err != nil {
			logger.Error(err, "Cookie is None!!", nil)
			return c.NoContent(http.StatusInternalServerError)
		}
		session.Auth(cookie.Value)
		fmt.Println("auth middle execute!!")
		return next(c)
	}
}