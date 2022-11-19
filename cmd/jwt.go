/*
Copyright © 2022 Kristofer Linnestjerna <krippz@krippz.se>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	jwtHelper "github.com/krippz/wtools/internal/jwt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	jwtToken string
	Plain    bool
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

		token, err := jwtHelper.GetJwtTokenFromString(jwtToken)
		if err != nil {
			errPrinter.Println("Could not parse string as JWT Token")
		}

		dataClaims, err := json.Marshal(token.Claims)
		if err != nil {
			errPrinter.Println("Could not convert Claims to bytes")
			os.Exit(1)
		}
		claims, err := jwtHelper.ConvertToJsonMap(&dataClaims)
		if err != nil {

			errPrinter.Println("Could not convert Claims to map")
			os.Exit(2)
		}

		prettyColorized, _ := jwtHelper.MapToColorizedJsonString(claims)
		prettyPlain, _ := jwtHelper.DataToJsonString(dataClaims)

		if Plain {
			fmt.Println(prettyPlain)
		} else {
			fmt.Println(prettyColorized)
		}
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jwtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	jwtCmd.Flags().BoolVarP(&Plain, "plain", "p", false, "Show output with no colorization")
}
