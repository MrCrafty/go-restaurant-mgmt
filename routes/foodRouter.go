package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func FoodRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/foods", controllers.GetFoods())
	incomingRoute.GET("/foods/:food_id", controllers.GetFood())
	incomingRoute.POST("/foods", controllers.CreateFood())
	incomingRoute.PATCH("/foods/:food_id", controllers.UpdateFood())
	incomingRoute.DELETE("/foods/:food_id", controllers.DeleteFood())
}
