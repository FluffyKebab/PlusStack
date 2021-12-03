package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"plusLang/evaluator"
	"plusLang/repl"
)

const usage = `Usage:
	plusLang [command]

Available Commands:
	run	Runs the file given
	repl	Starts the repl
`

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] == "run" {
		f, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}

		e := evaluator.New()
		e.Eval(string(f))
		fmt.Println(e.Inspect())
		return
	}

	if os.Args[1] == "repl" {
		repl.Start(os.Stdin, os.Stdout)
	}
}

//
//<10 10 10 15
// +10.
// +20 -<.()
// @0
// Make array with ten elements and add to the stack.
// Pop from the stack and for each of the elements do the folowting function:
// Make an array with 10 elements

// +10 . + (  )
//f <-

// Pop an array and add 2 to each item
// +10 <. 10 8

// Pop an array and multiply to each item
// 0+ .4 0+ .2 .+.
// MAKE INT POP INT 	MAKE INT POP INT	 POP PLUSS POP

// Make the array 10 30 4 5
//<10 30 5 5

//+0 .1 .3 .+5
