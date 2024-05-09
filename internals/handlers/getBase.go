package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mobiance/goblog/views/components"
	"github.com/mobiance/goblog/views/pages"
)

func GetBase(c echo.Context) error {
    return render(c, pages.HomePage() )
}

func GetProjects(c echo.Context) error {
    return render(c, pages.ProjectsPage() )
}

func GetHero(c echo.Context) error {
    return render(c, components.Hero() )
}
func GetAbout(c echo.Context) error {
    return render(c, components.About() )
}
