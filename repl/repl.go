package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/karchx/jj/lexer"
	"github.com/karchx/jj/token"
)

const PROMPT = ">> "

// Start start repl for jj.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
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
