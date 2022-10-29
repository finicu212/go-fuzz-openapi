package schemas

import (
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"text/template"
)

func convertType(jsonType string) string {
	switch jsonType {
	case "string":
		return "string"
	case "integer":
		return "int32"
	case "array":
		return "interface{} //TODO: Handle arrays"
	case "object":
		return "interface{} //TODO: Handle objects"
	default:
		return "interface{} //TODO: Handle others"
	}
}

type SchemaTemplateData struct {
	Name       string
	Properties []PropertyTemplateData
}

type PropertyTemplateData struct {
	Name string
	Type string
}

func EmbedStruct(name string, s *openapi3.Schema) (string, error) {
	data := SchemaTemplateData{Name: name, Properties: nil}
	for pName, pRef := range s.Properties {
		p := pRef.Value
		property := PropertyTemplateData{Name: pName, Type: p.Type}
		data.Properties = append(data.Properties, property)
	}

	t, err := template.ParseFiles("pkg/templates/schemas.template")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	return "done", nil
}
