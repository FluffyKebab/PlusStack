package evaluator

import (
	"errors"
	"fmt"
	"plusLang/objects"
	"sort"
)

func plus(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 0 {
		return arguments, errors.New("Internal interpreter error: Wrong number of arguments passed to plus function")
	}

	output := make([]float64, 1)
	output[0] = 0
	return []objects.Object{objects.NewNumberArray(output)}, nil
}

func plussOneIntArray(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 1 {
		return arguments, errors.New("Internal interpreter error: Wrong number of arguments passed to plussOneIntArray function")
	}

	inputArray, ok := arguments[0].(objects.NumberArray)
	if !ok {
		return arguments, errors.New("Internal interpreter error: Argument given to plussOneIntArray not of type intArray")
	}

	outputObjects := make([]objects.Object, 0)

	for i := 0; i < len(inputArray.Elements); i++ {
		curArray := make([]float64, int(inputArray.Elements[i]))
		for i := 0; i < len(curArray); i++ {
			curArray[i] = 0
		}

		outputObjects = append(outputObjects, objects.NewNumberArray(curArray))
	}
	return outputObjects, nil
}

func negative(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 1 {
		return arguments, errors.New("Negative needs one argument")
	}

	integerArray, ok := arguments[0].(objects.NumberArray)
	if !ok {
		return arguments, errors.New("Negative argument have to be of type int array")
	}

	output := make([]float64, 0)
	for i := 0; i < len(integerArray.Elements); i++ {
		output = append(output, -integerArray.Elements[i])
	}

	return []objects.Object{objects.NewNumberArray(output)}, nil
}

func arithmeticMultipleIntArrays(arithmeticFunc func(float64, float64) float64) func(arguments []objects.Object) ([]objects.Object, error) {
	return func(arguments []objects.Object) ([]objects.Object, error) {
		inputArrays := make([]objects.NumberArray, len(arguments))
		for i := 0; i < len(inputArrays); i++ {
			curIntArray, ok := arguments[i].(objects.NumberArray)
			inputArrays[i] = curIntArray
			if !ok {
				return arguments, errors.New("Internal interpreter error: Argument given to arithmeticMultipleIntArrays at pos " + fmt.Sprint(i) + " not of type intArray")
			}
		}

		lenEqual := true
		for i := 0; i < len(inputArrays); i++ {
			lenEqual = len(inputArrays[0].Elements) == len(inputArrays[i].Elements)
			if !lenEqual {
				break
			}
		}

		if lenEqual {
			outputArray := make([]float64, len(inputArrays[0].Elements))

			for i := 0; i < len(inputArrays); i++ {
				if i == 0 {
					for j := 0; j < len(inputArrays[0].Elements); j++ {
						outputArray[j] = inputArrays[i].Elements[j]
					}
					continue
				}

				// 200 10 10

				for j := 0; j < len(inputArrays[0].Elements); j++ {
					outputArray[j] = arithmeticFunc(outputArray[j], inputArrays[i].Elements[j])
				}
			}

			return []objects.Object{objects.NewNumberArray(outputArray)}, nil
		}

		sort.SliceStable(inputArrays, func(i, j int) bool {
			return len(inputArrays[i].Elements) < len(inputArrays[j].Elements)
		})

		if len(inputArrays[0].Elements) != 1 {
			return arguments, errors.New("Arguments given to function + does not match")
		}

		outputArrays := make([]objects.Object, 0)

		plusConstant := 0.0
		for i := 0; i < len(inputArrays); i++ {
			if len(inputArrays[i].Elements) == 1 {
				plusConstant += inputArrays[i].Elements[0]
				continue
			}

			curOutputArray := make([]float64, 0)
			for j := 0; j < len(inputArrays[i].Elements); j++ {
				curOutputArray = append(curOutputArray, arithmeticFunc(inputArrays[i].Elements[j], plusConstant))
			}

			outputArrays = append(outputArrays, objects.NewNumberArray(curOutputArray))
		}

		return outputArrays, nil
	}
}
