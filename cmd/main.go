package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mobiance/goblog/internals/handlers"
)

func main() {
    // Create a new echo instance
    e := echo.New()

    // Add middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    Page := handlers.GetBase
    e.GET("/", Page)
    // Start the server
    e.Logger.Fatal(e.Start(":3000"))
}
