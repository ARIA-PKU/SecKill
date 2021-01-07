package common

import (
	"SecKill_Product/datamodels"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_create")

type Claims struct{
    Name      string
	Authority string
	Telephone string
 	jwt.StandardClaims
}

//发放token
func ReleaseToken(user datamodels.Admin) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Name: user.Name,
		Authority: user.Authority,
		Telephone: user.Telephone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "XUYAO",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
