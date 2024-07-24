package middlewares

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func Auth(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := store.Get(c.Request(), "session_id")
			if session.Values["user"] == nil {
				return c.String(403, "Unauthorized")
			}
			return next(c)
		}
	}

}