package handlers

import (
	"segmenta/src/controllers"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	controllers.Register(c)
}

func Login(c *gin.Context) {
	controllers.Login(c)
}

func Logout(c *gin.Context) {
	controllers.Logout(c)
}

func Refresh(c *gin.Context) {
	controllers.Refresh(c)
}

func ForgotPassword(c *gin.Context) {
	controllers.ForgotPassword(c)
}

func ResetPassword(c *gin.Context) {
	controllers.ResetPassword(c)
}