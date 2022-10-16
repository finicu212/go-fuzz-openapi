package templates

import (
	"os"
	"text/template"
)

type config struct {
	Name string
}

func Execute() {
	t, err := template.ParseFiles("templates/fuzzTest.template")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("root_test.go")
	defer f.Close()

	cfg := config{Name: "Reverse"}

	err = t.Execute(f, cfg)
	if err != nil {
		panic(err)
	}
}
