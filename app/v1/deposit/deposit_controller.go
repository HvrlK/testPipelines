package deposit

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"testPipelines/app/v1/deposit/model"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

func (controller *Controller) GetDeposit(c *gin.Context) {
	result := &model.Deposit{}
	controller.db.Preload("TotalVolume", &result.TotalVolume).
		Preload("Amount", &result.Amount).
		Preload("Profit", &result.Profit).
		Where("Actor = ?", c.Param("actor")).
		First(result)
	c.JSON(200, result)
}