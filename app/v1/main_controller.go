package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"testPipelines/app/v1/deposit"
	"testPipelines/app/v1/profit"
	"testPipelines/app/v1/redeem"
)

type MainController struct {
	DepositController  	*deposit.Controller
	RedeemController 	*redeem.Controller
	ProfitController	*profit.Controller
}

func NewMainController(engine *gin.Engine, db *gorm.DB) *MainController {
	return &MainController{
		DepositController:  deposit.NewController(db),
		RedeemController: 	redeem.NewController(db),
		ProfitController: 	profit.NewController(),
	}
}

func (controller *MainController) NoRoute(c *gin.Context) {
	errorMessage := fmt.Sprintf("Path not implemented: %s", c.Request.URL.String())
	c.JSON(http.StatusNotFound, errorMessage)
}