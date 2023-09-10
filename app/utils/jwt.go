package utils

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// metadata for your jwt
type MetaToken struct {
	ID            int
	Email         string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecrePublicKeyEnvName string, ExpiredAt time.Duration) (string, error) {

	expiredAt := time.Now().Add(time.Duration(time.Second) * ExpiredAt).Unix()

	jwtSecretKey := SecrePublicKeyEnvName

	// metadata for your jwt
	claims := jwt.MapClaims{}
	claims["expiredAt"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")

	if tokenHeader == "" {
		return nil, nil
	}
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]

	if accessToken == "" {
		return nil, nil
	}
	jwtSecretKey := SecrePublicKeyEnvName

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func VerifyToken(accessToken, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := SecrePublicKeyEnvName

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}

func GetUserFromToken(c *gin.Context) string {

	getToken, _ := VerifyTokenHeader(c, viper.GetString("secret.user"))

	decodeToken := DecodeToken(getToken)
	return decodeToken.Claims.Email
}
