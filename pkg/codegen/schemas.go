package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"path/filepath"
	"text/template"
)

func EmbedStruct(schemas []SchemaTemplateData, output io.Writer) error {
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

func SchemaToStruct(name string, s *openapi3.Schema) SchemaTemplateData {
	return SchemaTemplateData{
		Name: name,
		Properties: utils.Map(s.Properties, func(s string, v *openapi3.SchemaRef) PropertyTemplateData {
			return PropertyTemplateData{
				Name: asTitle(s),
				Type: oaSchemaRefToPrimitive(v),
			}
		}),
	}
}

// SchemaTemplateData contains the templating data necessary to generate the corresponding struct for each schema
type SchemaTemplateData struct {
	Name       string
	Properties []PropertyTemplateData
}

// PropertyTemplateData contains the templating data necessary to generate the corresponding primitive for each property of the schema
type PropertyTemplateData struct {
	Name string
	Type string
}

// Exports names. (first letter of the keyword as capital letter)
// 1. Allows access outside output package,
// 2. Avoids collision with builtin keywords (i.e. `type`)
func asTitle(s string) string {
	return cases.Title(language.English, cases.Compact).String(s)
}

func refPathToType(ref string) string {
	//paths := strings.Split(ref, "/")
	//return paths[len(paths)-1]
	return filepath.Base(ref)
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
	case "string":
		return "string"
	case "integer":
		return oaSchemaFormatToPrimitive(s.Value.Format) // int subtypes
	case "array":
		return "[]" + oaSchemaRefToPrimitive(s.Value.Items)
	case "object":
		return asTitle(refPathToType(s.Ref))
	default:
		return "interface{} //TODO: Handle others"
	}
}
