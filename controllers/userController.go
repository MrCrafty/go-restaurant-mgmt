package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {
	hashP := ""
	return hashP
}

func VerifyPassword(password string, hash string) (bool, string) {

	return false, ""
}
