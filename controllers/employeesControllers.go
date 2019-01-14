package controllers

import (
	"github.com/estrategiaconcursos/UsersTest/models"
	"github.com/labstack/echo"
	"net/http"
)

// EmployeeIndex e a controller da Api de employees para get
func EmployeeIndex(c echo.Context) error {
	var employees []models.Employees

	if err := models.EmployeesModel.Find("id", 777).All(&employees); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar recuperar os dados",
		})
	}

	data := map[string]interface{}{
		"funcionarios": employees,
		"titulo":       "Listagem de funcionarios",
	}

	return c.JSON(http.StatusOK, data)
}
