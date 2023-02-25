package xtraflag

import (
	"fmt"
	"strings"
)

type parser func(string) (map[string][]string, error)

func StringToStringSliceParser(entriesSep, assignSep, valsSep string) parser {
	return func(val string) (map[string][]string, error) {
		out := make(map[string][]string, 0)
		for _, e := range strings.Split(val, entriesSep) {
			url, operations, ok := strings.Cut(e, assignSep)
			if !ok {
				return nil, fmt.Errorf("%s is not valid input for an entry of map[string][]string: missing `%s` separator", e, assignSep)
			}
			out[url] = strings.Split(operations, valsSep)
		}
		return out, nil
	}
}
