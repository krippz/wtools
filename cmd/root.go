// Package cmd /*
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   string
	gitCommit string //nolint:gochecknoglobals
	buildDate string //nolint:gochecknoglobals
	appName   string //nolint:gochecknoglobals
)

// rootCmd represents the base command when called without any subcommands.
//
//goland:noinspection GoLinter
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals,exhaustruct,exhaustivestruct
	Use:     "wtools",
	Version: version,
	Short:   "Web helper tools",
	Long: `This tool is a collection of helpers. 
You can use them to help with some web related tasks, 
such as decoding a jwt token into plain text.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.SetVersionTemplate(fmt.Sprintf("%v version: %v \r\ngit commit: %v\r\nbuild-date: %v", appName, version, gitCommit, buildDate)) //nolint:lll
}
