Mevcut Yapı
Projen klasik interpreter mimarisini takip ediyor:

Lexer (lexer/) - Kaynak kodu tokenlara ayırır (INT, PLUS, IF gibi)
Parser (parser/) - Tokenleri AST'ye (Abstract Syntax Tree) dönüştürür
AST (ast/) - Programın ağaç yapısını temsil eder
Evaluator (evaluator/) - AST'yi yorumlayıp çalıştırır
Object (object/) - Çalışma zamanı değerlerini tutar (Integer, Boolean, Nil)
# HajiLang - Detaylı Mimari Açıklaması

## 📚 İçindekiler
1. [Interpreter Nedir?](#interpreter-nedir)
2. [HajiLang'in 5 Katmanlı Yapısı](#hajilang-5-katman)
3. [Her Katman Detaylı Açıklama](#katmanlar-detay)
4. [Yeni Özellik Ekleme Rehberi](#yeni-ozellik)
5. [Go Dili Özellikleri](#go-ozellikleri)

---

## 🎯 Interpreter Nedir? {#interpreter-nedir}

Bir programlama dili çalıştırmak için 2 yol var:

1. **Compiler** (Derleyici): Kodu makine diline çevirir (C, Rust gibi)
2. **Interpreter** (Yorumlayıcı): Kodu satır satır okuyup çalıştırır (Python, JavaScript gibi)

HajiLang bir **interpreter**. Yani `let x = 5 + 3;` yazdığında:
- Okur
- Anlar
- Hesaplar
- Sonucu döner

---

## 🏗️ HajiLang'in 5 Katmanlı Yapısı {#hajilang-5-katman}

```
KAYNAK KOD (test.haji)
        ↓
[1] LEXER (Tokenizer)
        ↓
[2] PARSER (Syntax Analyzer)
        ↓
[3] AST (Abstract Syntax Tree)
        ↓
[4] EVALUATOR (Interpreter)
        ↓
[5] RESULT (Sonuç)
```

### Örnek: `let x = 5 + 3;`

```
[1] LEXER:     [LET] [x] [=] [5] [+] [3] [;]
[2] PARSER:    LetStatement { name: x, value: InfixExpression }
[3] AST:       Tree yapısı oluşturur
[4] EVALUATOR: 5 + 3 = 8 hesaplar, x'e atar
[5] RESULT:    x = 8
```

---

## 📦 Katmanlar Detaylı Açıklama {#katmanlar-detay}

---

## [1] TOKEN - Temel Yapı Taşları

**Dosya:** `token/token.go`

### Token Nedir?

Token, dilin en küçük anlamlı parçasıdır. Kelimelere benzer.

```go
type TokenType string  // Token'ın tipi (INT, PLUS, IF...)

type Token struct {
    Type    TokenType  // Ne tür bir token?
    Literal string     // Gerçek değeri ne?
}
```

### Örnek:

```
Kod:   let x = 42;

Token'lar:
┌─────────┬─────────┐
│  Type   │ Literal │
├─────────┼─────────┤
│  LET    │  "let"  │
│  IDENT  │  "x"    │
│  ASSIGN │  "="    │
│  INT    │  "42"   │
│  SEMI   │  ";"    │
└─────────┴─────────┘
```

### Token Tipleri:

```go
const (
    // Özel
    ILLEGAL = "ILLEGAL"  // Tanınmayan karakter
    EOF     = "EOF"      // Dosya sonu
    
    // Tanımlayıcılar
    IDENT = "IDENT"      // değişken isimleri (x, y, foo)
    INT   = "INT"        // tamsayılar (5, 42)
    STRING = "STRING"    // string'ler ("merhaba")
    
    // Operatörler
    ASSIGN = "="
    PLUS   = "+"
    MINUS  = "-"
    
    // Anahtar Kelimeler
    LET   = "let"
    IF    = "if"
    ELSE  = "else"
)
```

---

## [2] LEXER - Token Üretici

**Dosya:** `lexer/lexer.go`

### Lexer Nedir?

Metni okuyup token'lara ayıran makine. Kitabı kelimelere ayırmak gibi.

```go
type Lexer struct {
    input        string  // Kaynak kod
    position     int     // Şu anki pozisyon
    readPosition int     // Bir sonraki pozisyon
    ch           byte    // Şu anki karakter
}
```

### Nasıl Çalışır?

```go
func (l *Lexer) NextToken() token.Token {
    l.skipWhitespace()  // Boşlukları atla
    
    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            return token.Token{Type: token.EQ, Literal: "=="}
        }
        return token.Token{Type: token.ASSIGN, Literal: "="}
    
    case '+':
        return token.Token{Type: token.PLUS, Literal: "+"}
    
    case '"':
        return token.Token{
            Type: token.STRING, 
            Literal: l.readString()
        }
    }
}
```

### Örnek Akış:

```
Input: "let x = 5;"
       ↓
Position: 0, Char: 'l'
       ↓
readIdentifier() → "let"
       ↓
Token: {Type: LET, Literal: "let"}
       ↓
Position: 4, Char: 'x'
       ↓
readIdentifier() → "x"
       ↓
Token: {Type: IDENT, Literal: "x"}
```

### String Okuma:

```go
func (l *Lexer) readString() string {
    position := l.position + 1  // " işaretinden sonra başla
    
    for {
        l.readChar()
        if l.ch == '"' || l.ch == 0 {  // Kapanış " veya EOF
            break
        }
    }
    
    return l.input[position:l.position]  // Tırnaklar hariç
}
```

**Örnek:**
```
Input: "merhaba"
       ^       ^
       |       |
     pos+1   position
       
Result: "merhaba" (tırnaksız)
```

---

## [3] AST - Ağaç Yapısı

**Dosya:** `ast/ast.go`

### AST Nedir?

Abstract Syntax Tree (Soyut Sözdizim Ağacı). Kodun anlamını ağaç şeklinde gösterir.

```
Kod: let x = 5 + 3;

AST:
      LetStatement
          ├── Name: "x"
          └── Value: InfixExpression
                  ├── Left: IntegerLiteral(5)
                  ├── Operator: "+"
                  └── Right: IntegerLiteral(3)
```

### Temel Interface:

```go
type Node interface {
    TokenLiteral() string  // Token'ın string değeri
    String() string        // Debug için
}

type Statement interface {
    Node
    statementNode()  // Bu bir statement (let, if, return...)
}

type Expression interface {
    Node
    expressionNode()  // Bu bir expression (5+3, x, "merhaba")
}
```

### Statement vs Expression

**Statement:** Bir şey yapar, değer döndürmez
```go
let x = 5;        // Değişken tanımla
if (x > 3) { }    // Koşul kontrol et
```

**Expression:** Değer üretir
```go
5 + 3             // 8 döner
x > 10            // true/false döner
"a" + "b"         // "ab" döner
```

### Önemli AST Yapıları:

#### 1. IntegerLiteral (Sayı)

```go
type IntegerLiteral struct {
    Token token.Token  // token.INT
    Value int64        // Gerçek sayı değeri
}

// Örnek: 42
&IntegerLiteral{
    Token: {Type: "INT", Literal: "42"},
    Value: 42
}
```

#### 2. StringLiteral (String)

```go
type StringLiteral struct {
    Token token.Token  // token.STRING
    Value string       // String içeriği
}

// Örnek: "merhaba"
&StringLiteral{
    Token: {Type: "STRING", Literal: "merhaba"},
    Value: "merhaba"
}
```

#### 3. InfixExpression (İkili İşlem)

```go
type InfixExpression struct {
    Token    token.Token  // Operatör token'ı (+, -, *, /)
    Left     Expression   // Sol taraf
    Operator string       // Operatör ("+", "-", "*", "/")
    Right    Expression   // Sağ taraf
}

// Örnek: 5 + 3
&InfixExpression{
    Left: &IntegerLiteral{Value: 5},
    Operator: "+",
    Right: &IntegerLiteral{Value: 3}
}
```

#### 4. IfExpression (Koşul)

```go
type IfExpression struct {
    Token       token.Token      // 'if' token'ı
    Condition   Expression       // Koşul
    Consequence *BlockStatement  // True bloku
    Alternative *BlockStatement  // Else bloku (opsiyonel)
}

// Örnek: if (x > 5) { 10; } else { 20; }
&IfExpression{
    Condition: &InfixExpression{...},  // x > 5
    Consequence: &BlockStatement{...}, // { 10; }
    Alternative: &BlockStatement{...}  // { 20; }
}
```

---

## [4] PARSER - Ağaç Oluşturucu

**Dosya:** `parser/parser.go`

### Parser Nedir?

Token'ları alıp AST ağacı oluşturan katman.

```go
type Parser struct {
    l      *lexer.Lexer  // Lexer referansı
    errors []string      // Hata listesi
    
    curToken  token.Token  // Şu anki token
    peekToken token.Token  // Bir sonraki token
    
    // FONKSİYON MAP'LERİ - ÇOK ÖNEMLİ!
    prefixParseFns map[token.TokenType]prefixParseFn
    infixParseFns  map[token.TokenType]infixParseFn
}
```

### Pratt Parsing - Fonksiyon Map Sistemi

Bu sistem sayesinde yeni token tipi eklemek çok kolay!

#### Fonksiyon Tipleri:

```go
type (
    // Prefix: Token'ın başında olur (-5, !true, "merhaba")
    prefixParseFn func() ast.Expression
    
    // Infix: İki değer arasında olur (5 + 3, x == 10)
    infixParseFn func(ast.Expression) ast.Expression
)
```

#### Map Sistemi:

```go
func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l, errors: []string{}}
    
    // PREFIX FONKSİYONLARI KAYDET
    p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
    
    p.registerPrefix(token.INT, p.parseIntegerLiteral)
    //                ↑                ↑
    //            Token Tipi      Çalışacak Fonksiyon
    
    p.registerPrefix(token.STRING, p.parseStringLiteral)
    p.registerPrefix(token.IDENT, p.parseIdentifier)
    p.registerPrefix(token.BANG, p.parsePrefixExpression)
    p.registerPrefix(token.MINUS, p.parsePrefixExpression)
    
    // INFIX FONKSİYONLARI KAYDET
    p.infixParseFns = make(map[token.TokenType]infixParseFn)
    
    p.registerInfix(token.PLUS, p.parseInfixExpression)
    p.registerInfix(token.MINUS, p.parseInfixExpression)
    p.registerInfix(token.ASTERISK, p.parseInfixExpression)
    
    return p
}
```

#### parseExpression - Ana Fonksiyon

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
    // 1. PREFIX: Başlangıç token'ını işle
    prefix := p.prefixParseFns[p.curToken.Type]
    //         ↑
    //    Map'ten fonksiyonu al
    
    if prefix == nil {
        p.noPrefixParseFnError(p.curToken.Type)
        return nil
    }
    
    leftExp := prefix()  // Fonksiyonu çalıştır
    //         ↑
    //    parseIntegerLiteral() veya parseStringLiteral() vs.
    
    // 2. INFIX: Operatörleri soldan sağa işle
    for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        
        p.nextToken()
        leftExp = infix(leftExp)
        //        ↑
        //   parseInfixExpression(leftExp)
    }
    
    return leftExp
}
```

#### Örnek Akış: `5 + 3`

```
Adım 1: curToken = INT(5)
        ↓
prefix = prefixParseFns[INT] = parseIntegerLiteral
        ↓
leftExp = parseIntegerLiteral() = IntegerLiteral{Value: 5}

Adım 2: peekToken = PLUS
        ↓
infix = infixParseFns[PLUS] = parseInfixExpression
        ↓
nextToken() → curToken = PLUS
        ↓
leftExp = parseInfixExpression(IntegerLiteral{5})
        ↓
InfixExpression{
    Left: IntegerLiteral{5},
    Operator: "+",
    Right: parseExpression() → IntegerLiteral{3}
}
```

#### Parse Fonksiyonları:

```go
// INTEGER parse et
func (p *Parser) parseIntegerLiteral() ast.Expression {
    value, _ := strconv.ParseInt(p.curToken.Literal, 0, 64)
    return &ast.IntegerLiteral{
        Token: p.curToken,
        Value: value
    }
}

// STRING parse et
func (p *Parser) parseStringLiteral() ast.Expression {
    return &ast.StringLiteral{
        Token: p.curToken,
        Value: p.curToken.Literal  // Lexer zaten tırnakları temizledi
    }
}

// INFIX parse et (5 + 3)
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
    expression := &ast.InfixExpression{
        Token:    p.curToken,
        Operator: p.curToken.Literal,
        Left:     left,  // Sol taraf zaten parse edilmiş
    }
    
    precedence := p.curPrecedence()
    p.nextToken()
    expression.Right = p.parseExpression(precedence)  // Sağ tarafı parse et
    
    return expression
}
```

---

## [5] EVALUATOR - Yorumlayıcı

**Dosya:** `evaluator/evaluator.go`

### Evaluator Nedir?

AST ağacını dolaşıp hesaplamalar yapan katman. Gerçek çalıştırma burada oluyor.

```go
func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    
    case *ast.StringLiteral:
        return &object.String{Value: node.Value}
    
    case *ast.InfixExpression:
        left := Eval(node.Left, env)    // Sol tarafı hesapla
        right := Eval(node.Right, env)  // Sağ tarafı hesapla
        return evalInfixExpression(node.Operator, left, right)
    
    case *ast.LetStatement:
        val := Eval(node.Value, env)  // Değeri hesapla
        env.Set(node.Name.Value, val) // Değişkene ata
        return nil
    
    case *ast.Identifier:
        return env.Get(node.Value)  // Değişkenin değerini al
    }
}
```

### Environment - Değişken Hafızası

```go
type Environment struct {
    store map[string]object.Object  // Değişken deposu
}

