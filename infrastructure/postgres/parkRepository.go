package postgres

import (
	"basket-map-api/domain"
	"basket-map-api/usecase/service"
	"gorm.io/gorm"
)

type parkRepository struct {
	db *gorm.DB
}

func NewParkRepository(db *gorm.DB) service.ParkService {
	return &parkRepository{
		db: db,
	}
}

func (p *parkRepository) GetAllParks() ([]domain.Park, error) {
	var parks []domain.Park
	if err := p.db.Table("park").Find(&parks).Error; err != nil {
		return parks, err
	}

	return parks, nil
}
