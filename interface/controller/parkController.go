package controller

import (
	"basket-map-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ParkController struct {
	parkUsecase usecase.ParkUsecase
}

func NewParkController(parkUsecase usecase.ParkUsecase) ParkController {
	return ParkController{parkUsecase: parkUsecase}
}

type (
	parkController interface {
		GetAllParks(ctx echo.Context) error
	}
	GetAllParksResponse struct {
		Rows []*GetAllParksResponseRow `json:"rows"`
	}
	GetAllParksResponseRow struct {
		ParkId    uint   `json:"park_id"`
		ParkName  string `json:"park_name"`
		CourtType string `json:"court_type"`
		IsFree    bool   `json:"is_free"`
	}
)

func (pc *ParkController) GetAllParks(ctx echo.Context) error {
	pr, err := pc.parkUsecase.GetAllParks()
	if err != nil {
		return err
	}
	rows := make([]*GetAllParksResponseRow, 0)
	for _, park := range pr.Rows {
		rows = append(rows, &GetAllParksResponseRow{
			ParkId:    park.ParkId,
			ParkName:  park.ParkName,
			CourtType: park.CourtType,
			IsFree:    park.IsFree,
		})

	}
	return ctx.JSON(http.StatusOK, GetAllParksResponse{
		Rows: rows,
	})
}
