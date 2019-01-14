package middleware

import (
	"github.com/labstack/echo"
	casbin "github.com/casbin/casbin"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
  )

  // InitPermission initializes middlewares
func InitPermission(app *echo.Echo) {
	app.Use(casbin_mw.Middleware(casbin.NewEnforcer("middleware/keymatch_model.conf", "middleware/keymatch_policy.csv")))

	// ce := casbin.NewEnforcer("middleware/keymatch_model.conf", "")
	// ce.AddPolicy("alice", "/students/*", "GET")
	// app.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
	// 	Enforcer: ce,
	// }))
}