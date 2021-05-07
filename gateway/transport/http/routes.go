package http

import (
	"github.com/hidayatullahap/evermos/gateway/entity"
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo, app *entity.App) {
	h := NewHandler(app)

	v1 := e.Group("/api/v1")
	v1.POST("/orders", h.Order)
}
