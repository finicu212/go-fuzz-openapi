package cmd

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"go_fuzz_openapi/pkg/codegen"
	"go_fuzz_openapi/pkg/utils"
	"os"
)

const (
	flagSpec   = "spec"
	flagUrl    = "url"
	flagOutput = "output"
)

type testData struct {
	example     interface{}
	operationId string        // updatePetWithForm
	params      []interface{} // ID, Name, Status
}

/*
 * Path processing:
 * paths -> [
 *   - pet-put:
 * 	   - operationId: "updatePet": Use in fuzz test name
 *     - requestBody.content.(json/xml).schema: Use valid request content
 *     - responses: [200, 400, 404, 405]
 *       - response.200.content.(json/xml).schema: Get valid updated content
 *	   - security (relevant?)
 * ]
 *
 * Schema processing
 * 	 create struct for each object.
 *   map json primitives to golang primitives
 * s := range components.schemas -> [
 *   - Pet: {
 *     - s.type: object => create struct
 *       p := range properties -> [
 *       - p.type == "array": can be primitive, non-primitive in p.items.type
 *       - p.type == $ref:
 *       ]
 *
 *   }
 *
 * ]
 */

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

		f, err := utils.GetTestFileInstance(out + "/main_test.go")()
		if err != nil {
			return err
		}

		// Defer closure of file
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		var schemasAsStructs []utils.SchemaTemplateData
		for sName, sRef := range doc.Components.Schemas {
			s := sRef.Value
			json, err := sRef.MarshalJSON()
			if err != nil {
				return err
			}
			fmt.Printf("%s :: %s\n", sName, json)
			schemasAsStructs = append(schemasAsStructs, utils.SchemaToStruct(sName, s))
			if err != nil {
				return err
			}
			fmt.Printf("\n")
		}
		err = codegen.EmbedStruct(schemasAsStructs, f)

		err = doc.Validate(context.TODO())
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP(flagUrl, "u", "https://petstore3.swagger.io/api/v3/", "Specify an URL to target")
	generateCmd.Flags().StringP(flagSpec, "s", "openapi.yaml", "The swagger/openapi spec file of the API")
	generateCmd.Flags().StringP(flagOutput, "o", "out", "The output directory in which to generate the fuzz tests")

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
