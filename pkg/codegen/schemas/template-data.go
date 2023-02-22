package schemas

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/utils"
)

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
