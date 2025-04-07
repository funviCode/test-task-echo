package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRoleCheckMiddleware_Admin(t *testing.T) {
	var buf bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(originalOutput)

	e := echo.New()
	testHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}

	mw := RoleCheckMiddleware()(testHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("User-Role", "admin")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := mw(c); err != nil {
		t.Fatalf("Middleware вернул ошибку: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидали статус %d, получили %d", http.StatusOK, rec.Code)
	}
	if rec.Body.String() != "OK" {
		t.Errorf("Ожидали тело ответа 'OK', получили '%s'", rec.Body.String())
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "red button user detected") {
		t.Errorf("Ожидалось сообщение в логе 'red button user detected', получили: %s", logOutput)
	}
}

func TestRoleCheckMiddleware_NonAdmin(t *testing.T) {
	var buf bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(originalOutput)

	e := echo.New()
	testHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}

	mw := RoleCheckMiddleware()(testHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("User-Role", "user")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := mw(c); err != nil {
		t.Fatalf("Middleware вернул ошибку: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Ожидали статус %d, получили %d", http.StatusOK, rec.Code)
	}
	if rec.Body.String() != "OK" {
		t.Errorf("Ожидали тело ответа 'OK', получили '%s'", rec.Body.String())
	}

	logOutput := buf.String()
	if strings.Contains(logOutput, "red button user detected") {
		t.Errorf("Не ожидалось сообщения 'red button user detected' в логе, получили: %s", logOutput)
	}
}
