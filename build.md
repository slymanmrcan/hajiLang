# HajiLang Derleme ve Çalıştırma Rehberi

Bu proje Go ile yazılmış bir yorumlayıcıdır ve tek bir ikili dosya üretir. Aşağıdaki komutlar, farklı platformlar için nasıl derleyip çalıştıracağını gösterir.

## macOS (geliştirme)
```bash
go build -o hajilang main.go
./hajilang test.haji
```

## Linux için çapraz derleme
```bash
GOOS=linux GOARCH=amd64 go build -o hajilang-linux main.go
chmod +x hajilang-linux
./hajilang-linux test.haji
```

## Windows için çapraz derleme
```bash
GOOS=windows GOARCH=amd64 go build -o hajilang.exe main.go
hajilang.exe test.haji
```

## Notlar
- İkili dosyayı istediğin klasöre kopyalayabilir ve PATH'ine ekleyebilirsin.
- `./hajilang` yalnızca verdiğin `.haji` dosyasını okur; ek bir yapıma ihtiyaç duymaz.
- Çalışma zamanı hataları ve sözdizimi hataları konsola yazılır.
