package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTimeHandler(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/time", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := TimeHandler(c); err != nil {
		t.Fatal("Handler returned an error", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "incorrect date") && !strings.Contains(body, "Number of days") {
		t.Errorf("Incorrect body answer %s", body)
	}
}
