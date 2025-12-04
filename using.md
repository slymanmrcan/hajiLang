# HajiLang Kılavuz

## Sözdizimi
- Değişken: `let isim = ifade;`
- Literaller: `123`, `true`, `false`
- Yorum: `// satır sonuna kadar`
- Bloklar: `{ ... }`

## Operatörler ve Öncelik
- Unary: `-x`, `!x`
- Çarpma/bölme: `*`, `/`
- Toplama/çıkarma: `+`, `-`
- Karşılaştırma: `<`, `>`, `==`, `!=`
- Öncelik sırası yukarıdaki gibi; parantez önceliği değiştirir.

## Akış Kontrolü
```javascript
if (kosul) {
  // koşul true ise bu bloğun son ifadesi dönüyor
} else {
  // opsiyonel; yoksa sonuç nil
}
```

## Çalışan Örnek
```javascript
let limit = 100;
let deger = 12 * 3 + 4;

if (deger < limit) {
  deger + 1;
} else {
  0;
}
```

## Kullanım
```bash
hajilang            # REPL (durumlu environment)
hajilang foo.haji   # dosyadan yürütme
```

## Özelleştirme
- Anahtar kelimeler: `token/token.go` içindeki `keywords` haritasını değiştir.
- Semboller: `lexer/lexer.go` karakter kontrollerini güncelle (`{`, `}`, `=` vb.).
- Gramer genişletmeleri: `parser/parser.go` (ör. yeni ifade/operatör eklemek).

## Şu Anki Sınırlar
String/koleksiyon tipleri, fonksiyonlar ve `return` yok. Kapsam tek seviyeli (global environment). Mantıksal `&&/||` henüz desteklenmiyor.
