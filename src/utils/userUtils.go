package utils

import (
	"errors"
	"fmt"
	"net/smtp"
	"time"

	"segmenta/src/configs"
	"segmenta/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashing(password string) (string, error) {
	hashedBytes, errorHandler := bcrypt.GenerateFromPassword([]byte(password), 8)
	hashedPassword := string(hashedBytes)
	if errorHandler != nil {
		return "", errorHandler
	}
	return hashedPassword, nil
}

func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"name":    user.FullName,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.GetJWTSecretKey()))
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, errorHandler := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.GetJWTSecretKey()), nil
	})
	if errorHandler != nil {
		return nil, errorHandler
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func GetUserID(c *gin.Context, prefix string) (uint, bool) {
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		SendErrorResponse(c, prefix+" Unauthorized", 401)
		return 0, false
	}
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		SendErrorResponse(c, prefix+" Invalid user ID", 400)
		return 0, false
	}
	return uint(userIDFloat), true
}

func GenerateResetPasswordLink(email string) string {
	token, errorHandler := GenerateJWT(&models.User{
		Email: email,
	})
	if errorHandler != nil {
		return ""
	}
	
	if configs.LoadAppConfig().DomainName == "localhost" {
		return "http://" + configs.LoadAppConfig().DomainName + ":" + configs.LoadAppConfig().FrontendPort + "/reset-password?token=" + token
	}
	return "http://" + configs.LoadAppConfig().DomainName + "/reset-password?token=" + token
}

func SendResetPasswordtoEmail(email string, resetLink string) error {
	fromEmail := configs.LoadAppConfig().SMTPEmail
	fromPassword := configs.LoadAppConfig().SMTPEmailPassword
	smtpHost := configs.LoadAppConfig().SMTPHost
	smtpPort := configs.LoadAppConfig().SMTPPort

	message := []byte("Subject: Reset Password\n\n" +
		"Click the following link to reset your Segmenta password :\n" + resetLink)

	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, []string{email}, message)
}

func ParseResetPasswordToken(tokenString string) (string, error) {
	claims, errorHandler := ValidateJWT(tokenString)
	if errorHandler != nil {
		return "", errorHandler
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email claim not found in token")
	}

	return email, nil
}
