package models

import (
	"gorm.io/gorm"
	"time"
)

type Hotel struct {
	gorm.Model
	Name     string    `db:"name" json:"name" validate:"required,lte=255"`
	CheckIn  time.Time `db:"CheckIn" json:"CheckIn" validate:"required"`
	CheckOut time.Time `db:"CheckOut" json:"CheckOut" validate:"required,lte=255"`
	About    string    `db:"About" json:"About" validate:"required,lte=255"`
	Logo     string    `db:"Logo" json:"Logo" validate:"required,lte=255"`
	Status   string    `db:"Status" json:"Status" validate:"required,lte=255"`
	Timezone string    `db:"Timezone" json:"Timezone" validate:"required,lte=255"`
}
