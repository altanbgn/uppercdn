package utils

import (
  "time"

	"upperfile.com/api/internal/config"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func NewAccessToken(ID string) (token string, err error) {
  claims := &UserClaims{
    ID: ID,
    StandardClaims: jwt.StandardClaims{
      IssuedAt: time.Now().Unix(),
      ExpiresAt: time.Now().Add(
        time.Duration(config.Env.JWT_EXPIRE) * time.Second,
      ).Unix(),
    },
  }

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(config.Env.JWT_SECRET_KEY))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, err := jwt.ParseWithClaims(
    accessToken,
    &UserClaims{},
    func(token *jwt.Token) (interface{}, error) {
      return []byte(config.Env.JWT_SECRET_KEY), nil
    },
  )

  if err != nil {
    return nil
  }

	return parsedAccessToken.Claims.(*UserClaims)
}
