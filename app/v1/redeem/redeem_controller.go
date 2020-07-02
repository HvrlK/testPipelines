package redeem

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"testPipelines/app/v1/redeem/model"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

func (controller *Controller) GetRedeem(c *gin.Context) {
	result := &model.Redeem{}
	controller.db.Preload("TotalVolume", &result.TotalVolume).
		Preload("Profit", &result.Profit).
		Where("Actor = ?", c.Param("actor")).
		First(result)
	c.JSON(200, result)
}