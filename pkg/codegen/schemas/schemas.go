package schemas

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/utils"
	"io"
	"strings"
	"text/template"
)

func Generate(wr io.Writer, schemas []SchemaTemplateData) error {
	fs := template.FuncMap{
		"ToLower": strings.ToLower,
	}
	schTempl := template.New("schemas.template").Funcs(fs)

	t, err := schTempl.ParseFiles("templates/schemas.template")
	if err != nil {
		return fmt.Errorf("failed parsing schemas template: %w", err)
	}

	err = t.Execute(wr, schemas)
	if err != nil {
		return fmt.Errorf("failed executing schemas template: %w", err)
	}

	return nil
}

func oaSchemaFormatToPrimitive(format string) string {
	switch format {
	case "uint8":
	case "uint16":
	case "uint32":
	case "uint64":
	case "int8":
	case "int16":
	case "int32":
	case "int64":
	case "int":
		return format
	default:
	}
	return "int32"
}

func oaSchemaRefToPrimitive(s *openapi3.SchemaRef) string {
	switch s.Value.Type {
	case "bool":
	case "boolean":
		return "bool"
	case "string":
		return "string"
	case "integer":
		return oaSchemaFormatToPrimitive(s.Value.Format) // int subtypes
	case "array":
		return "[]" + oaSchemaRefToPrimitive(s.Value.Items) + "`fakesize:\"2\"`" // TODO: fakesize to func
	case "object":
		return utils.AsTitle(utils.RefPathToType(s.Ref))
	default:
	}
	return "interface{} // TODO: add this type to oaSchemaRefToPrimitive!"
}
