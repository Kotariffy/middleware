package mw

import (
	"log"

	"github.com/labstack/echo"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		value := ctx.Request().Header.Get("User-Role")

		if value == "admin" {
			log.Println("Admin role detected.")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
