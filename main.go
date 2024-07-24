package main

import (
	"learn/myapp/config"
	"learn/myapp/routers"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {

	// loading the .env file
	godotenv.Load(".env")

	// creating a new echo instance
	e := echo.New()

	// setting up middlewares
	store := sessions.NewCookieStore([]byte("your-secret-key"))
	e.Use(session.Middleware(store))
    
	// connecting to db
	db := config.Init()

	// setting up routers
	routers.RegisterWeb(e, store, db)

	// serving the application
	e.Logger.Fatal(e.Start(":8080"))
}