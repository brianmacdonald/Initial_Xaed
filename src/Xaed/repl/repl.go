package repl

import (
	"bufio"
	"fmt"
	"io"
	"Xaed/lexer"
	"Xaed/parser"
	"Xaed/evaluator"
)

const PROMPT = "[%d] >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var i = 0
	for {
		fmt.Printf(PROMPT, i)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
		i++
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

