package repl

import (
	"bufio"
	"fmt"
	"io"
	"Xaed/lexer"
	"Xaed/token"
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
		i++
	}
}
