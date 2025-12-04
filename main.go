package main

import (
	"fmt"
	"os"

	"github.com/slymanmrcan/hajilang/evaluator"
	"github.com/slymanmrcan/hajilang/lexer"
	"github.com/slymanmrcan/hajilang/object" // <--- EKLENDİ
	"github.com/slymanmrcan/hajilang/parser"
	"github.com/slymanmrcan/hajilang/repl"
)

func main() {
	if len(os.Args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
		return
	}
	if len(os.Args) < 2 && os.Args[1] == "--version" {
		fmt.Println("Kullanım: ./hajilang dosya.haji")
		fmt.Println("  veya   hajilang (REPL için)")
		return
	}

	dosyaAdi := os.Args[1]
	icerik, err := os.ReadFile(dosyaAdi)
	if err != nil {
		fmt.Printf("Dosya okunamadı: %s\n", err)
		return
	}

	kod := string(icerik)
	fmt.Printf(">> '%s' çalıştırılıyor...\n", dosyaAdi)
	// 1. Lexer
	l := lexer.New(kod)
	// 2. Parser
	p := parser.New(l)
	program := p.ParseProgram()

	// Hata var mı?
	if len(p.Errors()) != 0 {
		fmt.Println("❌ KODUNDA HATALAR VAR:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t -> %s\n", msg)
		}
		return
	}

	// 3. EVALUATOR (ÇALIŞTIRICI)
	// Burada programın içindeki her satırı tek tek çalıştırıp sonucunu yazıyoruz.

	// A) Hafızayı oluştur (RAM'i tak)
	env := object.NewEnvironment()
	verbose := false // veya flag ile al

	// B) Kodları sırayla çalıştır
	for i, stmt := range program.Statements {
		sonuc := evaluator.Eval(stmt, env)

		if sonuc != nil {
			// Son satır veya verbose modda göster
			if verbose || i == len(program.Statements)-1 {
				fmt.Println(sonuc.Inspect())
			}
		}
	}
}
