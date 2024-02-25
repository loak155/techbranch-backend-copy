package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/loak155/techbranch-backend/internal/pkg/config"
)

type JwtManager struct {
	secret  string
	expires time.Duration
}

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func NewJwtManager(conf *config.Config) JwtManager {
	return JwtManager{
		secret:  conf.JWT.Secret,
		expires: conf.JWT.Expires,
	}
}

func (m *JwtManager) Generate(user_id int) (string, error) {
	claims := Claims{
		user_id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.expires).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *JwtManager) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("invalid token")
			}
			return []byte(m.secret), nil
		},
	)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
