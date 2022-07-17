package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenMaker struct {
	Key       []byte        // JWT secret key
	ValidTime time.Duration // JWT token valid time
}

type JWTClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

// Generates a JWT token for the user
func (j *TokenMaker) Generate(userIdPayload string) (t string, expiresAt int64, err error) {
	expirationTime := time.Now().Add(j.ValidTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserID: userIdPayload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	})

	// Create the JWT string
	tokenString, err := token.SignedString(j.Key)
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}

// func (*JWTToken) Verify(tokenString string) (userId string, e error) {
// }

// func (*JWTToken) Refresh(tokenString string) (t string, expiresAt int64, err error) {
// }

// TODO: Split jwt functions into internal methods
