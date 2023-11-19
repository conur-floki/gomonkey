package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/conur-floki/gomonkey/lexer"
	"github.com/conur-floki/gomonkey/token"
)

const PROMPT = "-->"

// Just simple Start function to start the repl instance
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
