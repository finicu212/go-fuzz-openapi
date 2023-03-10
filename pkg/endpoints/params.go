package endpoints

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/functional"
	"golang.org/x/exp/maps"
	"strings"
)

func FuzzPathParams(doc *openapi3.T, path string) string {
	p := doc.Paths.Find(path)
	ps := getPathParams(getSpecParams(p))
	// Build a map of name => fuzz value
	vals := map[string]string{}
	for _, param := range ps {
		fmt.Printf("Found param: %+v\n", param.Value.Name)
		vals[param.Value.Name] = getFakerFuncBySchema(param.Value.Schema)
	}

	// Replace path parameters with their fuzzed values
	for paramName, val := range vals {
		path = strings.Replace(path, "{"+paramName+"}", val, -1)
	}

	return path
}

func getFakerFuncBySchema(schema *openapi3.SchemaRef) string {
	switch schema.Value.Type {
	case "integer":
		return string(gofakeit.Int32())
	case "string":
		return "user1"
	}
	return ""
}

func getSpecParams(p *openapi3.PathItem) openapi3.Parameters {
	// Some specs set params on path-level
	if len(p.Parameters) > 0 {
		return p.Parameters
	}
	// Other specs set params on every operation
	return maps.Values(p.Operations())[0].Parameters
}

func getPathParams(ps openapi3.Parameters) openapi3.Parameters {
	return functional.Filter(ps, func(p *openapi3.ParameterRef) bool {
		return p.Value.In == "path"
	})
}
