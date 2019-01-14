package models

import (
	"github.com/estrategiaconcursos/UsersTest/database"
	"strconv"
	"time"

	//"github.com/labstack/echo"
	//"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// StudentsModel recebe a tabela do banco de dados
var StudentsModelTest = database.Sess.Collection("students")

type studentMock struct {
	mock.Mock
}

func (m *studentMock) GeradorToken(students Students) (string, error)  {
	args := m.Called(students)
	return args.String(0), args.Error(1)
}

func TestLogin(t *testing.T) {
	cpf := "66858185909"
	mobile := "9999999999"

	//ts := tspb.Timestamp{}
	//dateNow, err := ptypes.TimestampProto(time.Now())
	//dateNowString := strconv.FormatInt(dateNow.Seconds, 10)

	//Insert Students for tests
	var datetime = time.Now()
	timeFormat := datetime.Format(time.RFC3339)
	s, _ := strconv.Atoi(timeFormat)
	//fmt.Println(timeFormat)
	//fmt.Println(dateComplect)
	student := Students{
		Name: "Teste Unitario",
		Email: "testeunitario@estrategiaconcursos.com.br",
		CPF: &cpf,
		Mobile: &mobile,
		Password: "e427b9d67d9d19ca638fb63ac3508822",
		CreatedAt: s,
		UpdaterAt: s,
	}

	idCreated, err := StudentsModelTest.Insert(student)

	println(idCreated)
	println(err)


	//t.Run("A=1", func(t *testing.T) {
	//	aluno := Students{
	//		1,
	//		"Teste 2",
	//		"teste22@estrategiaconcursos.com.br",
	//		&cpf,
	//		&mobile,
	//		"e427b9d67d9d19ca638fb63ac3508822",
	//		nil,
	//	}
	//	theStudentMock := studentMock{}
	//	theStudentMock.On("GeradorToken", aluno).Return("TokenGeradoComSucesso", nil)
	//	aluno.StudentInterface = &theStudentMock
	//
	//	res, token := aluno.Login("teste22@estrategiaconcursos.com.br", "Estrategia123")
	//	assert.Equal(t,res, nil, "token valido erro nulo")
	//	assert.Equal(t, token, token, "token valido")
	//})
	//
	//t.Run("A=2", func(t *testing.T) {
	//	aluno1 := Students{
	//		1,
	//		"Teste 2",
	//		"teste22@estrategiaconcursos.com.br",
	//		&cpf,
	//		&mobile,
	//		"e427b9d67d9d19ca638fb63ac3508822",
	//		nil,
	//	}
	//	theStudentMock1 := studentMock{}
	//	theStudentMock1.On("GeradorToken", aluno1).Return("TokenGeradoComSucesso", nil)
	//	aluno1.StudentInterface = &theStudentMock1
	//
	//	res1, token1 := aluno1.Login("teste22@estrategiaconcursos.com.brr", "Estrategia123")
	//	assert.Equal(t, res1, echo.ErrUnauthorized, "nao autorizado")
	//	assert.Equal(t, token1["token"], nil, "token nao gerado")
	//})
	//
	//t.Run("A=3", func(t *testing.T) {
	//	aluno2 := Students{
	//		1,
	//		"Teste 2",
	//		"teste22@estrategiaconcursos.com.br",
	//		&cpf,
	//		&mobile,
	//		"e427b9d67d9d19ca638fb63ac3508822",
	//		nil,
	//	}
	//	theStudentMock2 := studentMock{}
	//	theStudentMock2.On("GeradorToken", aluno2).Return("", echo.ErrInternalServerError)
	//	aluno2.StudentInterface = &theStudentMock2
	//
	//	res2, token2 := aluno2.Login("teste22@estrategiaconcursos.com.br", "Estrategia123")
	//	assert.Equal(t, res2, echo.ErrInternalServerError, "nao autorizado")
	//	assert.Equal(t, token2["token"], nil, "token nao gerado")
	//})
}

func BenchmarkHello(b *testing.B) {
	//var cpf *string = nil
	//var mobile *string = nil
	//aluno := Students{
	//	1,
	//	"Teste 2",
	//	"teste22@estrategiaconcursos.com.br",
	//	cpf,
	//	mobile,
	//	"e427b9d67d9d19ca638fb63ac3508822",
	//	nil,
	//}
	//theStudentMock1 := studentMock{}
	//theStudentMock1.On("GeradorToken", aluno).Return("TokenGeradoComSucesso", nil)
	//aluno.StudentInterface = &theStudentMock1
	//for i := 0; i < b.N; i++ {
	//	aluno.Login("teste22@estrategiaconcursos.com.br", "Estrategia123")
	//}
}

func BeforeStudent(student Students) {

}