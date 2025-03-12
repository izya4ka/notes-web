package middleware

import (
	"github.com/izya4ka/notes-web/gateway/util"
	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		util.LogInfof("%s %s", c.Request().Method, c.Path())
		return nil
	}
}
