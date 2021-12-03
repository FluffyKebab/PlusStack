package repl

import (
	"bufio"
	"fmt"
	"io"
	"plusLang/evaluator"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	evaluator := evaluator.New()

	for {
		fmt.Printf("PROMPT> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		err := evaluator.Eval(scanner.Text())
		if err != nil {
			fmt.Println("Folowing error accured when running code given: ", err.Error())
		}

		io.WriteString(out, evaluator.Inspect())
		io.WriteString(out, "\n")
	}
}
