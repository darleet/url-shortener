package url

import (
	"github.com/labstack/echo/v4"
	"shortener/internal/delivery/http/url/handlers"
)

func InitRouter(usecase handlers.UsecaseURL) *echo.Echo {
	e := echo.New()
	m := handlers.NewManager(usecase)
	e.POST("/shorten", m.Shorten)
	e.POST("/expand", m.Expand)
	return e
}
