package models

import "github.com/jinzhu/gorm"

type Amount struct {
	gorm.Model
	Amount string
	IndexId string
}
