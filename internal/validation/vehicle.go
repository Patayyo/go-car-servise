package validation

import (
	"car-service/internal/model"
	"errors"
	"time"
)

func ValidateVehicle(v *model.Vehicle) error {
	currentYear := time.Now().Year()

	if v.Make != "" && v.Mark != "" && v.Year >= 1886 && v.Year <= currentYear {
		return nil
	}
	return errors.New("invalid input")
}
