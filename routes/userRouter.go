package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.Register())
	incomingRoutes.POST("/users/login", controllers.Login())
}
