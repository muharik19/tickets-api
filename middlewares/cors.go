package middlewares

import "github.com/labstack/echo"

func AllowOriginSkipper(c echo.Context) bool {
	return false
}
