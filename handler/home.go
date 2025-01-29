package handler

import (
	"learn/pon/view/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func (h *PageHandler) HandleHomePage(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}