func (e *Environment) Get(name string) (object.Object, bool) {
    obj, ok := e.store[name]
    return obj, ok
}

func (e *Environment) Set(name string, val object.Object) {
    e.store[name] = val
}
```

**Örnek:**
```
let x = 5;
let y = x + 3;

Environment:
┌────┬───────────────────┐
│ x  │ Integer{Value: 5} │
│ y  │ Integer{Value: 8} │
└────┴───────────────────┘
```

### Infix Expression Değerlendirme

```go
func evalInfixExpression(operator string, left, right object.Object) object.Object {
    
    // STRING + STRING → Birleştir
    if left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ {
        leftVal := left.(*object.String).Value
        rightVal := right.(*object.String).Value
        
        if operator == "+" {
            return &object.String{Value: leftVal + rightVal}
        }
    }
    
    // INTEGER + INTEGER → Topla
    if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
        leftVal := left.(*object.Integer).Value
        rightVal := right.(*object.Integer).Value
        
        switch operator {
        case "+":
            return &object.Integer{Value: leftVal + rightVal}
        case "-":
            return &object.Integer{Value: leftVal - rightVal}
        case "*":
            return &object.Integer{Value: leftVal * rightVal}
        case "/":
            if rightVal == 0 {
                return newError("sıfıra bölünemez!")
            }
            return &object.Integer{Value: leftVal / rightVal}
        }
    }
    
    return newError("tür uyumsuzluğu")
}
```

---

## [6] OBJECT - Çalışma Zamanı Değerleri

**Dosya:** `object/object.go`

### Object Nedir?

Program çalışırken bellekte tutulan değerler.

```go
type ObjectType string

