package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/controllers"
)

func MenuRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/menus", controllers.GetMenus())
	incomingRoute.GET("/menus/:menu_id", controllers.GetMenu())
	incomingRoute.POST("/menus", controllers.CreateMenu())
	incomingRoute.PATCH("/menus/:menu_id", controllers.UpdateMenu())
	incomingRoute.DELETE("/menus/:menu_id", controllers.DeleteMenu())
}
