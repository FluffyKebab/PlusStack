package arrayStack

import (
	"errors"
	"fmt"
	"plusLang/objects"
)

type ArrayStack struct {
	Items []objects.Object
}

func (stack *ArrayStack) Push(newItems ...objects.Object) {
	for _, item := range newItems {
		stack.Items = append(stack.Items, item)
	}
}

func (stack *ArrayStack) Pop() (objects.Object, error) {
	if len(stack.Items) == 0 {
		return objects.NewIntArray(make([]int, 0)), errors.New("Unable to pop stack with 0 items")
	}

	item := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	return item, nil
}

func (stack *ArrayStack) GetObjectTypeAtPos(pos int) string {
	if len(stack.Items)-pos-1 < 0 {
		return ""
	}

	item := stack.Items[len(stack.Items)-1-pos]
	return item.Type()
}

func (stack *ArrayStack) ToString() string {
	if len(stack.Items) <= 0 {
		return "No arrays on the stack"
	}

	output := ""

	for i := len(stack.Items) - 1; i >= 0; i-- {
		_, objectIsArray := stack.Items[i].(objects.IntArray)
		_, objectIsFunction := stack.Items[i].(objects.Function)
		curStackPos := fmt.Sprint(len(stack.Items)-i-1) + ") "

		if objectIsArray {
			output += fmt.Sprint(curStackPos + "[ " + stack.Items[i].ToString() + "]")
		} else if objectIsFunction {
			output += fmt.Sprint(curStackPos + "Function " + stack.Items[i].ToString())
		}

		if i != 0 {
			output += "\n"
		}
	}

	return output
}

func New(newItems ...objects.Object) *ArrayStack {
	stack := ArrayStack{
		Items: make([]objects.Object, 0),
	}

	for _, item := range newItems {
		stack.Push(item)
	}

	return &stack
}
