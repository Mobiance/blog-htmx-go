package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mobiance/goblog/views/shared"
)

func GetBase(c echo.Context) error {
    return render(c, shared.Page( "Hello" ))
}
