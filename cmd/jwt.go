// Package cmd /*
package cmd

//goland:noinspection SpellCheckingInspection

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/fatih/color"
	jwthelper "github.com/krippz/wtools/internal/jwt"
	"github.com/spf13/cobra"
)

var (
	Plain bool //nolint:gochecknoglobals
)

const one, two = 1, 2

// jwtCmd represents the jwt command.
//
//goland:noinspection ALL
var jwtCmd = &cobra.Command{ //nolint:exhaustruct,exhaustivestruct
	Use:   "jwt [JWT-TOKEN]",
	Short: "Decode a jwt token to plain json",
	Long: `Decode a jwt token and display the claims it contains in the terminal. For example:
	
	wtools jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`, //nolint:lll
	Args:          cobra.MaximumNArgs(two),
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {

		errPrinter := color.New(color.FgRed).Add(color.Underline)

		var jwtToken string
		if len(args) == one {
			jwtToken = args[0]
		}
		if len(args) == two {
			jwtToken = args[1]
		}

		var match, _ = regexp.MatchString("(^[\\w-]*\\.[\\w-]*\\.[\\w-]*$)", jwtToken)
		if !match {
			_, _ = errPrinter.Println("No a valid jwt token was supplied")
			_ = cmd.Help()

			return
		}

		token, err := jwthelper.GetJwtTokenFromString(jwtToken)
		if err != nil {
			_, _ = errPrinter.Println("Could not parse string as JWT Token")
			log.Fatal(err)
		}

		dataClaims, err := json.Marshal(token.Claims)
		if err != nil {
			_, _ = errPrinter.Println("Could not convert Claims to bytes")
			log.Fatal(err)
		}

		claims, err := jwthelper.ConvertToJSONMap(&dataClaims)
		if err != nil {
			_, _ = errPrinter.Println("Could not convert Claims to map")
			log.Fatal(err)
		}

		prettyColorized, _ := jwthelper.MapToColorizedJSONString(claims)
		prettyPlain, _ := jwthelper.DataToJSONString(dataClaims)

		if Plain {
			fmt.Println(prettyPlain) //nolint:forbidigo
		} else {
			fmt.Println(prettyColorized) //nolint:forbidigo
		}
	},
}

func init() {
	rootCmd.AddCommand(jwtCmd)
	jwtCmd.Flags().BoolVarP(&Plain, "plain", "p", false, "Show output with no colorization")
}
