package go_fuzz_openapi

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Hello")
		})
	}
}
