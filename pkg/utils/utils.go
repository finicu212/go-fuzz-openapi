package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
)

// GetTestFileInstance handles output file via a singleton file pointer.
//
//	If file pointer not nil, return it. Otherwise:
//	 -> If file exists, clean it up. Otherwise:
//	  -> create file and all directories leading up to it.
//
// Needs to be run once invoked, as is a closure: GetTestFileInstance("main_test.go")()
func GetTestFileInstance(name string) (f func() (*os.File, error)) {
	var file *os.File
	f = func() (*os.File, error) {
		if file != nil {
			// Singleton already previously initialized. Return it
			return file, nil
		}
		_, err := os.Stat(name)
		if err == nil {
			// file exists, but singleton not initialized. DIRTY FILE!
			file, err = os.OpenFile(name, os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				return nil, err
			}
			err = file.Truncate(0) // Cleanup file contents.
			if err != nil {
				return nil, err
			}
			return file, nil
		}
		// File had never been created before
		err = os.MkdirAll(filepath.Dir(name), 0750)
		if err != nil {
			return nil, err
		}
		file, err = os.Create(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return f
}

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