const (
    INTEGER_OBJ = "INTEGER"
    STRING_OBJ  = "STRING"
    BOOLEAN_OBJ = "BOOLEAN"
    NULL_OBJ    = "NULL"
    ERROR_OBJ   = "ERROR"
)

type Object interface {
    Type() ObjectType
    Inspect() string  // Debug için string dönüşümü
}
```

### Object Tipleri:

```go
// INTEGER
type Integer struct {
    Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// STRING
type String struct {
    Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// BOOLEAN
type Boolean struct {
    Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// ERROR
type Error struct {
    Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "HATA: " + e.Message }
```

---

## 🚀 Yeni Özellik Ekleme Rehberi {#yeni-ozellik}

### Örnek: ARRAY (Dizi) Ekleme

#### Adım 1: Token Ekle
```go
// token/token.go
const (
    LBRACKET = "["
    RBRACKET = "]"
)
```

#### Adım 2: Lexer'da Token'ı Tanı
```go
// lexer/lexer.go
func (l *Lexer) NextToken() token.Token {
    switch l.ch {
    case '[':
        tok = newToken(token.LBRACKET, l.ch)
    case ']':
        tok = newToken(token.RBRACKET, l.ch)
    }
}
```

#### Adım 3: AST Node Ekle
```go
// ast/ast.go
type ArrayLiteral struct {
    Token    token.Token   // '[' token'ı
    Elements []Expression  // [1, 2, "a"] gibi elemanlar
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
    elements := []string{}
    for _, el := range al.Elements {
        elements = append(elements, el.String())
    }
    return "[" + strings.Join(elements, ", ") + "]"
}
```

#### Adım 4: Parser'a Ekle
```go
// parser/parser.go

// New() fonksiyonuna ekle
p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)

// Parse fonksiyonu yaz
func (p *Parser) parseArrayLiteral() ast.Expression {
    array := &ast.ArrayLiteral{Token: p.curToken}
    array.Elements = p.parseExpressionList(token.RBRACKET)
    return array
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
    list := []ast.Expression{}
    
    if p.peekTokenIs(end) {
        p.nextToken()
        return list
    }
    
    p.nextToken()
    list = append(list, p.parseExpression(LOWEST))
    
    for p.peekTokenIs(token.COMMA) {
        p.nextToken()
        p.nextToken()
        list = append(list, p.parseExpression(LOWEST))
    }
    
    if !p.expectPeek(end) {
        return nil
    }
    
    return list
}
```

#### Adım 5: Object Ekle
```go
// object/object.go
const ARRAY_OBJ = "ARRAY"

type Array struct {
    Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
    elements := []string{}
    for _, e := range ao.Elements {
        elements = append(elements, e.Inspect())
    }
    return "[" + strings.Join(elements, ", ") + "]"
}
```

#### Adım 6: Evaluator'a Ekle
```go
// evaluator/evaluator.go
func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    
    case *ast.ArrayLiteral:
        elements := evalExpressions(node.Elements, env)
        if len(elements) == 1 && isError(elements[0]) {
            return elements[0]
        }
        return &object.Array{Elements: elements}
    }
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
    result := []object.Object{}
    
    for _, e := range exps {
        evaluated := Eval(e, env)
        if isError(evaluated) {
            return []object.Object{evaluated}
        }
        result = append(result, evaluated)
    }
    
    return result
}
```

#### Kullanım:
```
let arr = [1, 2, "merhaba", true];
arr;  // [1, 2, merhaba, true]
```

---

## 🔧 Go Dili Özellikleri {#go-ozellikleri}

### 1. Struct (Yapı)

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Ali", Age: 25}
fmt.Println(p.Name)  // Ali
```

