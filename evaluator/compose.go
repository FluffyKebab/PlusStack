package evaluator

import "plusLang/objects"

func compose(arguments []objects.Object) ([]objects.Object, error) {
	outputFunctionDef := ""

	for i := len(arguments) - 1; i >= 0; i-- {
		if i != len(arguments)-1 {
			outputFunctionDef += "A"
		}
		curFunction := arguments[i].(objects.Function)
		outputFunctionDef += curFunction.FunctionDef
	}

	return []objects.Object{objects.NewFunction(outputFunctionDef)}, nil
}
