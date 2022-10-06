package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fuzzctl",
	Short: "fuzzctl is a black-box API fuzz test generator",
	Long: `A fuzz test generator, which aims to generate 
				gofuzz tests based on an API swagger file./`,
	Run: func(cmd *cobra.Command, args []string) {
	},
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
