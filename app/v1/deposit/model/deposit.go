package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testPipelines/app/v1/models"
)

type Deposit struct {
	gorm.Model
	Uid string
	Actor string
	IndexId string
	TotalVolume []models.TotalVolume `gorm:"foreignkey:Uid;association_foreignkey:Uid"`
	Amount []models.Amount           `gorm:"foreignkey:Uid;association_foreignkey:Uid"`
	Profit []models.Profit           `gorm:"foreignkey:Uid;association_foreignkey:Uid"`
}

func MadeUid(actor string, indexId string) string {
	return fmt.Sprintf("%s_%s_deposit", actor, indexId)
}