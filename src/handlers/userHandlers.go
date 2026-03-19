package handlers

import (
	"segmenta/src/controllers"

	"github.com/gin-gonic/gin"
)

func GetOneUserByUserID(c *gin.Context) {
	controllers.GetOneUserByUserID(c)
}

func UpdateUserByUserID(c *gin.Context) {
	controllers.UpdateUserByUserID(c)
}

func DeleteUserByUserID(c *gin.Context) {
	controllers.DeleteUserByUserID(c)
}