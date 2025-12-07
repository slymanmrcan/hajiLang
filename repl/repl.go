package repl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/peterh/liner"
	"github.com/slymanmrcan/hajilang/evaluator"
	"github.com/slymanmrcan/hajilang/lexer"
	"github.com/slymanmrcan/hajilang/object"
	"github.com/slymanmrcan/hajilang/parser"
)

const PROMPT = "haji> "
const HISTORY_FILE = ".hajilang_history"

// Start başlatır REPL oturumunu
func Start(in io.Reader, out io.Writer) {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	// Geçmişi yükle
	if f, err := os.Open(getHistoryPath()); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	env := object.NewEnvironment()

	fmt.Fprintf(out, "HajiLang REPL v0.2 (Liner destekli)\n")
	fmt.Fprintf(out, "Çıkmak için 'exit' veya Ctrl+C\n\n")

	for {
		input, err := line.Prompt(PROMPT)
		if err == liner.ErrPromptAborted {
			fmt.Fprintf(out, "İptal edildi.\n")
			break
		} else if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(out, "Hata: %s\n", err)
			break
		}

		// Boş satır skip
		if input == "" {
			continue
		}

		// Exit komutu
		if input == "exit" || input == "çık" {
			break
		}

		// Komutu geçmişe ekle
		line.AppendHistory(input)

		// Lexer -> Parser -> Evaluator pipeline
		l := lexer.New(input)
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

	// Geçmişi kaydet
	if f, err := os.Create(getHistoryPath()); err == nil {
		line.WriteHistory(f)
		f.Close()
	}
}

func getHistoryPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return HISTORY_FILE
	}
	return filepath.Join(home, HISTORY_FILE)
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "Hata haji!\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "  %s\n", msg)
	}
}
