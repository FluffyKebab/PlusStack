package objects

import (
	"strconv"
)

type Object interface {
	Type() string
	ToString() string
}

type NumberArray struct {
	Elements []float64
}

func (a NumberArray) Type() string { return INT_ARRAY }
func (a NumberArray) ToString() string {
	output := ""
	for i := 0; i < len(a.Elements); i++ {
		output += strconv.FormatFloat(a.Elements[i], 'f', -1, 64) + " "
	}
	return output
}

func NewNumberArray(elements []float64) NumberArray {
	return NumberArray{
		Elements: elements,
	}
}

type String struct {
	Elements []rune
}

func (s String) Type() string     { return STRING }
func (s String) ToString() string { return string(s.Elements) }

func NewString(s string) String {
	return String{
		Elements: []rune(s),
	}
}

type Function struct {
	FunctionDef string
}

func (a Function) Type() string     { return FUNCTION }
func (a Function) ToString() string { return a.FunctionDef }

func NewFunction(functionDef string) Function {
	return Function{
		FunctionDef: functionDef,
	}
}
