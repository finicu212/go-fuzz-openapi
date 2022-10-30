package codegen

import (
	"go_fuzz_openapi/pkg/utils"
	"io"
	"text/template"
)

func EmbedStruct(schemas []utils.SchemaTemplateData, output io.Writer) error {
	t, err := template.ParseFiles("templates/schemas.template")
	if err != nil {
		return err
	}

	err = t.Execute(output, schemas)
	if err != nil {
		return err
	}
	return nil
}
