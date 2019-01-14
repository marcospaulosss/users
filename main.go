package main

import (
	"github.com/estrategiaconcursos/UsersTest/middleware"
	"github.com/estrategiaconcursos/UsersTest/routers"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {

	App := echo.New()

	middleware.Init(App)
	//middleware.InitPermission(App)
	routers.Init(App)

	App.Start(":3000")
}
