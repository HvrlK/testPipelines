package deposit

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller      *Controller
	getDeposit        gin.IRoutes
}

func NewRouter(controller *Controller, group *gin.RouterGroup) *Router {
	return &Router{
		controller,
		group.GET("/:actor", controller.GetDeposit),
	}
}
