package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func OrderRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/order", controllers.GetOrders())
	incomingRoute.GET("/order/:order_id", controllers.GetOrder())
	incomingRoute.POST("/order", controllers.CreateOrder())
	incomingRoute.PATCH("/order/:order_id", controllers.UpdateOrder())
}
