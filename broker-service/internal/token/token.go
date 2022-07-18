package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("invalid token")
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

// Verifies a JWT token and returns decoded UserId
func (j *TokenMaker) Verify(tokenString string) (userId string, e error) {
	claims := JWTClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return j.Key, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)
	if err != nil || !token.Valid || claims.UserID == "" {
		return "", ErrInvalidToken
	}

	return claims.UserID, nil
}

// func (*JWTToken) Refresh(tokenString string) (t string, expiresAt int64, err error) {
// }

// TODO: Split jwt functions into internal methods
