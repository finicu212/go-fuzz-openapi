package endpoints

import (
	"path/filepath"
	"strings"
)

// RefPathToType converts `#/components/schemas/Category` to `Category`
func RefPathToType(ref string) string {
	return filepath.Base(ref)
}

// RefPathToEndpoint converts `/store/order/{pet_id}` to `/store/order`
func RefPathToEndpoint(ref string) string {
	// TODO: This seems wrong...
	paths := strings.Split(ref, "{")
	return paths[0]
}
