package evaluator

import (
	"fmt"
	"plusLang/objects"
)

func appendIntArray(arguments []objects.Object) ([]objects.Object, error) {
	outputObject := objects.NewNumberArray(make([]float64, 0))

	for i := 0; i < len(arguments); i++ {
		curIntArray, ok := arguments[i].(objects.NumberArray)
		if !ok {
			return arguments, fmt.Errorf("Argument given to append not of type intArray")
		}

		outputObject.Elements = append(outputObject.Elements, curIntArray.Elements...)
	}

	return []objects.Object{outputObject}, nil
}

func appendFunction(arguments []objects.Object) ([]objects.Object, error) {
	return arguments, nil
}

func appendString(arguments []objects.Object) ([]objects.Object, error) {
	return arguments, nil
}
