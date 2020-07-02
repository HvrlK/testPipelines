package models

import "github.com/jinzhu/gorm"

type TotalVolume struct {
	gorm.Model
	TotalVolume string
	IndexId string
}
