package handler

import (
	"learn/pon/view/partial/filter"
	"time"

	"github.com/labstack/echo/v4"
)

type FilterHandler struct{}

func (h *FilterHandler) DateRange(c echo.Context) error {
	ctx := c.(*DataContext)
	timeStr := ctx.FormValue("time")
	newTime, err := time.Parse(filter.DateRangeFormat, timeStr)
	if err != nil {
		return err
	}
	switch ctx.Param("type") {
	case "start":
		ctx.Update().dateRange.Start = newTime
	case "end":
		ctx.Update().dateRange.End = newTime
	default:
		return echo.NotFoundHandler(c)
	}

	return render(c, filter.DateRange(ctx.dateRange.Start, ctx.dateRange.End, ctx.dateRange.Valid()))
}
