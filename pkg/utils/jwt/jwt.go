package jwt

import (
	"time"
	"yurii-lib/pkg/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JWT interface {
	CreateToken(id int, isAdmin bool) string
	Authorize(token string) (claim, bool, error)
}

type claim struct {
	ID      int
	IsAdmin bool
	jwt.RegisteredClaims
}

type JWTUtil struct {
	expiration time.Duration
	secret     string
}

func InitJWTUtil() JWT {
	return JWTUtil{
		expiration: time.Duration(viper.GetInt(config.JWTExpire)) * time.Hour,
		secret:     viper.GetString(config.JWTSecret),
	}
}

func (j JWTUtil) CreateToken(id int, isAdmin bool) string {
	expires := time.Now().Add(j.expiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim{
		ID:      id,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expires,
			},
		},
	})

	signed, _ := token.SignedString([]byte(j.secret))

	return signed
}

func (j JWTUtil) Authorize(token string) (claim, bool, error) {
	var userClaim claim

	jwtToken, err := jwt.ParseWithClaims(token, userClaim, func(t *jwt.Token) (interface{}, error) {
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
