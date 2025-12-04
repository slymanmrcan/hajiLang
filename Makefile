.PHONY: build run test install clean help build-all

# Varsayılan hedef
help:
	@echo "HajiLang Makefile Komutları:"
	@echo "  make build      - Binary'yi derle"
	@echo "  make build-all  - Tüm platformlar için derle"
	@echo "  make run        - test.haji'yi çalıştır"
	@echo "  make repl       - REPL modunu başlat"
	@echo "  make install    - Sistem geneline kur"
	@echo "  make clean      - Binary'yi sil"
	@echo "  make test       - Testleri çalıştır"

# Binary'yi derle
build:
	@echo "🔨 Derleniyor..."
	go build -o hajilang
	@echo "✅ Hazır: ./hajilang"

# test.haji'yi çalıştır
run:
	@echo "▶️  test.haji çalıştırılıyor..."
	go run main.go test.haji

# REPL başlat
repl:
	@echo "🚀 REPL başlatılıyor..."
	go run main.go

# Sistem geneline kur
install: build
	@echo "📦 /usr/local/bin'e kuruluyor..."
	sudo cp hajilang /usr/local/bin/
	@echo "✅ Kurulum tamamlandı!"
	@echo "   Artık 'hajilang' komutunu her yerden kullanabilirsin"

# Temizle
clean:
	@echo "🧹 Temizleniyor..."
	rm -f hajilang
	@echo "✅ Temizlendi"

# Testleri çalıştır (şimdilik basit)
test:
	@echo "🧪 Testler çalıştırılıyor..."
	go test ./...

# Tüm platformlar için derle
build-all:
	@echo "📦 Tüm platformlar için derleniyor..."
	GOOS=darwin GOARCH=amd64 go build -o dist/hajilang-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o dist/hajilang-darwin-arm64
	GOOS=linux GOARCH=amd64 go build -o dist/hajilang-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o dist/hajilang-linux-arm64
	GOOS=windows GOARCH=amd64 go build -o dist/hajilang-windows-amd64.exe
	@echo "✅ Tamamlandı: dist/ klasöründe"