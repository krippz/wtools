package jwthelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TylerBrock/colorjson"
	"github.com/golang-jwt/jwt"
)

var ErrUnableToParse = errors.New("could not parse jwt token string")

func GetJwtTokenFromString(token string) (*jwt.Token, error) {
	jwtToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return nil, ErrUnableToParse
	})

	return jwtToken, nil
}

func DataToJSONString(data []byte) (string, error) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, data, "", "    ")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

func MapToColorizedJSONString(data map[string]interface{}) (string, error) {
	formatter := colorjson.NewFormatter()
	formatter.Indent = 4

	colorizedData, err := formatter.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(colorizedData), nil
}

func ConvertToJSONMap(data *[]byte) (map[string]interface{}, error) {
	var mappedClaims map[string]interface{}

	err := json.Unmarshal(*data, &mappedClaims)
	if err != nil {
		return nil, err
	}

	return mappedClaims, nil
}

//goland:noinspection GoUnusedExportedFunction
func IterMap(claims *map[string]interface{}) {
	for key, value := range *claims {
		fmt.Println("key:", key, "=>", "value:", value) //nolint:forbidigo
	}
}
