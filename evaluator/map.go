package evaluator

import (
	"errors"
	"plusLang/objects"
)

func mapFunction(arguments []objects.Object) ([]objects.Object, error) {
	function, ok := arguments[len(arguments)-1].(objects.Function)
	if !ok {
		return nil, errors.New("Map function have to get a function as its last element")
	}

	outputObjects := make([]objects.Object, 0)

	for i := 0; i < len(arguments)-1; i++ {
		curIntArray := arguments[i].(objects.NumberArray)
		for j := 0; j < len(curIntArray.Elements); j++ {
			curOutput := execCustomFunction(function.FunctionDef, []objects.Object{objects.NewNumberArray([]float64{float64(j)}), objects.NewNumberArray([]float64{curIntArray.Elements[j]})})
			if len(curOutput) != 1 {
				return arguments, errors.New("Function given to map (-<) have to return one intArray with one element")
			}
			curIntArray.Elements[j] = curOutput[0].(objects.NumberArray).Elements[0]
		}

		outputObjects = append(outputObjects, curIntArray)
	}

	return outputObjects, nil
}
