package lib

import (
	"errors"
	"gin-socmed/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var mySigningKey = []byte("mysecretkey")

type JWTClaims struct {
	ID *uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.User) (string, error) {
	claims := JWTClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (*uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalide token signature")
		}

		return nil, errors.New("your token has expired")
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok || !token.Valid {
		return nil, errors.New("your token was expired")
	}

	return claims.ID, nil
}
