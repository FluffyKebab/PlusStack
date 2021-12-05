package evaluator

import (
	"fmt"
	"plusLang/arrayStack"
	"plusLang/lexer"
	"plusLang/objects"
	"plusLang/tokens"
	"strconv"
)

type evaluator struct {
	stack             *arrayStack.ArrayStack
	selectedArguments []objects.Object
	selectedFunction  objects.Function
}

func New() *evaluator {
	return &evaluator{
		stack:             arrayStack.New(),
		selectedArguments: make([]objects.Object, 0),
		selectedFunction:  objects.NewFunction(tokens.NONE),
	}
}

func (e *evaluator) Eval(input string) error {
	l := lexer.New(input)
	e.selectedArguments = make([]objects.Object, 0)
	e.selectedFunction = objects.NewFunction(tokens.NONE)

	for {
		curToken := l.NextToken()

		if curToken.Type == tokens.EOF {
			break
		}

		if curToken.Type == tokens.ILLEGAL {
			return fmt.Errorf("Token %s is illegal", curToken.Literal)
		}

		if curToken.Role == tokens.FUNCTION {
			e.selectedFunction = objects.NewFunction(curToken.Literal)
			err := e.evalCurFunction()
			if err != nil {
				return err
			}

			e.selectedArguments = make([]objects.Object, 0)
			continue
		}

		if curToken.Literal == tokens.POP && e.stack.GetObjectTypeAtPos(0) == tokens.FUNCTION {
			poppedItem, err := e.stack.Pop()
			if err != nil {
				return err
			}

			functionDef := poppedItem.(objects.Function).FunctionDef
			e.stack.Push(execCustomFunction(functionDef, e.selectedArguments)...)
			continue
		}

		if curToken.Role == tokens.ARGUMENT {
			if curToken.Type == tokens.INT {
				curArgument, err := strconv.Atoi(curToken.Literal)
				if err != nil {
					return err
				}

				e.selectedArguments = append(e.selectedArguments, objects.NewNumberArray([]float64{float64(curArgument)}))
			}

			if curToken.Type == tokens.FLOAT {
				curArgument, err := strconv.ParseFloat(curToken.Literal, 64)
				if err != nil {
					return err
				}

				e.selectedArguments = append(e.selectedArguments, objects.NewNumberArray([]float64{curArgument}))
			}

			if curToken.Type == tokens.STRING {
				e.selectedArguments = append(e.selectedArguments, objects.NewString(curToken.Literal))
			}

			if curToken.Type == tokens.POP {
				poppedItem, err := e.stack.Pop()
				if err != nil {
					return err
				}

				e.selectedArguments = append(e.selectedArguments, poppedItem)
			}

			if curToken.Type == tokens.POP_ALL {
				popedItems := e.stack.Items
				e.stack.Items = make([]objects.Object, 0)
				e.selectedArguments = append(e.selectedArguments, popedItems...)
			}

			if curToken.Type == tokens.FUNCTION {
				e.selectedArguments = append(e.selectedArguments, objects.NewFunction(curToken.Literal))
			}
		}
	}

	return nil
}

func (e *evaluator) evalCurFunction() error {
	functionLiteral := e.selectedFunction.FunctionDef

	if e.selectedFunction.FunctionDef == tokens.POP {
		poppedFunction, err := e.stack.Pop()
		if err != nil {
			return err
		}

		functionLiteral = poppedFunction.ToString()
	}

	returnValue, err := execFunction(functionLiteral, e.selectedArguments)
	if err != nil {
		return err
	}

	if returnValue != nil {
		e.stack.Push(returnValue...)
	}

	return nil
}

func (e *evaluator) Inspect() string {
	return e.stack.ToString()
}
