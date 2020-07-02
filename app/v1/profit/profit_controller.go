package profit

import "github.com/gin-gonic/gin"

type Controller struct {}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) GetProfit(c *gin.Context) {
	c.JSON(200, gin.H{
		"profit": "100",
	})
}