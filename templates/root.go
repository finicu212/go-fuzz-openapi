package templates

import (
	"net/http"
	"os"
	"text/template"
)

type config struct {
	name       string
	httpMethod string
}

func Execute() {
	t, err := template.ParseFiles("templates/fuzzTest.template")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("root_test.go")
	defer f.Close()

	cfg := config{
		name:       "Reverse",
		httpMethod: http.MethodGet,
	}

	err = t.Execute(f, cfg)
	if err != nil {
		panic(err)
	}
}