### 2. Interface (Arayüz)

```go
type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Hav!"
}

var animal Animal = Dog{}
animal.Speak()  // "Hav!"
```

### 3. Method (Metod)

```go
type Rectangle struct {
    Width  int
    Height int
}

// Receiver: (r Rectangle)
func (r Rectangle) Area() int {
    return r.Width * r.Height
}

rect := Rectangle{Width: 10, Height: 5}
rect.Area()  // 50
```

### 4. Pointer (İşaretçi)

```go
// * ile değeri al
// & ile adresi al

x := 42
p := &x      // x'in adresi
*p = 21      // İşaretçi üzerinden değeri değiştir
fmt.Println(x)  // 21
```

### 5. Type Assertion (Tür Dönüşümü)

```go
var obj object.Object = &object.Integer{Value: 5}

// Tür kontrolü ve dönüşüm
if intObj, ok := obj.(*object.Integer); ok {
    fmt.Println(intObj.Value)  // 5
}
```

### 6. Map (Sözlük)

```go
// Anahtar-değer çiftleri
m := make(map[string]int)
m["bir"] = 1
m["iki"] = 2

value, exists := m["bir"]  // 1, true
```

### 7. Switch (Type Switch)

```go
func Eval(node ast.Node) {
    switch node := node.(type) {
    case *ast.IntegerLiteral:
        // node artık *ast.IntegerLiteral tipi
        return node.Value
    
    case *ast.StringLiteral:
        // node artık *ast.StringLiteral tipi
        return node.Value
    }
}
```

