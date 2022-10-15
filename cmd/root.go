package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fuzzctl",
	Short: "fuzzctl - generate API fuzz tests",
	Long:  `fuzzctl is a black-box API fuzz test generator which aims to generate gofuzz tests based on an API swagger file.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
}
