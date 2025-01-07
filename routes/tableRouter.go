package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func TableRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/table", controllers.GetTables())
	incomingRoute.GET("/table/:table_id", controllers.GetTable())
	incomingRoute.POST("/table", controllers.CreateTable())
	incomingRoute.PATCH("/table/:table_id", controllers.UpdateTable())
}
