package Infrastructure

import(
		"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndComparePassword_Success(t *testing.T) {
	password := "mysecretpassword"

	
	hashed, err := Hashpassword(password)
	assert.Nil(t, err, "Hashing should not return an error")
	assert.NotEmpty(t, hashed, "Hashed password should not be empty")
	assert.NotEqual(t, password, hashed, "Hashed password should differ from plain password")

	
	err = ComparePassword(hashed, password)
	assert.Nil(t, err, "ComparePassword should succeed for correct password")
}

func TestComparePassword_Failure(t *testing.T) {
	password := "mysecretpassword"
	wrongPassword := "wrongpassword"

	hashed, _ := Hashpassword(password)

	err := ComparePassword(hashed, wrongPassword)
	assert.NotNil(t, err, "ComparePassword should return error for wrong password")
}

func TestHashAndCompare_EmptyPassword(t *testing.T) {
	password := ""

	hashed, err := Hashpassword(password)
	assert.Nil(t, err, "Hashing empty password should not return an error")
	assert.NotEmpty(t, hashed, "Hashed password should not be empty")

	err = ComparePassword(hashed, password)
	assert.Nil(t, err, "ComparePassword should succeed for empty password")
}
