package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mobiance/goblog/internals/handlers"
)

func main() {
    // Create a new echo instance
    e := echo.New()

    e.Static("/static", "/static")
    // Add middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    Page := handlers.GetBase
    e.GET("/", Page)
    e.GET("/projects", handlers.GetProjects)
    e.GET("/hero", handlers.GetHero)
    e.GET("/about", handlers.GetAbout)
    // Start the server
    e.Logger.Fatal(e.Start(":4000"))
}
