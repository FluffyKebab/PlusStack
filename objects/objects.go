package objects

import (
	"fmt"
)

type Object interface {
	Type() string
	ToString() string
}

type Integer struct {
	Value int
}

func (i Integer) Type() string     { return INT }
func (i Integer) ToString() string { return fmt.Sprintf("%d", i.Value) }

func NewInt(i int) Object {
	return Integer{
		Value: i,
	}
}

type IntArray struct {
	Elements []Integer
}

func (a IntArray) Type() string { return INT_ARRAY }
func (a IntArray) ToString() string {
	output := ""
	for i := 0; i < len(a.Elements); i++ {
		output += a.Elements[i].ToString() + " "
	}
	return output
}

func NewIntArray(elements []int) IntArray {
	intArrayElem := make([]Integer, len(elements))
	for i := 0; i < len(intArrayElem); i++ {
		intArrayElem[i] = Integer{Value: elements[i]}
	}

	return IntArray{
		Elements: intArrayElem,
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
