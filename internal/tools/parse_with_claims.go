package tools

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"service_topic/config"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID     uuid.UUID
	UserStatus string
	UserPerm   int
}

func ParseTokenClaims(c *gin.Context) (Claims, error) {
	claims := Claims{}

	strToken := c.Request.Header.Get("Authorization")

	token, err := jwt.ParseWithClaims(strToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("не подходящий алгоритм шифрования: %v", t.Header["alg"])
		}
		return []byte(config.Env.SecretKey), nil
	})
	if err != nil {
		return Claims{}, err
	}

	if !token.Valid {
		return Claims{}, errors.New("токен не валиден")
	}

	return claims, err
}
