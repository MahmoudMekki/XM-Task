package utils

import (
	"fmt"
	"github.com/MahmoudMekki/XM-Task/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type AuthCustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID int64) string {
	claims := &AuthCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    config.GetEnvVar("AUTH_ISSUER"),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(config.GetEnvVar("AUTH_SECRET")))
	if err != nil {
		panic(err)
	}
	return signedToken
}
func GetToken(encodedToken string) (*jwt.Token, error) {
	claims := &AuthCustomClaims{}
	return jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(config.GetEnvVar("AUTH_SECRET")), nil
	})
}
