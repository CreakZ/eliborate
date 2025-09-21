package utils

import (
	"eliborate/pkg/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JWT interface {
	CreateToken(id int, isAdmin bool) string
	Authorize(token string) (claim, bool, error)
}

type claim struct {
	ID      int  `json:"id"`
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

func (c claim) Valid() error {
	if c.ID < 1 {
		return fmt.Errorf("id value is less than 1")
	}
	return c.RegisteredClaims.Valid()
}

type jwtUtil struct {
	expiresAt time.Duration
	secret    string
}

func InitJWTUtil() JWT {
	return jwtUtil{
		expiresAt: time.Duration(viper.GetInt(config.JWTAccessExpire)) * time.Minute,
		secret:    viper.GetString(config.JWTSecret),
	}
}

func (j jwtUtil) CreateToken(id int, isAdmin bool) string {
	expiresAt := time.Now().Add(j.expiresAt)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim{
		ID:      id,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expiresAt,
			},
		},
	})

	accessSigned, _ := accessToken.SignedString([]byte(j.secret))

	return accessSigned
}

func (j jwtUtil) Authorize(token string) (claim, bool, error) {
	var userClaim claim

	jwtToken, err := jwt.ParseWithClaims(token, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return claim{}, false, err
	}

	if !jwtToken.Valid {
		return claim{}, false, nil
	}

	return userClaim, true, nil
}
