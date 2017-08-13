package repl

import (
	"bufio"
	"fmt"
	"io"
	"Xaed/lexer"
	"Xaed/parser"
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
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
		i++
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

