package main

import (
	"learn/pon/handler"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dev = true

func main() {
	app := echo.New()
	pageHandler := handler.PageHandler{}

	store := sessions.NewCookieStore([]byte("supersecretninjastuff"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
	}

	if dev {
		app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Response().Header().Set("Cache-Control", "no-store")
				return next(c)
			}
		})
	}

	app.Use(session.MiddlewareWithConfig(session.Config{
		Skipper: middleware.DefaultSkipper,
		Store:   store,
	}))

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   5,
		Skipper: middleware.DefaultSkipper,
	}))

	app.Static("/static", "view/static")
	app.GET("/", pageHandler.HandleHomePage)

	// api := app.Group("/api")

	if err := app.Start(":3000"); err != nil {
		panic(err)
	}
}
