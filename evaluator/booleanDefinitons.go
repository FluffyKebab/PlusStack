package evaluator

import (
	"errors"
	"plusLang/objects"
	"sort"
)

func not(arguments []objects.Object) ([]objects.Object, error) {
	output := make([]objects.Object, 0)

	for i := 0; i < len(arguments); i++ {
		curNumberArray, ok := arguments[i].(objects.NumberArray)
		if !ok {
			return arguments, errors.New("The no function argument have to be of type number array")
		}

		for j := 0; j < len(curNumberArray.Elements); j++ {
			if curNumberArray.Elements[j] == 1 {
				curNumberArray.Elements[j] = 0
			} else if curNumberArray.Elements[j] == 0 {
				curNumberArray.Elements[j] = 1
			} else {
				return arguments, errors.New("In not all elements in argument arrays have to be [ 1 ] or [ 0 ]")
			}
		}

		output = append(output, curNumberArray)
	}

	return output, nil
}

func equal(arguments []objects.Object) ([]objects.Object, error) {
	if len(arguments) != 2 {
		return arguments, errors.New("The equal function (=) needs two objects as input")
	}

	isEqual := arguments[0].ToString() == arguments[1].ToString()

	if isEqual {
		return []objects.Object{objects.NewNumberArray([]float64{1})}, nil
	}
	return []objects.Object{objects.NewNumberArray([]float64{0})}, nil
}

func notEqual(arguments []objects.Object) ([]objects.Object, error) {
	isEqual, err := equal(arguments)
	if err != nil {
		return arguments, errors.New("Error when executing equal in not equal: " + err.Error())
	}

	return not(isEqual)
}

func comperasion(f func(a, b float64) bool) func(arguments []objects.Object) ([]objects.Object, error) {
	return func(arguments []objects.Object) ([]objects.Object, error) {
		if len(arguments) != 2 {
			return arguments, errors.New("All comperasion functions needs two arguments")
		}

		array1, ok1 := arguments[0].(objects.NumberArray)
		array2, ok2 := arguments[1].(objects.NumberArray)
		if !ok1 || !ok2 {
			return arguments, errors.New("All comperasion function arguments have to be number arrays")
		}

		if len(array1.Elements) == len(array2.Elements) {
			output := make([]float64, 0)
			for i := 0; i < len(array1.Elements); i++ {
				compered := f(array1.Elements[i], array2.Elements[i])
				if compered {
					output = append(output, 1)
				} else {
					output = append(output, 0)
				}
			}
			return []objects.Object{objects.NewNumberArray(output)}, nil
		}

		arrays := []objects.NumberArray{array1, array2}

		sort.SliceStable(arrays, func(i, j int) bool {
			return len(arrays[i].Elements) < len(arrays[j].Elements)
		})

		if len(arrays[0].Elements) != 1 {
			return arguments, errors.New("Arguments given to comperasion function does not match")
		}

		output := make([]float64, 0)
		for i := 0; i < len(arrays[1].Elements); i++ {
			compered := f(arrays[1].Elements[i], arrays[0].Elements[0])
			if compered {
				output = append(output, 1)
			} else {
				output = append(output, 0)
			}
		}
		return []objects.Object{objects.NewNumberArray(output)}, nil
	}
}
