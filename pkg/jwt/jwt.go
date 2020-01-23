package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"time"
)

var (
	SecretPath = os.Getenv("SECRET_PATH")
	PublicPath = os.Getenv("PUBLIC_PATH")
)

func Create(userName, secretPath string) (string, error) {
	signBytes, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return "", err
	}
	// token, claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userName
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signBytes)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// TODO: implement validate jwt
func Validate() error {
	return nil
}

