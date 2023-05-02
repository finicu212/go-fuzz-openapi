package routes

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"go_fuzz_openapi/pkg/endpoints"
	"io"
	"text/template"
)

type FuzzTestTemplateData struct {
	URL              string
	PathTemplateData []PathTemplateData
}
type PathTemplateData struct {
	Name                  string
	OperationTemplateData []OperationTemplateData
}
type OperationTemplateData struct {
	Name         string
	OpId         string
	RequestBody  string
	ResponseBody string
}

func ExtractPathsTemplateData(url string, ps openapi3.Paths) (*FuzzTestTemplateData, error) {
	fuzzTestTemplateData := &FuzzTestTemplateData{
		URL:              url,
		PathTemplateData: []PathTemplateData{},
	}
	for pName, p := range ps {
		json, err := p.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("path %s is not valid json: %w", pName, err)
		}
		fmt.Printf("-(path)->%s :: %s\n", pName, json)

		pthTmplData, err := extractPathTemplateData(pName, p)
		fuzzTestTemplateData.PathTemplateData = append(fuzzTestTemplateData.PathTemplateData, *pthTmplData)
		if err != nil {
			return nil, fmt.Errorf("failed appending schema to struct slice: %w", err)
		}
		fmt.Printf("\n")
	}
	return fuzzTestTemplateData, nil
}

func extractPathTemplateData(name string, p *openapi3.PathItem) (*PathTemplateData, error) {
	pthTmplData := &PathTemplateData{Name: endpoints.RefPathToEndpoint(name)}
	pthTmplData.OperationTemplateData = make([]OperationTemplateData, 0)
	for opName, op := range p.Operations() {
		json, err := p.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("operation %s is not valid json: %w", opName, err)
		}
		fmt.Printf("\t-(operation)->%s :: %s\n", opName, json)
		pthTmplData.OperationTemplateData = append(pthTmplData.OperationTemplateData, extractOperationTemplateData(opName, op))
		if err != nil {
			return nil, fmt.Errorf("failed appending schema to struct slice: %w", err)
		}
		fmt.Printf("\n")
	}

	return pthTmplData, nil
}

func extractOperationTemplateData(name string, o *openapi3.Operation) OperationTemplateData {
	if o == nil {
		return OperationTemplateData{}
	}
	reqBody := ""
	if o.RequestBody != nil && o.RequestBody.Value != nil {
		reqBody = tryGetSchemaOfRequestBody(o.RequestBody.Value.Content)
	}
	respBody := ""
	validResp := o.Responses.Get(200)
	if validResp != nil && validResp.Value != nil {
		respBody = tryGetSchemaOfRequestBody(validResp.Value.Content)
	}
	fmt.Printf("\nOperation: %+v\n", o)
	return OperationTemplateData{
		Name:         name,
		OpId:         o.OperationID,
		RequestBody:  reqBody,  // Could be empty. Must be checked in templates
		ResponseBody: respBody, // Could be empty. Must be checked in templates
	}
}

func Generate(wr io.Writer, tmplData *FuzzTestTemplateData) error {
	t, err := template.ParseFiles("templates/request.template")
	if err != nil {
		return fmt.Errorf("failed parsing request template: %w", err)
	}

	err = t.Execute(wr, *tmplData)
	if err != nil {
		return fmt.Errorf("failed executing schemas template: %w", err)
	}

	return nil
}

// tryGetSchemaOfRequestBody attempts to find the name of the used schema as the RequestBody, by iterating through mime types used in swagger specs.
// The name of the schema used for the first non-empty mime type is returned
func tryGetSchemaOfRequestBody(content openapi3.Content) string {
	attemptedMimeTypesWhitelist := []string{"application/json", "application/xml", "application/x-www-form-urlencoded",
		"multipart/form-data", "text/plain", "text/plain; charset=utf-8", "text/html"}
	for _, mimeType := range attemptedMimeTypesWhitelist {
		mediaType := content.Get(mimeType)
		if mediaType != nil {
			return endpoints.RefPathToType(mediaType.Schema.Ref)
		}
	}

	return ""
}
