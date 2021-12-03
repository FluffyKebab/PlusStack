package evaluator

import (
	"plusLang/objects"
	"reflect"
)

func swapFunction(arguments []objects.Object) ([]objects.Object, error) {
	reverseAny(arguments)
	return arguments, nil
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
