package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ItemsbyOrder(id string) (OrderItems []primitive.M, err error) {
	items := make([]primitive.M, 0)
	return items, nil
}
