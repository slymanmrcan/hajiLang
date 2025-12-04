# HajiLang

HajiLang, Go ile yazılmış, eğlence amaçlı küçük bir yorumlayıcıdır. Lexer, parser, AST ve evaluator katmanlarını baştan kurarak basit bir dilin nasıl çalıştığını gösterir.

## Neler Destekleniyor?
- `let` ile değişken tanımlama ve yeniden değerlendirme (tek tip atama formu)
- Tamsayı ve boolean değerler; tek satır yorumlar (`// ...`)
- Aritmetik: `+`, `-`, `*`, `/`; unary `-` ve `!`
- Karşılaştırmalar: `<`, `>`, `==`, `!=` (sayılar ve booleanlarda)
- Parantezlerle öncelik kontrolü
- `if { ... } else { ... }` blokları; koşul sağlanmazsa `nil` dönme
- Hata yakalama: sözdizimi hataları, tanımsız değişken, tür uyumsuzluğu, sıfıra bölme

Henüz stringler, fonksiyonlar ve `return` gibi yapılar yok; dil kasıtlı olarak küçük tutuldu.

## Proje Yapısı
- `lexer/`, `parser/`, `ast/`: Ön uç ve ağaç üretimi
- `evaluator/`: AST yorumlayıcısı (çalıştırıcı)
- `object/`: Çalışma zamanı türleri ve ortam (environment)
- `token/`: Token tanımları ve anahtar kelimeler
- `main.go`: Dosya okuma, derleme hattı ve yürütme

## Hızlı Başlangıç
Ön koşul: Go 1.21+.

```bash
# ikiliyi üret
go build -o hajilang main.go

# bir programı çalıştır
./hajilang test.haji
```

`./hajilang` bir dosya yolu bekler; kodda hata varsa konsola raporlar.

## Örnek Program
```javascript
// basit matematik ve koşul
let x = 12 + 3 * 2;
let y = 10;

if (x > y) {
  x - y;
} else {
  0;
}
```

## Dil Kuralları (Kısa Özet)
- İfadeler `;` ile sonlanabilir; parser son satırı da kabul eder ama tutarlılık için noktalı virgül kullanın.
- Parantezler önceliği değiştirir, aksi halde `*`/`/` > `+`/`-` > karşılaştırma > `==`/`!=` > unary.
- `if` koşulu parantez ister; bloklar `{ ... }` içinde yazılır; `else` isteğe bağlıdır.
- Yorumlar `//` ile başlar ve satır sonuna kadar gider; evaluator tarafından yok sayılır.

## Yol Haritası
- `return` ve bloktan erken çıkış
- Fonksiyon tanımlama/çağırma
- Daha fazla tür (string, dizi, map) ve yerleşik fonksiyonlar
