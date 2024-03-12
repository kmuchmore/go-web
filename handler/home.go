package handler

import (
	"encoding/gob"
	"fmt"
	"learn/pon/domain"
	"learn/pon/view/pages"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type PageHandler struct{}

func init() {
	gob.Register(domain.Filters{})
	gob.Register(domain.DateRange{})
}

func (h *PageHandler) HandleHomePage(c echo.Context) error {
	cc := c.(*DataContext)
	// if cc.sess.IsNew {
	// 	if err := cc.sess.Save(c.Request(), c.Response()); err != nil {
	// 		return err
	// 	}
	// }
	return render(c, pages.Home(cc.filters, cc.files))
}

type DataContext struct {
	echo.Context
	filters domain.Filters
	files   []string
	sess    *sessions.Session
	change  bool
}

func (dc *DataContext) Update() *DataContext {
	dc.change = true
	return dc
}

func (h *PageHandler) DataCtxMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}

		cc := &DataContext{
			Context: c,
			sess:    sess,
			filters: domain.Filters{},
		}

		if sess.IsNew {
			fmt.Println("new session")
			now := time.Now()
			cc.filters.DateRange = domain.DateRange{
				Start: now.Add(-24 * time.Hour),
				End:   now,
			}
			cc.files = []string{}
			cc.change = true
		} else {
			var ok bool
			cc.filters, ok = sess.Values["filters"].(domain.Filters)
			if !ok {
				cc.change = true
			}
			cc.files, ok = sess.Values["files"].([]string)
			if !ok {
				cc.change = true
			}
		}

		return next(cc)
	}
}
