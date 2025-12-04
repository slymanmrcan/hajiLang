# HajiLang Kullanım Rehberi

## Temel Sözdizimi
- Değişken tanımı: `let isim = ifade;`
- İfadeler noktalı virgül ile bitmeli; parser son ifadeyi noktalı virgülsüz de okur ama tutarlılık için kullanın.
- Bloklar `{ ... }` ile sarılır; girinti zorunlu değildir.
- Boolean literaller: `true`, `false`.
- Tek satır yorum: `// yorum`.

## Operatörler ve Öncelik
- Unary: `-x`, `!x`
- Çarpma/bölme: `*`, `/`
- Toplama/çıkarma: `+`, `-`
- Karşılaştırma: `<`, `>`, `==`, `!=`
- Parantez önceliği değiştirir; aksi halde yukarıdaki sırayla değerlendirilir.

## Akış Kontrolü
```javascript
if (kosul) {
  // true durumunda son ifade döner
} else {
  // opsiyonel; yoksa sonuç nil olur
}
```

## Çalışan Örnek
```javascript
// tek satır yorum
let limit = 100;
let deger = 12 * 3 + 4;

if (deger < limit) {
  deger + 1;
} else {
  0;
}
```

## Anahtar Kelimeleri Özelleştirme
- Sözcükleri Türkçeleştirmek için `token/token.go` içindeki `keywords` haritasını değiştir (örn. `let` yerine `olsun`).
- Sembol tercihlerini değiştirmek istersen `lexer/lexer.go` içindeki karakter kontrollerini düzenle (`{`, `}`, `=` vb.).
- Dilin gramerini (ör. `if` parantezsiz olsun) değiştirmek parser tarafında kapsamlı düzenleme gerektirir: `parser/parser.go`.

## Neler Eksik?
- Fonksiyonlar, `return`, string/dizi/harita tipleri henüz yok.
- Hata mesajları konsola yazılır; çalışma durur veya `nil` döner.
