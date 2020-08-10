package jwt

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/projectpathap"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

// New JWT tokenの発行
func New(txTime time.Time, email string, pass string) string {

	signBytes, err := ioutil.ReadFile(projectpathap.GetRoot() + "/config/jwt.rsa")

	chk.SE(err)

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	chk.SE(err)

	// create token
	token := jwt.New(jwt.SigningMethodRS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = email
	claims["exp"] = txTime.Add(10 * time.Hour).Unix()

	tokenString, err := token.SignedString(signKey)
	chk.SE(err)

	return tokenString
}

// Auth JWT token認証
func Auth(r *http.Request) *jwt.Token {

	verifyBytes, err := ioutil.ReadFile(projectpathap.GetRoot() + "/config/jwt.rsa.pub.pkcs8")
	chk.SE(err)

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	chk.SE(err)

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil

	})

	chk.SE(err)

	if !token.Valid {
		chk.SE(errors.New("tokenが違います"))
	}

	return token
}
