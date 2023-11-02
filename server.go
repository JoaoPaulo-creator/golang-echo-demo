package main

import (
	"example.io/echo-demo/config"
	"example.io/echo-demo/controllers"
	"github.com/labstack/echo/v4/middleware"
	//	"example.io/echo-demo/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.PGConnection()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/transaction/create", controllers.CreateUserController)

	e.Logger.Fatal(e.Start(":3000"))

}
