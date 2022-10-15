package cmd

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

// URL: https://petstore.swagger.io/v2/
// Swagger: https://petstore.swagger.io/v2/swagger.json
var generateCmd = &cobra.Command{
	Use:   "generate --swagger <file> --url <url>",
	Short: `Generate the fuzz tests`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		swagger, _ := cmd.Flags().GetString("swagger")
		out, _ := cmd.Flags().GetString("output")
		fmt.Printf("Running generate with %s, %s, %s\n", url, out, swagger)
		_, _ = openapi3.NewLoader().LoadFromFile("swagger.json")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("url", "u", "", "Specify an URL to target")
	generateCmd.Flags().StringP("swagger", "s", "", "The swagger of the API")
	generateCmd.Flags().StringP("output", "o", "", "The output directory in which to generate the fuzz tests")

	if err := generateCmd.MarkFlagFilename("swagger"); err != nil {
		return
	}
	if err := generateCmd.MarkFlagDirname("output"); err != nil {
		return
	}
	if err := generateCmd.MarkFlagRequired("swagger"); err != nil {
		return
	}
	if err := generateCmd.MarkFlagRequired("url"); err != nil {
		return
	}
}
