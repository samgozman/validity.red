package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type TokenMaker struct {
	Key []byte // JWT secret key
}

type JWTClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

// Generate - generates a JWT token for the user.
//
// maxAge - JWT token max age (in seconds)
func (j *TokenMaker) Generate(userID string, maxAge int) (t string, err error) {
	expirationTime := time.Now().Add(time.Duration(maxAge) * time.Second).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	})

	// Create the JWT string
	tokenString, err := token.SignedString(j.Key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verifies a JWT token and returns decoded UserId
func (j *TokenMaker) Verify(tokenString string) (userId string, e error) {
	claims, err := j.parse(tokenString)
	if err != nil {
		return "", err
	}

	return claims.UserID, nil
}

// Refresh - generates new JWT token for the user.
//
// maxAge - JWT token max age (in seconds)
func (j *TokenMaker) Refresh(tokenString string, maxAge int) (t string, err error) {
	claims, err := j.parse(tokenString)
	if err != nil {
		return "", err
	}

	// TODO: Ensure that a new token is not issued until enough time has passed
	// TODO: Return previous token if it's far from expired

	// Create new token with current payload
	return j.Generate(claims.UserID, maxAge)
}

// Parse token string and return decoded JWTClaims
func (j *TokenMaker) parse(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return j.Key, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil || !token.Valid || claims.UserID == "" {
		return &JWTClaims{}, ErrInvalidToken
	}

	return claims, nil
}
