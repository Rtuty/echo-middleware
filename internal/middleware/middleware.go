package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strings"
)

const roleAdmin = "admin"

func RoleChecker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(value, roleAdmin) {
			log.Print("red button user detected")
		}

		if err := next(ctx); err != nil {
			return err
		}

		return nil
	}
}
