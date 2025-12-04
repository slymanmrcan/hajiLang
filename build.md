# Derleme ve Çalıştırma

Önkoşul: Go 1.25+.

## Hızlı Komutlar
```bash
make build   # ./hajilang üretir
make repl    # REPL
make run     # test.haji'yi çalıştırır
make test    # go test ./...
make install # /usr/local/bin/hajilang (sudo ister)
```

## Elle Derleme
```bash
go build -o hajilang main.go
./hajilang           # REPL
./hajilang file.haji # dosya çalıştır
```

## Çapraz Derleme
```bash
GOOS=linux   GOARCH=amd64 go build -o hajilang-linux main.go
GOOS=windows GOARCH=amd64 go build -o hajilang.exe   main.go
```

İkiliyi PATH'ine koyarak her yerden `hajilang` çalıştırabilirsin. Çalışma zamanı ek bağımlılık yok; tek okuma hedefi verdiğin `.haji` dosyası.
