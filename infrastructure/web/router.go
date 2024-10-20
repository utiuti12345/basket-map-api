package web

import (
	"basket-map-api/interface/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, controller controller.ParkController) {
	e.GET("/parks", controller.GetAllParks)
}
