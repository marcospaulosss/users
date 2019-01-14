package models

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/estrategiaconcursos/UsersTest/database"
	"github.com/labstack/echo"
	"time"
)

type studentInterface interface {
	GeradorToken(students Students) (string, error)
}

// Students representa a tabela de estudante na base de dados
type Students struct {
	ID     uint    `db:"id" json:"id"`
	Name   string  `db:"name" json:"nome"`
	Email  string  `db:"email" json:"e-mail"`
	CPF    *string `db:"cpf" json:"cpf"`
	Mobile *string `db:"mobile" json:"celular,omitempty"`
	Password string `db:"password" json:"password"`
	CreatedAt int `db:"created_at" json:"creado"`
	UpdaterAt int `db:"updater_id" json:"alterado,omitempty"`
	StudentInterface studentInterface
}

// StudentsModel recebe a tabela do banco de dados
var StudentsModel = database.Sess.Collection("students")

// FindStudentID get a student
func FindStudentID(id int) map[string]interface{} {
	// students instancia da model de students
	var students []Students

	resultSQL := StudentsModel.Find("id=?", id)
	if count, _ := resultSQL.Count(); count < 1 {
		data := map[string]interface{}{
			"mensagem": "Estudante inexistente",
		}
		return data
	}

	if err := resultSQL.All(&students); err != nil {
		data := map[string]interface{}{
			"mensagem": "Erro ao tentar recuperar os dados",
		}
		return data
	}

	data := map[string]interface{}{
		"funcionarios": students,
		"titulo":       "Listagem de funcionarios",
	}

	return data
}

func (aluno Students) Login (username string, password string) (error, map[string]interface{}) {
	var students []Students

	passwordByte := []byte(password)
	passwordEncript := md5.Sum(passwordByte)
	passwordMd5 := fmt.Sprintf("%x", passwordEncript)

	resultSQL := StudentsModel.Find("email=? AND password=?", username, passwordMd5)

	if count, _ := resultSQL.Count(); count < 1 {
		return echo.ErrUnauthorized, map[string]interface{}{
			"token": nil,
		}
	}

	resultSQL.Limit(1).All(&students)

	if token, err := aluno.StudentInterface.GeradorToken(students[0]); err != nil {
		return err, map[string]interface{}{
			"token": nil,
		}
	} else {
		return err, map[string]interface{}{
			"token": token,
		}
	}
}

func (aluno *Students) GeradorToken(students Students) (string, error)  {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = students.ID
	claims["name"] = students.Name
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("Estrategia2019"))
}


