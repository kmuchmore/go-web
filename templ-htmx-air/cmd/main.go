package main

import (
	"learn/pon/handler"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	pageHandler := handler.PageHandler{}
	filterHandler := handler.FilterHandler{}
	// fileHander := handler.FileHandler{}

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   5,
		Skipper: middleware.DefaultSkipper,
	}))
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("thisisagoodsecret"))))
	app.Static("/static", "view/static")
	app.GET("/", pageHandler.HandleHomePage)

	api := app.Group("/api")
	api.POST("/times", filterHandler.HandleDateRange)

	app.Start(":3000")
}
