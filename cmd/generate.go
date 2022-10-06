package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: `Generate the fuzz tests`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello world!\n")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("url", "u", "", "Specify an URL to target")
	generateCmd.Flags().StringP("swagger", "s", "", "The swagger of the API")
	generateCmd.Flags().StringP("output", "o", "", "The output directory in which to generate the fuzz tests")
}
