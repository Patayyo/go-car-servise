package model

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	Make string `json:"make"`
	Mark string `json:"mark"`
	Year int    `json:"year"`
}
