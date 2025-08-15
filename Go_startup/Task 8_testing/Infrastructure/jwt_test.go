package Infrastructure

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	userID := "123"
	role := "admin"

	token, err := GenerateToken(userID, role)

	assert.Nil(t, err, "should not return an error")
	assert.NotEmpty(t, token, "token should not be empty")
}

func TestValidateToken_Success(t *testing.T) {
	userID := "123"
	role := "admin"

	
	token, _ := GenerateToken(userID, role)

	claims, err := ValidateToken(token)

	assert.Nil(t, err, "should not return an error")
	assert.Equal(t, userID, claims["user_id"])
	assert.Equal(t, role, claims["role"])
}

func TestValidateToken_InvalidToken(t *testing.T) {
	invalidToken := "this"

	claims, err := ValidateToken(invalidToken)

	assert.NotNil(t, err, "should return an error")
	assert.Nil(t, claims, "claims should be nil")
}

func TestValidateToken_Expired(t *testing.T) {
	expiredToken := func() string {
		claims := jwt.MapClaims{
			"user_id": "123",
			"role":    "admin",
			"exp":     time.Now().Add(-1 * time.Second).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tkn, _ := token.SignedString(jwtKey)
		return tkn
	}()

	claims, err := ValidateToken(expiredToken)

	assert.NotNil(t, err, "should return an error for expired token")
	assert.Nil(t, claims, "claims should be nil for expired token")
}