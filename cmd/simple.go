package cmd

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"go_fuzz_openapi/pkg/endpoints"
)

func simpleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simple",
		Short: `Test cmd`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			//url, _ := cmd.Flags().GetString(flagUrl)
			swagger, _ := cmd.Flags().GetString(flagSpec)
			//out, _ := cmd.Flags().GetString(flagOutput)

			loader := openapi3.NewLoader()
			doc, err := loader.LoadFromFile(swagger)
			if err != nil {
				return err
			}

			err = doc.Validate(context.Background())
			if err != nil {
				return err
			}

			path := endpoints.FuzzPathParams(doc, "/pet/{petId}")
			fmt.Printf("%+v\n", path)

			return nil
		},
	}

	cmd.Flags().SortFlags = false
	cmd.SilenceUsage = true

	return cmd
}

func init() {
	rootCmd.AddCommand(simpleCmd())
}
