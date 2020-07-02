package model

import (
	"github.com/jinzhu/gorm"
	"testPipelines/app/v1/models"
)

type Redeem struct {
	gorm.Model
	Actor string
	IndexId string
	TotalVolume []models.TotalVolume `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
	Profit []models.Profit           `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
}

