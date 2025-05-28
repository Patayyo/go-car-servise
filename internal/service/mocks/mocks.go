package mocks

import (
	"car-service/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockVehicleRepository struct {
	mock.Mock
}

func (m *MockVehicleRepository) Create(v *model.Vehicle) error {
	args := m.Called(v)
	return args.Error(0)
}

func (m *MockVehicleRepository) GetByID(id uint) (*model.Vehicle, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Vehicle), args.Error(1)
}

func (m *MockVehicleRepository) ListAll() ([]model.Vehicle, error) {
	args := m.Called()
	return args.Get(0).([]model.Vehicle), args.Error(1)
}
