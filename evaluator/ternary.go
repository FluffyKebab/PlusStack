package evaluator

import (
	"errors"
	"plusLang/objects"
)

func ternary(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 3 {
		return arguments, errors.New("The ternary function need three arguments")
	}

	intArray, ok := arguments[0].(objects.IntArray)
	if !ok {
		return arguments, errors.New("Argument given to the ternary function (?) at pos 0 not of type intArray")
	}

	if len(intArray.Elements) != 1 {
		return arguments, errors.New("Argument given to the ternary function (?) at pos 0 does not have a lenght of 1")
	}

	boleanValue := intArray.Elements[0].Value
	if boleanValue != 0 && boleanValue != 1 {
		return arguments, errors.New("Argument given to the ternary function (?) at pos 0 not a bolean value (1 or 0)")
	}

	function1, ok := arguments[1].(objects.Function)
	if !ok {
		return arguments, errors.New("Argument given to the ternary function (?) at pos 1 not of type function")
	}

	function2, ok := arguments[2].(objects.Function)
	if !ok {
		return arguments, errors.New("Argument given to the ternary function (?) at pos 2 not of type function")
	}

	if boleanValue == 1 {
		return []objects.Object{function1}, nil
	}

	return []objects.Object{function2}, nil
}
