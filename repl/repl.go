package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/slymanmrcan/hajilang/evaluator"
	"github.com/slymanmrcan/hajilang/lexer"
	"github.com/slymanmrcan/hajilang/object"
	"github.com/slymanmrcan/hajilang/parser"
)

const PROMPT = "haji> "

// Start başlatır REPL oturumunu
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Fprintf(out, "HajiLang REPL v0.1\n")
	fmt.Fprintf(out, "Çıkmak için 'exit' veya Ctrl+C\n\n")

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Exit komutu
		if line == "exit" || line == "çık" {
			fmt.Fprintf(out, "Görüşürüz haji!\n")
			return
		}

		// Boş satır skip
		if line == "" {
			continue
		}

		// Lexer -> Parser -> Evaluator pipeline
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		// Parser hataları varsa göster
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// Evaluate et
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprintf(out, "%s\n", evaluated.Inspect())
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "Hata haji!\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "  %s\n", msg)
	}
}
