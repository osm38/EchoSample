package bind

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

func Bind(c echo.Context, i interface{}) error {
	err := c.Bind(i)
	if err != nil {
		err = fmt.Errorf("Bind Error! Wrap: %w", err)
		c.NoContent(http.StatusInternalServerError)
	}

	return err
}