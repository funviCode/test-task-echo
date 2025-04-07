package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const (
	userRoleHeader   = "User-Role"
	adminRole        = "admin"
	redButtonMessage = "red button user detected"
)

func RoleCheckMiddleware() echo.MiddlewareFunc { // Явно возвращает Middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			userRole := ctx.Request().Header.Get(userRoleHeader)

			if strings.Contains(strings.ToLower(userRole), adminRole) {
				log.Println(redButtonMessage)
			}

			return next(ctx)
		}
	}
}
