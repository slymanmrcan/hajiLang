# HajiLang

Go ile yazılmış, eğlence amaçlı küçük bir yorumlayıcı.

## Kurulum

### 1. Go ile (önerilen)

Go yüklüyse:

```bash
go install github.com/slymanmrcan/hajilang@latest
```

### 2. Otomatik Script ile

**Linux/macOS:**
```bash
curl -sSL https://raw.githubusercontent.com/slymanmrcan/hajilang/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/slymanmrcan/hajilang/main/install.ps1 | iex
```

### 3. Manuel Kurulum

[Releases](https://github.com/slymanmrcan/hajilang/releases/latest) sayfasından işletim sistemine uygun binary'yi indir.

**Linux/macOS:**
```bash
# İndir
wget https://github.com/slymanmrcan/hajilang/releases/latest/download/hajilang-linux-amd64
chmod +x hajilang-linux-amd64
sudo mv hajilang-linux-amd64 /usr/local/bin/hajilang
```

**Windows:**
Binary'yi indir, bir klasöre koy ve PATH'e ekle.

## Kullanım

```bash
# REPL modu
hajilang

# Dosyadan çalıştır
hajilang dosya.haji
```

## Özellikler

- ✅ Let ile değişken tanımlama
- ✅ Integer ve Boolean
- ✅ Aritmetik operatörler (+, -, *, /)
- ✅ Karşılaştırma (<, >, ==, !=)
- ✅ If/Else
- ✅ Yorumlar (//)
- ✅ REPL (interaktif mod)

## Örnek

```haji
// Basit matematik ve koşul
let x = 12 + 3 * 2;
let y = 10;

if (x > y) {
    x - y;
} else {
    0;
}
```

## Geliştirme

```bash
# Klonla
git clone https://github.com/slymanmrcan/hajilang
cd hajilang

# Derle
make build

# Test et
make run

# REPL
make repl
```

## Lisans

MIT