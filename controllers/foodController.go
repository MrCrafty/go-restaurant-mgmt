package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrcrafty/go-restaurant-mgmt/database"
	"github.com/mrcrafty/go-restaurant-mgmt/models"
)

var foodCollection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// defer cancel()
		c.JSON(200, foodCollection)

	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodId := c.Param("id")
		var food models.Food
		err := foodCollection.FindOne(ctx, gin.H{"_id": foodId}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(200, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {
	return 0
}

func toFixed(num float64, precision int) float64 {
	return 0
}
