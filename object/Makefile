.PHONY: build run test install clean help

# Varsayılan hedef
help:
	@echo "HajiLang Makefile Komutları:"
	@echo "  make build    - Binary'yi derle"
	@echo "  make run      - test.haji'yi çalıştır"
	@echo "  make repl     - REPL modunu başlat"
	@echo "  make install  - Sistem geneline kur"
	@echo "  make clean    - Binary'yi sil"
	@echo "  make test     - Testleri çalıştır"

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