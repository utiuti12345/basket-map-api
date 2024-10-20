package service

import "basket-map-api/domain"

type ParkService interface {
	GetAllParks() ([]domain.Park, error)
}
