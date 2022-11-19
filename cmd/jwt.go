/*
Copyright © 2022 Kristofer Linnestjerna <krippz@krippz.se>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
)

var (
	jwtToken string
)

// jwtCmd represents the jwt command
// jsonClaims
var jwtCmd = &cobra.Command{
	Use:   "jwt [JWT-TOKEN]",
	Short: "Decode a jwt token to plain json",
	Long: `Decode a jwt token and display the claims it contains in the terminal. For example:
	
	wtools jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
	Args:          cobra.MaximumNArgs(2),
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {

		errPrinter := color.New(color.FgRed).Add(color.Underline)
		if len(args) == 1 {
			jwtToken = args[0]
		}
		if len(args) == 2 {
			jwtToken = args[1]
		}

		var match, _ = regexp.MatchString("(^[\\w-]*\\.[\\w-]*\\.[\\w-]*$)", jwtToken)
		if !match {
			fmt.Println()
			_, _ = errPrinter.Println("No a valid jwt token was supplied")
			fmt.Println()
			_ = cmd.Help()
			return
		}

		token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})

		dataClaims, err := json.Marshal(token.Claims)
		if err != nil {
			errPrinter.Println("Could not convert Claims to bytes")
			os.Exit(1)
		}
		claims, err := convertToJsonMap(&dataClaims)
		if err != nil {

			errPrinter.Println("Could not convert Claims to map")
			os.Exit(2)
		}
		
		jsonClaims, _ := mapToColorizedJsonString(claims)
		pretty, _ := dataToJsonString(dataClaims)

		fmt.Println(pretty)
		fmt.Println(jsonClaims)
	},
}

func dataToJsonString(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "    "); err != nil {
		return "", nil
	}
	return prettyJSON.String(), nil
}

func mapToColorizedJsonString(data map[string]interface{}) (string, error) {

	formatter := colorjson.NewFormatter()
	formatter.Indent = 4
	colorizedData, err := formatter.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(colorizedData), nil
}

func iterMap(claims *map[string]interface{}) {
	for key, value := range *claims {
		fmt.Println("key:", key, "=>", "value:", value)
	}
}

func convertToJsonMap(data *[]byte) (map[string]interface{}, error) {

	var mappedClaims map[string]interface{}
	err := json.Unmarshal([]byte(*data), &mappedClaims)
	if err != nil {
		return nil, err
	}

	return mappedClaims, nil
}

func init() {
	rootCmd.AddCommand(jwtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jwtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jwtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
