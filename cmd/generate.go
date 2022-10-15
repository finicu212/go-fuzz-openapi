package cmd

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

const (
	flagSpec   = "spec"
	flagUrl    = "url"
	flagOutput = "output"
)

// URL: https://petstore3.swagger.io/api/v3/pet
// Swagger: https://petstore.swagger.io/v2/swagger.json
var generateCmd = &cobra.Command{
	Use:   fmt.Sprintf("generate --%s <file> --%s <url>", flagSpec, flagUrl),
	Short: `Generate the fuzz tests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString(flagUrl)
		swagger, _ := cmd.Flags().GetString(flagSpec)
		out, _ := cmd.Flags().GetString(flagOutput)
		fmt.Printf("Running generate with %s, %s, %s\n", url, out, swagger)

		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromFile(swagger)
		if err != nil {
			return err
		}

		err = doc.Validate(context.TODO())
		if err != nil {
			return err
		}

		return nil
	},
}

type PFlagFunc func(name string, shorthand string, value string, usage string) *string

func decoratePFlagFuncs(f PFlagFunc) PFlagFunc {
	return func(name string, shorthand string, value string, usage string) *string {
		fmt.Println("[decorate] before")
		ret := f(name, shorthand, value, usage)
		fmt.Println("[decorate] after")
		return ret
	}
}

var (
	ex = decoratePFlagFuncs(generateCmd.Flags().StringP)
)

func init() {
	rootCmd.AddCommand(generateCmd)
	decoratePFlagFuncs(generateCmd.Flags().StringP)(flagUrl, "u", "https://petstore3.swagger.io/api/v3/", "Specify an URL to target")
	generateCmd.Flags().StringP(flagSpec, "s", "openapi.json", "The swagger/openapi spec file of the API")
	generateCmd.Flags().StringP(flagOutput, "o", "output", "The output directory in which to generate the fuzz tests")

	if err := generateCmd.MarkFlagFilename(flagSpec); err != nil {
		return
	}
	if err := generateCmd.MarkFlagDirname(flagOutput); err != nil {
		return
	}
	//if err := generateCmd.MarkFlagRequired(flagSpec); err != nil {
	//	return
	//}
	//if err := generateCmd.MarkFlagRequired(flagUrl); err != nil {
	//	return
	//}
}
