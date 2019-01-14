package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init initializes middlewares
func Init(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Pre(middleware.RemoveTrailingSlash())
}
