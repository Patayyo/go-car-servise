package repo

import (
	"car-service/internal/model"

	"gorm.io/gorm"
)

type VehicleRepository interface {
	Create(vehicle *model.Vehicle) error
	GetByID(id uint) (*model.Vehicle, error)
	ListAll() ([]model.Vehicle, error)
}
type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *vehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) Create(vehicle *model.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *vehicleRepository) GetByID(id uint) (*model.Vehicle, error) {
	var v model.Vehicle
	if err := r.db.First(&v, id).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *vehicleRepository) ListAll() ([]model.Vehicle, error) {
	var v []model.Vehicle
	if err := r.db.Find(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}
