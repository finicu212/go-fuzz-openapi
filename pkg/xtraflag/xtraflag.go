package xtraflag

import (
	"fmt"
	"reflect"
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

type Value[T any] struct {
	parse func(val string) (T, error)
	value *T
}

func NewValue[T any](val T, p *T, parse func(val string) (T, error)) *Value[T] {
	v := new(Value[T])
	v.parse = parse
	v.value = p
	*v.value = val
	return v
}

func (v *Value[T]) Set(val string) error {
	var err error
	*v.value, err = v.parse(val)
	return err
}

func (v *Value[T]) Type() string {
	return reflect.TypeOf(v).Name()
}

func (v *Value[T]) String() string {
	if v.value == nil || reflect.ValueOf(*v.value).IsZero() {
		return ""
	}
	return fmt.Sprint(*v.value)
}
