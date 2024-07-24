package routers

import (
	"learn/myapp/controllers"
	"learn/myapp/middlewares"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterWeb(e *echo.Echo, store sessions.Store, db *gorm.DB) {
	router := e.Group("")
	
	// playing routes
	router.GET("/", controllers.Index)
	router.GET("/hello", controllers.Hello)
	router.GET("/hellowithparam/:name", controllers.HelloWithParams)
	router.POST("/hellowithbody", controllers.HelloWithBody)

	// auth routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login(store))
	router.GET("/logout", controllers.Logout(store))

	// testing auth route
	router.GET("/protected", controllers.Protected(), middlewares.Auth(store))

	// db route
	router.GET("/books", controllers.GetAllBooks(db), middlewares.Auth(store))
	router.GET("/books/:id", controllers.GetBook(db), middlewares.Auth(store))
	router.POST("/books", controllers.AddBook(db), middlewares.Auth(store))
	router.PATCH("/books/:id", controllers.UpdateBook(db), middlewares.Auth(store))
	router.DELETE("/books/:id", controllers.DeleteBook(db), middlewares.Auth(store))
}