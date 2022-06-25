package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const validTime = 10 * time.Minute

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Generates a JWT token for the user
func (u *User) GenerateJwtToken() (t string, expiresAt int64, err error) {
	expirationTime := time.Now().Add(validTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	})

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}
