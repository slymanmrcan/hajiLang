package main

import (
	"fmt"
	"hajilang/evaluator" // <--- EKLENDİ
	"hajilang/lexer"
	"hajilang/object" // <--- EKLENDİ
	"hajilang/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: ./hajilang dosya.haji")
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
	fmt.Println("--------------------------------------------")

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

	// B) Kodları sırayla çalıştır
	for _, stmt := range program.Statements {
		// Hafızayı (env) da gönderiyoruz artık!
		sonuc := evaluator.Eval(stmt, env)

		if sonuc != nil {
			fmt.Println("SONUÇ >", sonuc.Inspect())
		}
	}

	fmt.Println("--------------------------------------------")
}
