package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	cc := c.(*DataContext)
	if cc.change {
		cc.sess.Values["filters"] = cc.filters
		cc.sess.Values["files"] = cc.files
		if err := cc.sess.Save(cc.Request(), cc.Response()); err != nil {
			return err
		}
	}

	return component.Render(c.Request().Context(), c.Response())
}
