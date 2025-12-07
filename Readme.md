# 🚀 HajiLang

Go ile yazılmış, Türkçe sözdizimine sahip, modern ve eğlenceli bir programlama dili yorumlayıcısı (Interpreter).

![HajiLang Logo](https://via.placeholder.com/800x200?text=HajiLang)

## ✨ Özellikler

HajiLang, modern bir programlama dilinden beklenen temel özellikleri destekler:

- **Değişkenler:** `haji` (değişken) ve `kati` (sabit) tanımları
- **Fonksiyonlar:** First-class fonksiyonlar, closure desteği ve high-order fonksiyonlar
- **Döngüler:** C-tarzı `for` döngüleri
- **Veri Yapıları:** Dinamik diziler (Array) ve Hash map'ler
- **Operatörler:** Aritmetik (`+`, `-`, `*`, `/`, `%`), Mantıksal (`&&`, `||`), Karşılaştırma (`<=`, `>=`)
- **Koşullar:** `if`, `else if`, `else` blokları
- **REPL:** Renkli, geçmiş destekli ve otomatik tamamlamalı interaktif konsol
- **Hata Yönetimi:** Satır numaralı detaylı hata mesajları

## 📦 Kurulum

### 1. Go ile (Geliştiriciler İçin)

```bash
# Repoyu klonla
git clone https://github.com/slymanmrcan/hajilang.git
cd hajilang

# Derle ve çalıştır
go run main.go
```

### 2. Binary Kullanımı

Releases sayfasından işletim sistemine uygun sürümü indirebilirsiniz.

```bash
# Linux/macOS
chmod +x hajilang-linux-amd64
./hajilang-linux-amd64 test.haji

# Windows
hajilang-windows-amd64.exe test.haji
```

## 🎮 Kullanım Örnekleri

### 1. Merhaba Dünya ve Değişkenler
```javascript
haji isim = "Dünya"
yaz("Merhaba " + isim) // Merhaba Dünya

kati PI = 3.14159
// PI = 3.14 // HATA: Sabit değiştirilemez!
```

### 2. Fonksiyonlar ve Closure
```javascript
haji topla = fn(a, b) {
    return a + b
}
yaz(topla(5, 10)) // 15

// Closure Örneği
haji sayacYap = fn() {
    haji i = 0
    return fn() {
        i = i + 1
        return i
    }
}

haji sayac = sayacYap()
yaz(sayac()) // 1
yaz(sayac()) // 2
```

### 3. Döngüler ve Karşılaştırma
```javascript
haji toplam = 0
for (haji i = 1; i <= 10; i = i + 1) {
    if (i % 2 == 0) {
        toplam = toplam + i
        yaz(i, "çifttir")
    }
}
yaz("Toplam:", toplam)
```

### 4. Diziler ve Haritalar
```javascript
haji sayilar = [1, 2, 3, 4]
yaz(len(sayilar))    // 4
yaz(first(sayilar))  // 1
yaz(push(sayilar, 5)) // [1, 2, 3, 4, 5]

haji sozluk = {"ad": "Ali", "yas": 25}
yaz(sozluk["ad"]) // Ali
```

## 🛠️ VS Code Eklentisi

HajiLang kodlarınızı renklendirmek için VS Code eklentisi mevcuttur.

1. `vscode/` klasörünü VS Code ile açın.
2. `F5` tuşuna basarak eklentiyi test modunda başlatın.
3. `.haji` uzantılı dosyalarınız artık renkli!

## 🤝 Katkıda Bulunma

1. Forklayın
2. Feature branch oluşturun (`git checkout -b ozellik/yeni-ozellik`)
3. Commit leyin (`git commit -m 'Yeni özellik eklendi'`)
4. Pushlayın (`git push origin ozellik/yeni-ozellik`)
5. Pull Request açın

## 📝 Lisans

MIT License ile lisanslanmıştır.