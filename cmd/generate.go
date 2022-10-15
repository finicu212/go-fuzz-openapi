package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

const (
	flagSpec   = "spec"
	flagUrl    = "url"
	flagOutput = "output"
)

type testData struct {
	example     interface{}
	operationId string
}

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

		//fmt.Printf("Schemas:\n\n")
		//for _, schema := range doc.Components.Schemas {
		//	schema.
		//	props := schema.Value.Properties
		//	//json, _ := json.Marshal(props)
		//	//fmt.Printf("\t%s\n", string(json))
		//	for _, p := range props {
		//		fmt.Printf("\t%+v\n", p)
		//	}
		//
		//}

		for _, path := range doc.Paths {
			fmt.Printf("Debugging %s", path.Description)
			fmt.Printf("Path: %+v\n", path)
			for _, op := range path.Operations() {
				fmt.Printf("\tDebugging %s\n", op.OperationID)
				fmt.Printf("\tOperation: %+v\n", op)
				fmt.Printf("\tParams:\n")
				for _, param := range op.Parameters {
					json, _ := json.Marshal(param)
					fmt.Printf("\t\t- %s\n", json)
				}
				fmt.Println()

			}
			fmt.Println()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP(flagUrl, "u", "https://petstore3.swagger.io/api/v3/", "Specify an URL to target")
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
