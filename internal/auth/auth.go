package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	secretKey []byte
}

func NewTokenManager(secret string) (*TokenManager, error) {
	if secret == "" {
		return nil, errors.New("JWT_SECRET must be set")
	}
	return &TokenManager{
		secretKey: []byte(secret),
	}, nil
}

func (tm *TokenManager) CreateToken(host string, port int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"port": port,
			"host": host,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(tm.secretKey)
}

func (tm *TokenManager) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected Signing Method: %v", token.Header["alg"])
		}
		return tm.secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
