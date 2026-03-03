package utils

import (
	"fmt"
	"time"

	"segmenta/src/configs"
	"segmenta/src/models"

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

func getUserID(c *gin.Context, prefix string) (uint, bool) {
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		utils.SendErrorResponse(c, prefix+" Unauthorized", 401)
		return 0, false
	}
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		utils.SendErrorResponse(c, prefix+" Invalid user ID", 400)
		return 0, false
	}
	return uint(userIDFloat), true
}