---

## 📊 Tam Akış Örneği

### Kod: `let mesaj = "Merhaba " + "Dünya";`

```
┌─────────────────────────────────────────┐
│ 1. LEXER                                │
├─────────────────────────────────────────┤
│ Input: let mesaj = "Merhaba " + "Dünya";│
│                                         │
│ Tokens:                                 │
│   LET                                   │
│   IDENT("mesaj")                        │
│   ASSIGN                                │
│   STRING("Merhaba ")                    │
│   PLUS                                  │
│   STRING("Dünya")                       │
│   SEMICOLON                             │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 2. PARSER                               │
├─────────────────────────────────────────┤
│ AST:                                    │
│   LetStatement                          │
│     Name: Identifier("mesaj")           │
│     Value: InfixExpression              │
│       Left: StringLiteral("Merhaba ")   │
│       Operator: "+"                     │
│       Right: StringLiteral("Dünya")     │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 3. EVALUATOR                            │
├─────────────────────────────────────────┤
│ 1. Eval(LetStatement)                   │
│    → Eval(InfixExpression)              │
│      → Eval(StringLiteral("Merhaba "))  │
│        → String{Value: "Merhaba "}      │
│      → Eval(StringLiteral("Dünya"))     │
│        → String{Value: "Dünya"}         │
│      → evalStringInfixExpression        │
│        → "Merhaba " + "Dünya"           │
│        → String{Value: "Merhaba Dünya"} │
│    → env.Set("mesaj", String{...})      │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 4. ENVIRONMENT                          │
├─────────────────────────────────────────┤
│ Store:                                  │
│   "mesaj" → String{Value: "Merhaba Dünya"} │
└─────────────────────────────────────────┘
```

