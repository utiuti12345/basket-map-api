package usecase

import (
	"basket-map-api/usecase/service"
)

type ParkUsecase struct {
	parkService service.ParkService
}

func NewParkUsecase(parkService service.ParkService) ParkUsecase {
	return ParkUsecase{parkService: parkService}
}

type (
	ParkUseCase interface {
		GetAllParks() ([]*GetAllParksResponse, error)
	}
	GetAllParksResponse struct {
		Rows []*GetAllParksResponseRow
	}
	GetAllParksResponseRow struct {
		ParkId    uint
		ParkName  string
		CourtType string
		IsFree    bool
	}
)

func (pu *ParkUsecase) GetAllParks() (*GetAllParksResponse, error) {
	parks, err := pu.parkService.GetAllParks()
	if err != nil {
		return nil, err
	}
	rows := make([]*GetAllParksResponseRow, 0)
	for _, park := range parks {
		rows = append(rows, &GetAllParksResponseRow{
			ParkId:    park.ParkId,
			ParkName:  park.ParkName,
			CourtType: park.CourtType,
			IsFree:    park.IsFree,
		})
	}
	return &GetAllParksResponse{
		Rows: rows,
	}, nil
}
