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

	case regexp.MustCompile("\\A" + tokens.TO_NUM).MatchString(functionInString):
		return toNum, nil

	case regexp.MustCompile("\\A" + tokens.TO_STRING).MatchString(functionInString):
		return toString, nil

	//Matches for "+"
	case regexp.MustCompile("\\" + tokens.PLUS + "\\z").MatchString(functionInString):
		return plus, nil

	//Matches for "+" with one "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.PLUS + objects.INT_ARRAY + "\\z").MatchString(functionInString):
		return plussOneIntArray, nil

	//Matches for "+" with one or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.PLUS + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return arithmeticMultipleIntArrays(func(i1, i2 float64) float64 { return i1 + i2 }), nil

	//Matches for "-" with one "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.MINUS + objects.INT_ARRAY + "\\z").MatchString(functionInString):
		return negative, nil

	//Matches for "-" with one or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.MINUS + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return arithmeticMultipleIntArrays(func(i1, i2 float64) float64 { return i1 - i2 }), nil

	//Matches for "*" with one or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.MULT + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return arithmeticMultipleIntArrays(func(i1, i2 float64) float64 { return i1 * i2 }), nil

	//Matches for ":" with one or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.DIV + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return arithmeticMultipleIntArrays(func(i1, i2 float64) float64 { return i1 / i2 }), nil

	case regexp.MustCompile("\\A" + tokens.EQUAL + ".+").MatchString(functionInString):
		return equal, nil

	case regexp.MustCompile("\\" + tokens.NOT + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return not, nil

	//Matches for ] with zero or more "INT_ARRAY"
	case regexp.MustCompile("\\" + tokens.APPEND + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return appendIntArray, nil

	//Matches for ] with zero or more "STRING"
	case regexp.MustCompile("\\" + tokens.APPEND + "(" + objects.STRING + ")+").MatchString(functionInString):
		return appendString, nil

	//Matches for ] with zero or more "FUNCTION"
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

	case regexp.MustCompile("\\" + tokens.LESS_THEN + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return comperasion(func(a, b float64) bool { return a < b }), nil

	case regexp.MustCompile("\\" + tokens.GREATER_THEN + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return comperasion(func(a, b float64) bool { return a > b }), nil

	case regexp.MustCompile("\\" + tokens.EQUAL_OR_LESS_THEN + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return comperasion(func(a, b float64) bool { return a <= b }), nil

	case regexp.MustCompile("\\" + tokens.EQUAL_OR_GREATER_THEN + "(" + objects.INT_ARRAY + ")+").MatchString(functionInString):
		return comperasion(func(a, b float64) bool { return a >= b }), nil
	}

	return nilFunction, errors.New("Function " + function + " is not defined for arguments given: " + arguments)
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
	eval := New()
	eval.stack = arrayStack.New(arguments...)
	eval.Eval(functionDef)
	return eval.stack.Items
}

/*
What plus (+) function does:
+				Push new array with 1 element.
+Array			For each of the elements in the array push new array with the same length as the element.
+Array...		Add two arrays and push result.

What minus (-) function does:
-Array			Take every element in the array and return 0 - that element
-Array...

What append ([) function does:
<Array... 		Take all arrays given and combines them before pushing the result

*/
