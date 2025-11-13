package jwt

import (
	"errors"
	"gin-demo/config"
	"gin-demo/util/logging"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSignKey = []byte(config.JwtSignKey)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerToken(username string) (string, error) {
	claims := JwtClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "swift",
			Subject:   "fanjiale",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSignKey)

	return ss, err
}

func ParseToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		return jwtSignKey, nil
	})
	if err != nil {
		logging.Error(nil, "token解析失败")
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		logging.Warning(nil, "token不合法")
		return nil, errors.New("token不合法")
	}
}
