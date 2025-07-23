package middleware

import (
	"fmt"
	"strings"
	"task_manager/data"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func( c *gin.Context){
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
	c.JSON(401, gin.H{"error": "Authorization header is required"})
	c.Abort()
	return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
	c.JSON(401, gin.H{"error": "Invalid authorization header"})
	c.Abort()
	return
	}

	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return data.JwtSecret, nil
	})

	if err != nil || !token.Valid {
	c.JSON(401, gin.H{"error": "Invalid JWT"})
	c.Abort()
	return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok{
		c.JSON(401, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}
	userID := claims["user_id"].(string)

	role := claims["role"].(string)
	c.Set("user_id", userID)
	c.Set("role", role)
		c.Next()
	}
}
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.JSON(403, gin.H{"error": "Forbidden: insufficient privileges"})
			c.Abort()
			return
		}
		c.Next()
	}
}