package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTAuth struct {
	algorithm string
	secretKey []byte
	timeout   time.Duration
}

func NewJWTAuth(algorithm string, secretKey []byte, timeout time.Duration) JWTAuth {
	return JWTAuth{
		algorithm: algorithm,
		secretKey: secretKey,
		timeout:   timeout,
	}
}

func (auth JWTAuth) CreateToken(userId int64) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(auth.algorithm))

	token.Claims["userId"] = userId
	token.Claims["exp"] = time.Now().Add(auth.timeout).Unix()

	tokenValue, err := token.SignedString(auth.secretKey)

	if err != nil {
		return "", errors.New("unauthorized")
	}

	return tokenValue, nil
}
