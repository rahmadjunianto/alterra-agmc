package mocks

import (
	"day4/util"
	"io"
	"net/http/httptest"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoMock struct {
	E *echo.Echo
}

func (em *EchoMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	em.E.Validator = &util.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(method, path, body)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)

	return c, rec
}
