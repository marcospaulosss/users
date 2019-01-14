package routers

import (
	"github.com/estrategiaconcursos/UsersTest/controllers"
	"github.com/labstack/echo"
)

// Init initializes the routes
func Init(app *echo.Echo) {

	apiEmployees := app.Group("/employees")
	apiEmployees.GET("", controllers.EmployeeIndex)

	apiStudents := app.Group("/students")
	apiStudents.GET("/:id", controllers.StudentIndex)
}
