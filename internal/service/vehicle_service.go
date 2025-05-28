package service

import (
	"car-service/internal/model"
	"car-service/internal/repo"
)

type VehicleService struct {
	repo repo.VehicleRepository
}

func NewVehicleService(r repo.VehicleRepository) *VehicleService {
	return &VehicleService{repo: r}
}

func (s *VehicleService) Create(vehicle *model.Vehicle) error {
	return s.repo.Create(vehicle)
}

func (s *VehicleService) GetByID(id uint) (*model.Vehicle, error) {
	return s.repo.GetByID(id)
}

func (s *VehicleService) ListAll() ([]model.Vehicle, error) {
	return s.repo.ListAll()
}
