package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"time"
)

const (
	targetYear  = 2027
	targetMonth = time.January
	targetDay   = 1
)

func TimeHandler(ctx echo.Context) error {
	d := time.Date(targetYear, targetMonth, targetDay, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d).Hours()
	days := int(math.Ceil(dur / 24.0))

	s := fmt.Sprintf("Number of days: %d", days)
	if days < 1 {
		return ctx.String(http.StatusOK, "incorrect date")
	}

	return ctx.String(http.StatusOK, s)
}
