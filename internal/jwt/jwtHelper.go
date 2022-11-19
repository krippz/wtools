package jwtHelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/golang-jwt/jwt"
)

func GetJwtTokenFromString(token string) (*jwt.Token, error) {
	jwtToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return nil, errors.New("Could not parse jwt token string")
	})

	return jwtToken, nil
}

func DataToJsonString(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "    "); err != nil {
		return "", nil
	}
	return prettyJSON.String(), nil
}

func MapToColorizedJsonString(data map[string]interface{}) (string, error) {

	formatter := colorjson.NewFormatter()
	formatter.Indent = 4
	colorizedData, err := formatter.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(colorizedData), nil
}

func ConvertToJsonMap(data *[]byte) (map[string]interface{}, error) {

	var mappedClaims map[string]interface{}
	err := json.Unmarshal([]byte(*data), &mappedClaims)
	if err != nil {
		return nil, err
	}

	return mappedClaims, nil
}

func IterMap(claims *map[string]interface{}) {
	for key, value := range *claims {
		fmt.Println("key:", key, "=>", "value:", value)
	}
}
