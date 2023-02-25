package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	flagSpec   = "spec"
	flagUrl    = "url"
	flagOutput = "output"
)

var rootCmd = &cobra.Command{
	Use:   "fuzzctl",
	Short: "fuzzctl - generate API fuzz tests",
	Long:  `fuzzctl is a black-box API fuzz test generator which aims to generate gofuzz tests based on an API swagger file.`,
}

func init() {
	rootCmd.PersistentFlags().StringP(flagUrl, "u", "https://petstore3.swagger.io/api/v3/", "Specify an URL to target")
	rootCmd.PersistentFlags().StringP(flagSpec, "s", "openapi.yaml", "The swagger/openapi spec file of the API")

	rootCmd.PersistentFlags().SortFlags = false
	if err := rootCmd.MarkFlagFilename(flagSpec); err != nil {
		return
	}
	if err := rootCmd.MarkFlagDirname(flagOutput); err != nil {
		return
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
