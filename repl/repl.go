package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/KSSidhu/go-interpreter/lexer"
	"github.com/KSSidhu/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Read from input until hitting a new line
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		// Pass input to lexer to parse
		line := scanner.Text()
		l := lexer.New(line)

		// Iterate over tokens from input and print to console
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
