package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"path/filepath"
	"text/template"
)

func asTitle(s string) string {
	return cases.Title(language.English, cases.Compact).String(s)
}

func getTypeFromRef(ref string) string {
	//paths := strings.Split(ref, "/")
	//return paths[len(paths)-1]
	return filepath.Base(ref)
}

func convertType(s *openapi3.SchemaRef) string {
	switch s.Value.Type {
	case "string":
		return "string"
	case "integer":
		return "int32"
	case "array":
		return "[]" + convertType(s.Value.Items)
	case "object":
		return asTitle(getTypeFromRef(s.Ref))
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

func EmbedStruct(name string, s *openapi3.Schema, output io.Writer) error {
	data := SchemaTemplateData{Name: name, Properties: nil}

	for pName, pRef := range s.Properties {
		property := PropertyTemplateData{Name: asTitle(pName), Type: convertType(pRef)}
		data.Properties = append(data.Properties, property)
	}

	t, err := template.ParseFiles("pkg/templates/schemas.template")
	if err != nil {
		return err
	}

	err = t.Execute(output, data)
	if err != nil {
		return err
	}
	return nil
}
