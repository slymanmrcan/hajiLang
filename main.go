package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/slymanmrcan/hajilang/evaluator"
	"github.com/slymanmrcan/hajilang/lexer"
	"github.com/slymanmrcan/hajilang/object"
	"github.com/slymanmrcan/hajilang/parser"
	"github.com/slymanmrcan/hajilang/repl"
	"github.com/slymanmrcan/hajilang/runtime"
)

func main() {
	// No arguments -> REPL mode
	if len(os.Args) < 2 {
		fmt.Println("HajiLang REPL v1.0 (Press CTRL+C to exit)")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	command := os.Args[1]

	// Server mode
	if command == "server" {
		filename := "server.haji"
		if len(os.Args) > 2 {
			filename = os.Args[2]
		}

		// ← BURASI ÇOK ÖNEMLİ! ScriptBaseDir'i ayarla
		absPath, err := filepath.Abs(filename)
		if err == nil {
			runtime.ScriptBaseDir = filepath.Dir(absPath)
			fmt.Printf("📁 Base Directory: %s\n", runtime.ScriptBaseDir)
		}

		runtime.RunServer(filename)
		return
	}

	// Help/Version
	if command == "--help" || command == "-h" || command == "--version" {
		printHelp()
		return
	}

	// Execute file
	executeFileWithLogging(command)
}

func executeFileWithLogging(filename string) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ Dosya okunamadı: %s\n", err.Error())
		return
	}

	// ScriptBaseDir'i execute file için de ayarla
	absPath, err := filepath.Abs(filename)
	if err == nil {
		runtime.ScriptBaseDir = filepath.Dir(absPath)
	}

	code := string(source)
	fmt.Printf(">> '%s' dosyası çalıştırılıyor...\n", filename)

	// Ortamı ayarla
	env := object.NewEnvironment()
	runtime.RegisterJSON(env)
	runtime.RegisterUtils(env)

	// Lexer + Parser
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	fmt.Printf(">> %d ifade(ler) ayrıştırıldı\n", len(program.Statements))

	// Parser hatalarını kontrol et
	if len(p.Errors()) != 0 {
		fmt.Println("❌ AYRIŞTIRMA HATALARI:")
		for _, msg := range p.Errors() {
			fmt.Printf("   -> %s\n", msg)
		}
		return
	}

	// Değerlendir
	result := evaluator.Eval(program, env)

	// Sadece hataları göster
	if result != nil {
		if result.Type() == object.ERROR_OBJ {
			fmt.Println("❌ ÇALIŞMA ZAMANI HATASI:", result.Inspect())
		} else if result.Type() != object.NULL_OBJ {
			fmt.Println(result.Inspect())
		}
	}
}

func printHelp() {
	fmt.Println(`
HajiLang Interpreter v1.0
-------------------------
Usage:
  hajilang                    Start REPL mode
  hajilang <file.haji>        Execute a file
  hajilang server [file]      Start HTTP server
  hajilang --help             Show this help

Examples:
  hajilang                    # Interactive mode
  hajilang app.haji           # Run app.haji
  hajilang server api.haji    # Start server with api.haji
	`)
}
