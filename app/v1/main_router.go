package v1

import (
	"github.com/gin-gonic/gin"
	"testPipelines/app/v1/deposit"
	"testPipelines/app/v1/profit"
	"testPipelines/app/v1/redeem"
)

type MainRouter struct {
	Deposit *deposit.Router
	Redeem  *redeem.Router
	Profit 	*profit.Router
}

func NewMainRouter(controller *MainController, engine *gin.Engine) *MainRouter {

	engine.NoRoute(controller.NoRoute)

	v1 := engine.Group("/api/v1")

	depositRouter := deposit.NewRouter(controller.DepositController, v1.Group("/deposit"))

	redeemRouter := redeem.NewRouter(controller.RedeemController, v1.Group("/redeem"))

	profitRouter := profit.NewRouter(controller.ProfitController, v1.Group("/profit"))

	return &MainRouter{
		Deposit: 	depositRouter,
		Redeem: 	redeemRouter,
		Profit: 	profitRouter,
	}
}
