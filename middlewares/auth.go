package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		Authorization := c.Request().Header.Get("Authorization")
		if Authorization != "November 10, 2009" {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized!",
			})
			return errors.New("")
		}
		next(c)
		return nil
	}
}
