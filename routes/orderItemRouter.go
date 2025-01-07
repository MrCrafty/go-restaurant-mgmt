package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func OrderItemRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/orderItems", controllers.GetOrderItems())
	incomingRoute.GET("/orderItem/:orderitem_id", controllers.GetOrderItem())
	incomingRoute.GET("/orderItems-order/:orderitem_id", controllers.GetOrderItemsByOrder())
	incomingRoute.POST("/orderItem", controllers.CreateOrderItem())
	incomingRoute.PATCH("/orderItem/:orderitem_id", controllers.UpdateOrderItem())
}
