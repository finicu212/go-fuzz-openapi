package filemaker

import (
	"os"
	"path/filepath"
)

// TryGetElseCreateFile handles output file via a singleton file pointer.
//
//	If file pointer not nil, return it. Otherwise:
//	 -> If file exists, clean it up. Otherwise:
//	  -> create file and all directories leading up to it.
//
// Needs to be run once invoked, as is a closure: TryGetElseCreateFile("main_test.go")()
func TryGetElseCreateFile(name string) (f func() (*os.File, error)) {
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
