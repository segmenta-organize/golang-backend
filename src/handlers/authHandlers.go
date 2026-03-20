package handlers

import (
	"segmenta/src/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	services.Register(c)
}

func Login(c *gin.Context) {
	services.Login(c)
}

func Logout(c *gin.Context) {
	services.Logout(c)
}

func Refresh(c *gin.Context) {
	services.Refresh(c)
}

func ForgotPassword(c *gin.Context) {
	services.ForgotPassword(c)
}

func ResetPassword(c *gin.Context) {
	services.ResetPassword(c)
}