package evaluator

import (
	"errors"
	"plusLang/objects"
)

func reduceFunction(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) < 3 {
		return arguments, errors.New("The reduce function (/) have to have at least 3 arguments")
	}

	function, ok := arguments[len(arguments)-1].(objects.Function)
	if !ok {
		return arguments, errors.New("The reduce function's (/) last argument have to be a function")
	}

	startingPointArray, ok := arguments[len(arguments)-2].(objects.IntArray)

	if len(startingPointArray.Elements) != 1 {
		return arguments, errors.New("The reduce function's (/) next to last argument have to be a intArray of lenght 1")
	}
	startingPoint := startingPointArray.Elements[0].Value

	outputObjects := make([]objects.Object, 0)

	for i := 0; i < len(arguments)-2; i++ {
		curIntArray, ok := arguments[i].(objects.IntArray)
		if !ok {
			return arguments, errors.New("IIE: reduceFunction given wrong arguments")
		}

		curOutput := startingPoint
		for j := 0; j < len(curIntArray.Elements); j++ {
			curFunctionOutput, err := execFunction(function.FunctionDef, []objects.Object{objects.NewIntArray([]int{curOutput}), objects.NewIntArray([]int{curIntArray.Elements[j].Value})})
			if err != nil {
				return arguments, err
			}

			if len(curFunctionOutput) != 1 {
				return arguments, errors.New("Function given to reduce (/) have to return only one object")
			}

			outputInIntArray, ok := curFunctionOutput[0].(objects.IntArray)
			if !ok {
				return arguments, errors.New("Function given to reduce (/) have to return a intArray with one element")
			}

			if len(outputInIntArray.Elements) != 1 {
				return arguments, errors.New("Function given to reduce (/) have to return a intArray with one element")
			}

			curOutput = outputInIntArray.Elements[0].Value
		}

		outputObjects = append(outputObjects, objects.NewIntArray([]int{curOutput}))
	}

	return outputObjects, nil
}
