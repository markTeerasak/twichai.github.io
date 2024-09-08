package entities

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name string
}
