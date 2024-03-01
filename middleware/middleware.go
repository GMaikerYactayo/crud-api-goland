package middleware

import (
	"github.com/GMaikerYactayo/crud-api-goland/authorization"
	"github.com/labstack/echo"
	"net/http"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Token invalid"})
		}
		return f(c)
	}
}
