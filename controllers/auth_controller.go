package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {

	type RequestBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	body := new(RequestBody)
	if err := c.Bind(body); err != nil {
		return err
	}

	c.Request().SetBasicAuth(body.Email, body.Password)

	return c.JSON(200, map[string]string{"message": "Register Success"})
}

func Login (store sessions.Store) echo.HandlerFunc {
	return func (c echo.Context) error {

		type RequestBody struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		body := new(RequestBody)
		if err := c.Bind(body); err != nil {
			return err
		}

		// Validate the username and password
		if body.Email == "ibraheemhaseeb7@gmail.com" && body.Password == "a1s2d3f4" {
			// Create a new session
			session, _ := store.Get(c.Request(), "session_id")
			session.Values["user"] = body.Email
			session.Save(c.Request(), c.Response())
			return c.String(http.StatusOK, "Logged in")
		}
	
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func Protected() echo.HandlerFunc {
	return func (c echo.Context) error {
		return c.String(http.StatusOK, "Protected Accessed")
	}
}

func Logout(store sessions.Store) echo.HandlerFunc {
	return func (c echo.Context) error {
		session, _ := store.Get(c.Request(), "session_id")
		session.Values["user"] = nil
		session.Save(c.Request(), c.Response())

		return c.String(http.StatusOK, "Logged out")
	}
}