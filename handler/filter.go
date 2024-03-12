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
		ctx.Update().filters.DateRange.Start = newTime
	case "end":
		ctx.Update().filters.DateRange.End = newTime
	default:
		return echo.NotFoundHandler(c)
	}

	return render(c, filter.DateRange(ctx.filters.DateRange.Start, ctx.filters.DateRange.End, ctx.filters.DateRange.Valid()))
}