---
🎓 Özet

## 5 Katman Hızlı Özet:

| Katman | Dosya | Görevi |
|--------|-------|--------|
| **Token** | `token/token.go` | Dilin kelime dağarcığı (LET, INT, PLUS, FUNCTION...) |
| **Lexer** | `lexer/lexer.go` | Metni token'lara ayırır |
| **AST** | `ast/` | Token'ları ağaç yapısına dönüştürür |
| **Parser** | `parser/` | Token'ları okuyup AST ağacını oluşturur |
| **Evaluator** | `evaluator/` | AST'yi dolaşıp hesaplamaları yapar |
| **Object** | `object/` | Çalışma zamanında değerler (Integer, String, Function...) |

---

## 🆕 Son Eklenen Özellikler

### 1. Türkçe Keyword'ler

| Keyword | İngilizce | Kullanım |
|---------|-----------|----------|
| `haji` | `let` | Değişken tanımlama |
| `kati` | `const` | Sabit tanımlama |

```javascript
haji x = 5      // Değiştirilebilir
kati PI = 3.14  // Değiştirilemez
```

### 2. Fonksiyonlar (`fn`)

**Dosyalar:**
- `parser/functions.go` - `parseFunctionLiteral`, `parseFunctionParameters`
- `evaluator/functions.go` - `evalFunctionLiteral`, `applyFunction`, `extendFunctionEnv`
- `object/functions.go` - `Function` struct

**Syntax:**
```javascript
haji topla = fn(a, b) {
    return a + b
}
yaz(topla(3, 5))  // 8
```

**Closure Desteği:**
```javascript
haji carpici = fn(x) {
    return fn(y) {
        return x * y
    }
}

haji ikiKati = carpici(2)
yaz(ikiKati(5))  // 10
```

**Akış:**
```
fn(x, y) { return x + y; }
         ↓
[1] Parser: parseFunctionLiteral()
         ↓
[2] AST: FunctionLiteral { Parameters, Body }
         ↓
[3] Evaluator: evalFunctionLiteral() 
         ↓
[4] Object: Function { Parameters, Body, Env }
```

### 3. For Döngüsü

**Dosyalar:**
- `parser/controls.go` - `parseForStatement`
- `evaluator/conditionals.go` - `evalForStatement`
- `ast/statements.go` - `ForStatement`

**Syntax:**
```javascript
for (haji i = 0; i < 5; i = i + 1) {
    yaz(i)
}
```

**Enclosed Environment:**
For döngüsü kendi scope'unu oluşturur. `NewEnclosedEnvironment(env)` ile dış değişkenlere erişim sağlanır.

```go
func evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
    forEnv := object.NewEnclosedEnvironment(env)  // ← Yeni scope
    
    if node.Init != nil {
        Eval(node.Init, forEnv)
    }
    
    for {
        condition := Eval(node.Condition, forEnv)
        if !isTruthy(condition) {
            break
        }
        Eval(node.Body, forEnv)
        Eval(node.Post, forEnv)
    }
    
    return object.NULL
}
```

### 4. Scope Chaining (Environment)

**Dosya:** `object/environment.go`

