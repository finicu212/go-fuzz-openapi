package utils

import (
	"os"
	"path/filepath"
)

// GetTestFileInstance handles output file via a singleton file pointer.
//   If file pointer not nil, return it. Otherwise:
//    -> If file exists, clean it up. Otherwise:
//     -> create file and all directories leading up to it.
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

// Map modifies each pair of a map[k]v using the provided function, and returns the modified slice
//
// Reference: https://gist.github.com/finicu212/8b95436426b3336981a9d82d0cab2d94
func Map[M interface{ ~map[Key]Value }, Key comparable, Value any, ReturnType any](s M, f func(Key, Value) ReturnType) []ReturnType {
	sm := make([]ReturnType, len(s))
	for i, v := range s {
		sm = append(sm, f(i, v))
	}
	return sm
}
