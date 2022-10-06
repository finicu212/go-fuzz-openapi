package main

import (
	"golang.org/x/example/stringutil"
	"testing"
	"unicode/utf8"
)

func Reverse(s string) string {
	return stringutil.Reverse(s)
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		r := Reverse(s)
		rr := Reverse(r)
		if s != rr {
			t.Errorf("Before: %q, after: %q", s, rr)
		}
		if utf8.ValidString(s) && !utf8.ValidString(r) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", r)
		}
	})
}

//func TableTest(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		{"1"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			fmt.Println("Hello")
//		})
//	}
//}
