package handlers

import (
	"segmenta/src/services"

	"github.com/gin-gonic/gin"
)

func GetOneUserByUserID(c *gin.Context) {
	services.GetOneUserByUserID(c)
}

func UpdateUserByUserID(c *gin.Context) {
	services.UpdateUserByUserID(c)
}

func DeleteUserByUserID(c *gin.Context) {
	services.DeleteUserByUserID(c)
}