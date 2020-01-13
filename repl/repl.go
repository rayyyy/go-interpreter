package repl

import (
	"io"
	"go-interpreter/lexer"
	"go-interpreter/token"
	"fmt"
	"bufio"
)

// PROMPT コンソールの左側
const PROMPT = ">> "

// Start 開始関数
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