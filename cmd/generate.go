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
	Use: fmt.Sprintf("generate --%s <file> --%s <url>", flagSpec, flagUrl),
	//Use:   "generate",
	Short: `Generate the fuzz tests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.PersistentFlags().GetString(flagUrl)
		swagger, _ := cmd.PersistentFlags().GetString(flagSpec)
		out, _ := cmd.PersistentFlags().GetString(flagOutput)
		fmt.Printf("Running generate with %s, %s, %s\n", url, out, swagger)

		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromFile(swagger)
		if err != nil {
			return err
		}

		err = doc.Validate(context.Background())
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

		code, err := codegen.Generate(url, doc.Components.Schemas, doc.Paths)
		if err != nil {
			return err
		}

		_, err = f.Write(code)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP(flagOutput, "o", "out", "The output directory in which to generate the fuzz tests")

}
