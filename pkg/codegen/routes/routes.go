package routes

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/codegen/schemas"
	"io"
	"text/template"
)

type RequestTemplateData struct {
	URL        string
	Name       string
	Paths      openapi3.Paths
	Properties []schemas.PropertyTemplateData
}

func Generate(wr io.Writer, routes []RequestTemplateData) error {
	t, err := template.ParseFiles("templates/schemas.template")
	if err != nil {
		return fmt.Errorf("failed parsing schemas template: %w", err)
	}

	funcs := template.FuncMap{"Operations": openapi3.PathItem{}.Operations}
	t.Funcs(funcs)

	err = t.Execute(wr, routes)
	if err != nil {
		return fmt.Errorf("failed executing schemas template: %w", err)
	}

	return nil
}
