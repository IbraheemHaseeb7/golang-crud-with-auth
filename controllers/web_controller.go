package controllers

import "github.com/labstack/echo/v4"

func Index(c echo.Context) error {
	return c.JSON(200, map[string]string{"message": "Hello, World!"})
}

func Hello(c echo.Context) error {
	name := c.QueryParam("name")
	return c.JSON(200, map[string]string{"message": "Hello, " + name + "!"})
}

func HelloWithParams(c echo.Context) error {
	name := c.Param("name")
	return c.JSON(200, map[string]string{"message": "Hello, " + name + "!"})
}

func HelloWithBody(c echo.Context) error {
	type RequestBody struct {
		Name string `json:"name"`
	}
	
	body := new(RequestBody)
	if err := c.Bind(body); err != nil {
		return err
	}
	
	return c.JSON(200, map[string]string{"message": "Hello, " + body.Name + "!"})
}