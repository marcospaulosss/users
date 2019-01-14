package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http/httptest"
	"testing"
)

func TestGetEmployee(t *testing.T) {
	App := echo.New()

	req := httptest.NewRequest(echo.GET, "/employees/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := App.NewContext(req, rec)

	res := EmployeeIndex(c)
	fmt.Println(res)
}
