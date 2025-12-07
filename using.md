# HajiLang Kullanım Kılavuzu

## 🚀 Hızlı Başlangıç

```bash
hajilang            # REPL (interaktif mod)
hajilang test.haji  # Dosya çalıştır
```

---

## 📝 Değişkenler

### `haji` - Değişken Tanımlama
```javascript
haji isim = "Süleyman"
haji yas = 25
haji aktif = true
```

### `kati` - Sabit Tanımlama
```javascript
kati PI = 3.14159
kati MAX = 100
// kati değişkenler değiştirilemez!
```

---

## 🔢 Operatörler

| Operatör | Açıklama | Örnek |
|----------|----------|-------|
| `+` | Toplama / Birleştirme | `5 + 3` → `8`, `"a" + "b"` → `"ab"` |
| `-` | Çıkarma | `10 - 4` → `6` |
| `*` | Çarpma | `3 * 4` → `12` |
| `/` | Bölme | `10 / 2` → `5` |
| `==` | Eşitlik | `5 == 5` → `true` |
| `!=` | Eşitsizlik | `5 != 3` → `true` |
| `<` | Küçük | `3 < 5` → `true` |
| `>` | Büyük | `5 > 3` → `true` |
| `!` | Değil | `!true` → `false` |
| `-` | Negatif | `-5` |

---

## 🔀 Kontrol Akışı

### If/Else
```javascript
if (x > 10) {
    yaz("Büyük")
} else if (x > 5) {
    yaz("Orta")
} else {
    yaz("Küçük")
}
```

### For Döngüsü
```javascript
for (haji i = 0; i < 5; i = i + 1) {
    yaz(i)
}

// Dış değişken kullanımı
haji toplam = 0
for (haji j = 1; j < 11; j = j + 1) {
    toplam = toplam + j
}
yaz("Toplam:", toplam)  // 55
```

---

## 📦 Fonksiyonlar

### Fonksiyon Tanımlama
```javascript
haji topla = fn(a, b) {
    return a + b
}

yaz(topla(3, 5))  // 8
```

### Parametresiz Fonksiyon
```javascript
haji selamla = fn() {
    yaz("Merhaba Dünya!")
}

selamla()
```

### Closure (İç İçe Fonksiyon)
```javascript
haji carpici = fn(x) {
    return fn(y) {
        return x * y
    }
}

haji ikiKati = carpici(2)
yaz(ikiKati(5))   // 10
yaz(ikiKati(10))  // 20
```

---

## 📚 Diziler ve Hash

### Dizi (Array)
```javascript
haji sayilar = [1, 2, 3, 4, 5]
yaz(sayilar[0])        // 1
yaz(len(sayilar))      // 5
yaz(first(sayilar))    // 1
yaz(last(sayilar))     // 5
yaz(rest(sayilar))     // [2, 3, 4, 5]
yaz(push(sayilar, 6))  // [1, 2, 3, 4, 5, 6]
```

### Hash (Sözlük)
```javascript
haji kisi = {"isim": "Ali", "yas": 30}
yaz(kisi["isim"])  // Ali
```

---

## 🔧 Gömülü Fonksiyonlar

| Fonksiyon | Açıklama | Örnek |
|-----------|----------|-------|
| `yaz(...)` | Ekrana yazdır | `yaz("Merhaba", 42)` |
| `puts(...)` | Ekrana yazdır (satır sonu) | `puts("Merhaba")` |
| `len(x)` | Uzunluk | `len("abc")` → `3` |
| `first(arr)` | İlk eleman | `first([1,2,3])` → `1` |
| `last(arr)` | Son eleman | `last([1,2,3])` → `3` |
| `rest(arr)` | İlk hariç | `rest([1,2,3])` → `[2,3]` |
| `push(arr, x)` | Eleman ekle | `push([1,2], 3)` → `[1,2,3]` |
| `to_int(s)` | Stringe çevir | `to_int("42")` → `42` |
| `to_str(n)` | Stringe çevir | `to_str(42)` → `"42"` |

---

## 💡 Örnek Program

```javascript
// Faktöriyel hesaplama
haji faktoriyel = fn(n) {
    if (n < 2) {
        return 1
    }
    return n * faktoriyel(n - 1)
}

yaz("5! =", faktoriyel(5))  // 120

// Fibonacci
haji fib = fn(n) {
    if (n < 2) {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}

for (haji i = 0; i < 10; i = i + 1) {
    yaz("fib(" + to_str(i) + ") =", fib(i))
}
```

---

## 🎯 Sözdizimi Özeti

```
haji x = 5          // Değişken
kati PI = 3.14      // Sabit
fn(a, b) { ... }    // Fonksiyon
if (...) { } else { } // Koşul
for (...; ...; ...) { } // Döngü
return x            // Değer döndür
[1, 2, 3]           // Dizi
{"a": 1}            // Hash
// yorum            // Yorum satırı
```
