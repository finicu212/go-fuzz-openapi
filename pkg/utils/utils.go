package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AsTitle exports names. (first letter of the keyword as capital letter)
// 1. Allows access outside output package,
// 2. Avoids collision with builtin keywords (i.e. `type`)
func AsTitle(s string) string {
	return cases.Title(language.English, cases.Compact).String(s)
}
