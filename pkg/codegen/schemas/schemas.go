package schemas

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/utils"
	"io"
	"text/template"
)

func Generate(wr io.Writer, schemas []SchemaTemplateData) error {
	t, err := template.ParseFiles("templates/schemas.template")
	if err != nil {
		return fmt.Errorf("failed parsing schemas template: %w", err)
	}

	err = t.Execute(wr, schemas)
	if err != nil {
		return fmt.Errorf("failed executing schemas template: %w", err)
	}

	return nil
}

// SchemaTemplateData contains the templating data necessary to generate the corresponding struct for each schema
type SchemaTemplateData struct {
	Name       string
	Properties []PropertyTemplateData
}
type PropertyTemplateData struct {
	Name string
	Type string
}

func ExtractSchemasTemplateData(ss openapi3.Schemas) ([]SchemaTemplateData, error) {
	var schemasAsStructs = make([]SchemaTemplateData, 0)
	for sName, sRef := range ss {
		s := sRef.Value
		json, err := sRef.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("schema %s is not valid json: %w", sName, err)
		}
		fmt.Printf("%s :: %s\n", sName, json)
		schemasAsStructs = append(schemasAsStructs, extractSchemaTemplateData(sName, s))
		if err != nil {
			return nil, fmt.Errorf("failed appending schema to struct slice: %w", err)
		}
		fmt.Printf("\n")
	}
	return schemasAsStructs, nil
}

func extractSchemaTemplateData(name string, s *openapi3.Schema) SchemaTemplateData {
	return SchemaTemplateData{
		Name: name,
		Properties: utils.Map(s.Properties, func(s string, v *openapi3.SchemaRef) PropertyTemplateData {
			return PropertyTemplateData{
				Name: utils.AsTitle(s),
				Type: oaSchemaRefToPrimitive(v),
			}
		}),
	}
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
		return utils.AsTitle(utils.RefPathToType(s.Ref))
	default:
		return "interface{} //TODO: Handle others"
	}
}
