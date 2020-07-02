package profit

import "github.com/gin-gonic/gin"

type Router struct {
	controller      *Controller
	getRedeem        gin.IRoutes
}

func NewRouter(controller *Controller, group *gin.RouterGroup) *Router {
	return &Router{
		controller,
		group.GET("/:sender/:indexId", controller.GetProfit),
	}
}
