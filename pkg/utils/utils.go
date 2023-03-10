package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"
)

// AsTitle exports names. (first letter of the keyword as capital letter)
// 1. Allows access outside output package,
// 2. Avoids collision with builtin keywords (i.e. `type`)
func AsTitle(s string) string {
	return cases.Title(language.English, cases.Compact).String(s)
}

// RefPathToType converts `#/components/schemas/Category` to `Category`
func RefPathToType(ref string) string {
	return filepath.Base(ref)
}

// RefPathToEndpoint converts `/store/order/{pet_id}` to `/store/order`
func RefPathToEndpoint(ref string) string {
	paths := strings.Split(ref, "{")
	return paths[0]
}

// Map modifies each pair of a map[k]v using the provided function, and returns the modified slice
// (note to self: update gist if any changes to this func https://gist.github.com/finicu212/8b95436426b3336981a9d82d0cab2d94)
func Map[M interface{ ~map[Key]Value }, Key comparable, Value any, ReturnType any](s M, f func(Key, Value) ReturnType) []ReturnType {
	sm := make([]ReturnType, len(s))
	for i, v := range s {
		sm = append(sm, f(i, v))
	}
	return sm
}
