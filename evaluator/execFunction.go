package evaluator

import (
	"errors"
	"plusLang/arrayStack"
	"plusLang/objects"
	"plusLang/tokens"
	"regexp"
)

func getFunctionDefintion(function, arguments string) (func([]objects.Object) ([]objects.Object, error), error) {
	functionInString := function + arguments

	switch true {
	//Matches for "NONE"
	case regexp.MustCompile(tokens.NONE).MatchString(functionInString):
		return nilFunction, nil

	//Matches for "+"
	case regexp.MustCompile("\\" + tokens.PLUSS + "\\z").MatchString(functionInString):
		return pluss, nil

	//Matches for "+" with one "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.PLUSS + objects.INT_ARRAY + "\\z").MatchString(functionInString):
		return plussOneIntArray, nil

	//Matches for "+" with one or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.PLUSS + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return arithmeticMulitpleIntArrays(func(i1, i2 int) int { return i1 + i2 }), nil

	//Matches for "-" with one or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.MINUS + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return arithmeticMulitpleIntArrays(func(i1, i2 int) int { return i1 - i2 }), nil

	//Matches for [ with zero or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.APPEND + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return appendIntArray, nil

	//Matches for [ with zero or more "FUNCTION"
	case regexp.MustCompile("\\" + tokens.APPEND + "(" + objects.FUNCTION + ")+").MatchString(functionInString):
		return appendFunction, nil

	case regexp.MustCompile("\\" + tokens.REDUCE + "(" + objects.INT_ARRAY + ")+" + objects.FUNCTION + "\\z").MatchString(functionInString):
		return reduceFunction, nil

	case regexp.MustCompile("\\A" + tokens.SWAP + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return swapFunction, nil

	case regexp.MustCompile("\\A" + tokens.COMPOSE + "(" + objects.FUNCTION + ")+").MatchString(functionInString):
		return compose, nil

	case regexp.MustCompile("\\A" + tokens.DELETE + "(\\w)*").MatchString(functionInString):
		return delete, nil

	case regexp.MustCompile("\\A" + tokens.TERNARY + objects.INT_ARRAY + "(" + objects.FUNCTION + "){2}").MatchString(functionInString):
		return ternary, nil

	case regexp.MustCompile("\\" + tokens.MAP + "(" + objects.INT_ARRAY + ")+" + objects.FUNCTION).MatchString(functionInString):
		return mapFunction, nil
	}

	return nilFunction, errors.New("Function " + function + " is not defiend for arguments given: " + arguments)
}

func nilFunction(arguments []objects.Object) ([]objects.Object, error) {
	return nil, nil
}

func execFunction(functionName string, arguments []objects.Object) ([]objects.Object, error) {
	argumentTypes := ""
	for i := 0; i < len(arguments); i++ {
		argumentTypes += arguments[i].Type()
	}

	function, err := getFunctionDefintion(functionName, argumentTypes)
	if err != nil {
		return arguments, err
	}

	return function(arguments)
}

func execCustomFunction(functionDef string, arguments []objects.Object) []objects.Object {
	//TODO: cheack if function has points before first function

	//functionDef = "A" + functionDef
	eval := New()
	eval.stack = arrayStack.New(arguments...)
	eval.Eval(functionDef)
	return eval.stack.Items
}

/*
What pluss (+) function does:
+				Push new array with 1 element.
+Array			For each of the elements in the array push new array with the same lenght as the element.
+Array...		Add two arrays and push result.

What minus (-) function does:
-Array			Take every element in the array and return 0 - that element
-Array...

What append ([) function does:
<Array... 		Take all arrays given and combines them before pushing the result

*/
