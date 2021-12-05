package evaluator

import (
	"errors"
	"fmt"
	"plusLang/objects"
	"sort"
)

func pluss(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 0 {
		return arguments, errors.New("Internal interpreter error: Wrong number of arguments passed to pluss function")
	}

	output := make([]int, 1)
	output[0] = 0
	return []objects.Object{objects.NewIntArray(output)}, nil
}

func plussOneIntArray(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 1 {
		return arguments, errors.New("Internal interpreter error: Wrong number of arguments passed to plussOneIntArray function")
	}

	inputArray, ok := arguments[0].(objects.IntArray)
	if !ok {
		return arguments, errors.New("Internal interpreter error: Argument given to plussOneIntArray not of type intArray")
	}

	ouputObjects := make([]objects.Object, 0)

	for i := 0; i < len(inputArray.Elements); i++ {
		curArray := make([]int, inputArray.Elements[i].Value)
		for i := 0; i < len(curArray); i++ {
			curArray[i] = 0
		}

		ouputObjects = append(ouputObjects, objects.NewIntArray(curArray))
	}
	return ouputObjects, nil
}

func negative(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 1 {
		return arguments, errors.New("Negative needs one argument")
	}

	integerArray, ok := arguments[0].(objects.IntArray)
	if !ok {
		return arguments, errors.New("Negative argument have to be of type int array")
	}

	output := make([]int, 0)
	for i := 0; i < len(integerArray.Elements); i++ {
		output = append(output, -integerArray.Elements[i].Value)
	}

	return []objects.Object{objects.NewIntArray(output)}, nil
}

func arithmeticMulitpleIntArrays(arithmeticFunc func(int, int) int) func(arguments []objects.Object) ([]objects.Object, error) {
	return func(arguments []objects.Object) ([]objects.Object, error) {
		inputArrays := make([]objects.IntArray, len(arguments))
		for i := 0; i < len(inputArrays); i++ {
			curIntArray, ok := arguments[i].(objects.IntArray)
			inputArrays[i] = curIntArray
			if !ok {
				return arguments, errors.New("Internal interpreter error: Argument given to arithmeticMulitpleIntArrays at pos " + fmt.Sprint(i) + " not of type intArray")
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
			outputArray := make([]int, len(inputArrays[0].Elements))

			for i := 0; i < len(inputArrays); i++ {
				if i == 0 {
					for j := 0; j < len(inputArrays[0].Elements); j++ {
						outputArray[j] = inputArrays[i].Elements[j].Value
					}
					continue
				}

				// 200 10 10

				for j := 0; j < len(inputArrays[0].Elements); j++ {
					outputArray[j] = arithmeticFunc(outputArray[j], inputArrays[i].Elements[j].Value)
				}
			}

			return []objects.Object{objects.NewIntArray(outputArray)}, nil
		}

		sort.SliceStable(inputArrays, func(i, j int) bool {
			return len(inputArrays[i].Elements) < len(inputArrays[j].Elements)
		})

		if len(inputArrays[0].Elements) != 1 {
			return arguments, errors.New("Arguments given to function + does not match")
		}

		outputArrays := make([]objects.Object, 0)

		plussConstant := 0
		for i := 0; i < len(inputArrays); i++ {
			if len(inputArrays[i].Elements) == 1 {
				plussConstant += inputArrays[i].Elements[0].Value
				continue
			}

			curOutputArray := make([]int, 0)
			for j := 0; j < len(inputArrays[i].Elements); j++ {
				curOutputArray = append(curOutputArray, arithmeticFunc(inputArrays[i].Elements[j].Value, plussConstant))
			}

			outputArrays = append(outputArrays, objects.NewIntArray(curOutputArray))
		}

		return outputArrays, nil
	}
}
