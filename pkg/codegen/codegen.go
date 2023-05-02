package codegen

import (
	"bytes"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go/format"
	"go_fuzz_openapi/pkg/codegen/routes"
	"go_fuzz_openapi/pkg/codegen/schemas"
)

func Generate(url string, ss openapi3.Schemas, ps openapi3.Paths) ([]byte, error) {
	schemasTemplateData, err := schemas.ExtractSchemasTemplateData(ss)
	if err != nil {
		return nil, fmt.Errorf("failed generating schema template data: %w", err)
	}

	var buf bytes.Buffer
	if err = schemas.Generate(&buf, schemasTemplateData); err != nil {
		return nil, fmt.Errorf("failed generating schema code: %w", err)
	}

	routesData, err := routes.ExtractPathsTemplateData(url, ps)
	if err != nil {
		return nil, fmt.Errorf("failed generating paths template data: %w", err)
	}

	if err = routes.Generate(&buf, routesData); err != nil {
		return nil, err
	}

	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to format generated code: %w", err)
	}

	return formattedCode, nil
}
