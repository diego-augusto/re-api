package server

import (
	"net/http"
	"net/http/httptest"
	"re-go-challenge/pkg/packets"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_GetPacks(t *testing.T) {

	body := `{"sizes": [250, 500, 1000, 2000, 5000], "items": 10000}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/packs", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := server{
		packSrv: packets.New(),
	}
	err := s.GetPacks(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `[{"size":5000,"quantity":2}]`+"\n", rec.Body.String())
}

func Test_GetPacks_InvalidBody(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/packs", strings.NewReader("invalid"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := server{
		packSrv: packets.New(),
	}
	err := s.GetPacks(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
