package main

import (
	"fmt"
	"os"

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
		runtime.RunServer(filename)
		return
	}

	// Help/Version
	if command == "--help" || command == "-h" || command == "--version" {
		printHelp()
		return
	}

	// Execute file
	executeFile(command)
}

func executeFile(filename string) {
	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: Cannot read file '%s'\nReason: %s\n", filename, err)
		return
	}

	code := string(source)
	fmt.Printf(">> Executing '%s'...\n", filename)

	// Setup environment
	env := object.NewEnvironment()
	runtime.RegisterJSON(env)
	runtime.RegisterUtils(env)

	// Lexer + Parser
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	fmt.Printf(">> Parsed %d statement(s)\n", len(program.Statements))
	// BU SATIRLARI EKLEYİN:
	fmt.Println("=== PARSED AST ===")
	for i, stmt := range program.Statements {
		fmt.Printf("%d: %s\n", i, stmt.String())
	}
	fmt.Println("==================")
	// Check parser errors
	if len(p.Errors()) != 0 {
		fmt.Println("❌ PARSE ERRORS:")
		for _, msg := range p.Errors() {
			fmt.Printf("   -> %s\n", msg)
		}
		return
	}

	// Evaluate
	result := evaluator.Eval(program, env)

	// Only show errors (puts already prints itself)
	if result != nil {
		if result.Type() == object.ERROR_OBJ {
			fmt.Println("❌ RUNTIME ERROR:", result.Inspect())
		} else if result.Type() != object.NULL_OBJ {
			// Program sonucu NULL değilse göster (REPL gibi)
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