```go
type Environment struct {
    store     map[string]Object
    immutable map[string]bool
    outer     *Environment  // ← Dış scope referansı
}

// Get - Önce bu scope'ta, sonra dış scope'ta ara
func (e *Environment) Get(name string) (Object, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil {
        obj, ok = e.outer.Get(name)  // ← Recursive arama
    }
    return obj, ok
}

// Set - Dış scope'taki değişkeni güncelle
func (e *Environment) Set(name string, val Object) Object {
    if _, ok := e.store[name]; ok {
        e.store[name] = val
        return val
    }
    if e.outer != nil {
        if _, ok := e.outer.Get(name); ok {
            return e.outer.Set(name, val)  // ← Dış scope'ta güncelle
        }
    }
    e.store[name] = val
    return val
}
```

---

## 📁 Modüler Klasör Yapısı

```
hajiLang/
├── token/          # Token tanımları
│   └── token.go
├── lexer/          # Lexer (tokenizer)
│   └── lexer.go
├── ast/            # AST node'ları
│   ├── ast.go
│   ├── expressions.go   # Identifier, InfixExpression, FunctionLiteral...
│   ├── statements.go    # LetStatement, ForStatement, ReturnStatement...
│   └── literals.go      # IntegerLiteral, StringLiteral, ArrayLiteral...
├── parser/         # Parser
│   ├── parser.go        # Ana parser, precedence, register
│   ├── statements.go    # parseStatement, parseLetStatement...
│   ├── expressions.go   # parseExpression, parsePrefixExpression...
│   ├── controls.go      # parseIfExpression, parseForStatement, parseBlockStatement
│   ├── functions.go     # parseFunctionLiteral, parseCallExpression
│   ├── collections.go   # parseArrayLiteral, parseHashLiteral
│   └── helpers.go       # expectPeek, peekError, registerPrefix...
├── evaluator/      # Evaluator
│   ├── evaluator.go     # Ana Eval switch
│   ├── expressions.go   # evalInfixExpression, evalPrefixExpression...
│   ├── statements.go    # evalLetStatement, evalReturnStatement...
│   ├── conditionals.go  # evalIfExpression, evalForStatement
│   ├── functions.go     # evalFunctionLiteral, applyFunction, extendFunctionEnv
│   ├── literals.go      # evalIntegerLiteral, evalStringLiteral...
│   └── helpers.go       # newError, isError, isTruthy...
├── object/         # Runtime objects
│   ├── object.go        # Object interface, ObjectType
│   ├── primitives.go    # Integer, String, Boolean, Null
│   ├── functions.go     # Function, Builtin, ReturnValue
│   ├── collections.go   # Array, Hash
│   ├── environment.go   # Environment, scope chaining
│   └── builtins.go      # puts, yaz, len, first, last, push...
├── repl/           # Interactive shell
│   └── repl.go
├── runtime/        # HTTP runtime (opsiyonel)
└── main.go         # Entry point
```

---

## 🔧 Yeni Özellik Ekleme Sırası

1. **TOKEN** → `token/token.go` - Yeni token tipi ekle
2. **LEXER** → `lexer/lexer.go` - Token'ı tanı ve üret
3. **AST** → `ast/` - Yeni node yapısı ekle
4. **PARSER** → `parser/` - Parse fonksiyonu yaz, `registerPrefix/registerInfix`
5. **OBJECT** → `object/` - Çalışma zamanı tipi ekle
6. **EVALUATOR** → `evaluator/` - Değerlendirme loğiğini yaz



🏁 EOF (End of File) - Token Bitirme Sistemi
EOF Nedir?
EOF (End of File), dosyanın sonunu işaret eden özel bir token'dır. Lexer, parser ve evaluator'ın "kod bitti" dediği yerdir.
Lexer'da EOF
gofunc (l *Lexer) NextToken() token.Token {
    var tok token.Token
    
    l.skipWhitespace()
    
    switch l.ch {
    case 0:  // ← NULL karakter (ASCII 0)
        tok.Literal = ""
        tok.Type = token.EOF
        return tok
    
    // Diğer case'ler...
    }
    
    l.readChar()
    return tok
}
Neden 0 kontrolü?
gofunc (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0  // ← Dosya bitince 0 ata
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}
EOF Akışı:
Input: "let x = 5"
       ^        ^  ^
       |        |  |
     start    end EOF

Adım 1: position=0  → 'l'
Adım 2: position=4  → ' '
Adım 3: position=9  → end of string
Adım 4: position=10 → readPosition >= len(input)
                    → l.ch = 0
                    → NextToken() returns EOF
