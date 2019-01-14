package models

import (
	"github.com/estrategiaconcursos/UsersTest/database"
)

// Employees representa a tabela de funcionarios na base de dados
type Employees struct {
	ID       uint   `db:"id" json:"id"`
	Name     string `db:"name" json:"nome"`
	Email    string `db:"email" json:"e-mail"`
	Document string `db:"cpf" json:"documento"`
}

// EmployeesModel recebe a tabela do banco de dados
var EmployeesModel = database.Sess.Collection("employees")
