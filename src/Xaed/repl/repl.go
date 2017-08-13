package repl

import (
	"Xaed/lexer"
	"Xaed/parser"
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const PROMPT = "[%d] >> "

func Start(in io.Reader, out io.Writer) {
	var input_count = 0
	scanner := bufio.NewScanner(in)
	for {
		input_count++
		fmt.Printf(PROMPT, input_count)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors(), input_count)
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
func printParserErrors(out io.Writer, errors []string, input_count int) {
	for _, msg := range errors {
		io.WriteString(out, "[" + strconv.Itoa(input_count) + "] \t"+msg+"\n")
	}
}
