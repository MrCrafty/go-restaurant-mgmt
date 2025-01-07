package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/database"
	"github.com/mrcrafty/go-restaurant-mgmt/middleware"
	"github.com/mrcrafty/go-restaurant-mgmt/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	//Initializing User Routes before middleware
	routes.UserRoutes(router)

	//calling and using the authentication middleware
	router.Use(middleware.Authentication())

	//Initializing all the routes
	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	//starting the application on port
	router.Run(":" + port)
}