Parser'da EOF Kullanımı
gofunc (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}
    
    // EOF'a kadar döngü
    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    
    return program
}
Semicolon ve EOF
gofunc (p *Parser) parseExpression(precedence int) ast.Expression {
    prefix := p.prefixParseFns[p.curToken.Type]
    leftExp := prefix()
    
    // Semicolon VEYA EOF'ta dur
    for p.peekToken.Type != token.SEMICOLON && 
        p.peekToken.Type != token.EOF &&
        precedence < p.peekPrecedence() {
        
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        
        p.nextToken()
        leftExp = infix(leftExp)
    }
    
    return leftExp
}
Pratik Örnek:
go// test.haji
let x = 5;
let y = 10;
x + y

// Token akışı:
[LET] [IDENT:x] [ASSIGN] [INT:5] [SEMICOLON]
[LET] [IDENT:y] [ASSIGN] [INT:10] [SEMICOLON]
[IDENT:x] [PLUS] [IDENT:y]
[EOF]  ← ← ← Parser burada duruyor
EOF Olmadan Ne Olur?
go// EOF kontrolü yoksa:
for {
    stmt := p.parseStatement()  // Sonsuz döngü!
    program.Statements = append(program.Statements, stmt)
    p.nextToken()
}
// ❌ Program asla bitmez
Hata Yakalama ile EOF
gofunc (p *Parser) parseBlockStatement() *ast.BlockStatement {
    block := &ast.BlockStatement{Token: p.curToken}
    block.Statements = []ast.Statement{}
    
    p.nextToken()
    
    // '}' VEYA EOF'ta bitir (hata durumunda)
    for p.curToken.Type != token.RBRACE && 
        p.curToken.Type != token.EOF {
        
        stmt := p.parseStatement()
        if stmt != nil {
            block.Statements = append(block.Statements, stmt)
        }
        p.nextToken()
    }
    
    // Eğer EOF ile bittiyse hata!
    if p.curToken.Type == token.EOF {
        p.errors = append(p.errors, "beklenmeyen dosya sonu, '}' eksik")
    }
    
    return block
}
main.go'da EOF Kullanımı
gofunc main() {
    input := readFile("test.haji")
    
    l := lexer.New(input)
    p := parser.New(l)
    
    program := p.ParseProgram()  // EOF'a kadar parse eder
    
    if len(p.Errors()) > 0 {
        printParserErrors(p.Errors())
        return
    }
    
    env := object.NewEnvironment()
    result := evaluator.Eval(program, env)
    
    if result != nil {
        fmt.Println(result.Inspect())
    }
}

🔍 EOF ile İlgili Önemli Noktalar
1. EOF Her Yerde Kontrol Edilmeli
go// ✅ DOĞRU
for p.curToken.Type != token.EOF {
    // İşlemler
}

// ❌ YANLIŞ (sonsuz döngü riski)
for {
    // İşlemler
}
2. EOF vs Semicolon
go// Semicolon opsiyonel, EOF zorunlu
let x = 5;  // ✅ Semicolon var
let x = 5   // ✅ Semicolon yok ama EOF gelecek, parser bunu kabul eder
3. Hata Mesajlarında EOF
goif p.curToken.Type == token.EOF {
    return newError("beklenmeyen dosya sonu")
}

// Örnek hata:
// if (x > 5) {
//   10
// [EOF]  ← "'}' eksik" hatası verilir
4. REPL (Interactive Shell) için EOF
go// REPL modunda her satır ayrı parse edilir
for {
    fmt.Print(">> ")
    line := readLine()
    
    if line == "exit" {
        break
    }
    
    l := lexer.New(line)
    p := parser.New(l)
    program := p.ParseProgram()
    
    // Her satır kendi EOF'una sahip
    result := evaluator.Eval(program, env)
    fmt.Println(result.Inspect())
}

📚 Sonuç
EOF, interpreter'ın "dur" sinyalidir. Olmadan parser sonsuz döngüye girer.
Token → Lexer → Parser → Evaluator hattında EOF her katmanda kontrol edilir:

Lexer: Dosya bitince 0 karakteri görür, EOF token'ı üretir
Parser: EOF görünce program parse'ını bitirir
Evaluator: Parser EOF'ta durduğu için zaten tüm AST'yi alır