package middlewares

import (
	"github.com/DevAthhh/DoZen/internal/auth"
	"github.com/labstack/echo"
)

func AuthRequire(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var headerRequest struct {
			Auth string `header:"Authorization"`
		}

		c.Bind(&headerRequest)

		if _, err := auth.ValidateToken(headerRequest.Auth); err == nil {
			return next(c)
		}
		return nil
	}
}
