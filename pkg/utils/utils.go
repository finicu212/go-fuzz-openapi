package utils

import "os"

// GetTestFileInstance initializes and returns a singleton file pointer.
// If file doesn't exist, then create it.
func GetTestFileInstance(name string) (f func() (*os.File, error)) {
	var file *os.File
	f = func() (*os.File, error) {
		if file != nil {
			// Singleton already previously initialized. Return it
			return file, nil
		}
		_, err := os.Stat(name)
		if err == nil {
			// file exists, but singleton not initialized. Possibly dirty file?  TODO: Cleanup before / throw error / currently: return as is?
			file, err = os.Open(name)
			if err != nil {
				return nil, err
			}
			return file, nil
		}
		// File had never been created before
		file, err = os.Create(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return f
}
