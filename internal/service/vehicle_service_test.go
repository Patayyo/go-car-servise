package service

import (
	"car-service/internal/model"
	"errors"
	"testing"

	"car-service/internal/service/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreate_Success(t *testing.T) {
	mockRepo := new(mocks.MockVehicleRepository)
	service := NewVehicleService(mockRepo)

	expected := &model.Vehicle{
		Make: "Honda",
		Mark: "Civic",
		Year: 2020,
	}
	mockRepo.On("Create", expected).Return(nil)

	err := service.Create(expected)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreate_Failure(t *testing.T) {
	mockRepo := new(mocks.MockVehicleRepository)
	service := NewVehicleService(mockRepo)

	expected := &model.Vehicle{
		Make: "Honda",
		Mark: "Civic",
		Year: 2020,
	}
	mockRepo.On("Create", expected).Return(errors.New("db error"))

	err := service.Create(expected)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockVehicleRepository)
	service := NewVehicleService(mockRepo)

	expected := &model.Vehicle{
		Make: "Honda",
		Mark: "Civic",
		Year: 2020,
	}
	mockRepo.On("GetByID", uint(1)).Return(expected, nil)

	got, err := service.GetByID(uint(1))

	assert.Equal(t, expected, got)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_Failure(t *testing.T) {
	mockRepo := new(mocks.MockVehicleRepository)
	service := NewVehicleService(mockRepo)

	mockRepo.On("GetByID", uint(42)).Return((*model.Vehicle)(nil), errors.New("db error"))

	got, err := service.GetByID(uint(42))

	assert.Nil(t, got)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestListAll_Success(t *testing.T) {
	mockRepo := new(mocks.MockVehicleRepository)
	service := NewVehicleService(mockRepo)

	expected := model.Vehicle{
		Make: "Honda",
		Mark: "Civic",
		Year: 2020,
	}
	mockRepo.On("ListAll").Return([]model.Vehicle{expected}, nil)

	got, err := service.ListAll()

	assert.Len(t, got, 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestListAll_Failure(t *testing.T) {
	mockRepo := new(mocks.MockVehicleRepository)
	service := NewVehicleService(mockRepo)

	mockRepo.On("ListAll").Return([]model.Vehicle{}, errors.New("db error"))

	got, err := service.ListAll()

	assert.Len(t, got, 0)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
