package jwt

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/projectpathap"
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

var signKey *rsa.PrivateKey
var verifyKey *rsa.PublicKey

// Setup setup
func Setup() {

	// privateKey
	signBytes, err := ioutil.ReadFile(projectpathap.GetRoot() + "/config/jwt.rsa")
	chk.SE(err)
	s, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	chk.SE(err)
	signKey = s

	// publicKey
	verifyBytes, err := ioutil.ReadFile(projectpathap.GetRoot() + "/config/jwt.rsa.pub.pkcs8")
	chk.SE(err)
	v, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	chk.SE(err)
	verifyKey = v

}

// New JWT tokenの発行
func New(txTime time.Time, mid string, email string) string {

	// create token
	token := jwt.New(jwt.SigningMethodRS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = email
	claims["mid"] = mid
	claims["exp"] = txTime.Add(10 * time.Hour).Unix()

	tokenString, err := token.SignedString(signKey)
	chk.SE(err)

	return tokenString
}

// Auth JWT token認証
func Auth(r *http.Request) *jwt.Token {

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
