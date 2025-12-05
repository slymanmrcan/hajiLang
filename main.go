package main

import (
	"fmt"
	"os"

	"github.com/slymanmrcan/hajilang/evaluator"
	"github.com/slymanmrcan/hajilang/lexer"
	"github.com/slymanmrcan/hajilang/object"
	"github.com/slymanmrcan/hajilang/parser"
	"github.com/slymanmrcan/hajilang/repl"
	"github.com/slymanmrcan/hajilang/runtime" // <--- MUTLAKA EKLE
)

func main() {
	// 1. HİÇ ARGÜMAN YOKSA -> REPL MODU
	// Kullanım: go run main.go
	if len(os.Args) < 2 {
		fmt.Println("HajiLang REPL (Çıkış için CTRL+C)")
		repl.Start(os.Stdin, os.Stdout)
		return
	}
	// İlk argümanı al (server mı? dosya adı mı?)
	komut := os.Args[1]
	// 2. SERVER MODU
	// Kullanım: go run main.go server
	// 2. SERVER MODU
	if komut == "server" {
		dosya := "server.haji" // Varsayılan dosya adı
		// Eğer kullanıcı "go run main.go server benimapi.haji" dediyse:
		if len(os.Args) > 2 {
			dosya = os.Args[2]
		}
		runtime.RunServer(dosya)
		return
	}
	// 3. VERSİYON/YARDIM
	if komut == "--help" || komut == "--version" {
		yazdirYardim()
		return
	}
	// 4. DOSYA ÇALIŞTIRMA MODU
	// Kullanım: go run main.go dosya.haji
	calistirDosya(komut)
}

func calistirDosya(dosyaAdi string) {
	icerik, err := os.ReadFile(dosyaAdi)
	if err != nil {
		fmt.Printf("Hata: '%s' dosyası okunamadı.\nSebep: %s\n", dosyaAdi, err)
		return
	}

	kod := string(icerik)
	fmt.Printf(">> '%s' çalıştırılıyor...\n", dosyaAdi)

	// Environment (Ortam) Oluştur
	env := object.NewEnvironment()

	// Modülleri yükle
	runtime.RegisterJSON(env)
	runtime.RegisterUtils(env)

	l := lexer.New(kod)
	p := parser.New(l)
	program := p.ParseProgram()

	// 1. PARSER HATALARI (Yazım hataları)
	if len(p.Errors()) != 0 {
		fmt.Println("❌ YAZIM HATALARI VAR (PARSER):")
		for _, msg := range p.Errors() {
			fmt.Printf("\t-> %s\n", msg)
		}
		return
	}

	// 2. ÇALIŞTIRMA (Eval)
	// Eval fonksiyonu bir nesne döner. Bu nesne HATA olabilir!
	sonuc := evaluator.Eval(program, env)

	// 3. RUNTIME HATALARI (Çalışma zamanı hataları)
	if sonuc != nil {
		// Eğer dönen şey bir HATA nesnesi ise ekrana kırmızı bas
		if sonuc.Type() == object.ERROR_OBJ {
			fmt.Println("❌ ÇALIŞMA HATASI: " + sonuc.Inspect())
		}
		// İsteğe bağlı: Scriptin son sonucunu görmek istersen bunu aç:
		// else {
		// 	fmt.Println(sonuc.Inspect())
		// }
	}
}

func yazdirYardim() {
	fmt.Println(`
HajiLang Kullanım Kılavuzu:
---------------------------
1. REPL (Konsol) Modu:
   ./hajilang

2. Dosya Çalıştırma:
   ./hajilang dosya.haji

3. Web Sunucusu Başlatma:
   ./hajilang server
	`)
}
