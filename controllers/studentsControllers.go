package controllers

import (
	"github.com/estrategiaconcursos/UsersTest/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// StudentIndex returns the student by Id
func StudentIndex(c echo.Context) error {
	// studentId recebendo parametro vindo na requisição
	studentID, _ := strconv.Atoi(c.Param("id"))

	students := models.FindStudentID(studentID)

	return c.JSON(http.StatusOK, students)
}
