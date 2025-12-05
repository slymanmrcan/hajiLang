package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

// Değişken Tipleri
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	ERROR_OBJ        = "ERROR"
	STRING_OBJ       = "STRING"
	BOOLEAN_OBJ      = "BOOLEAN"
	BUILTIN_OBJ      = "BUILTIN"
	NULL_OBJ         = "NULL"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string // Ekrana yazılacak hali
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

var NULL = &Null{}

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Builtin fonksiyon tipi tanımı
type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

// Builtin için Object arayüzü metodları
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// --- TAMSAYI (INTEGER) ---
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// --- HATA MESAJI ---
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "HATA: " + e.Message }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type String struct {
	Value string
}

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}

// String ve Integer nesnelerinin Hash anahtarı olarak kullanılabilmesi için:

// (object.go içindeki String struct'ının altına bu metodu ekle)
func (s *String) HashKey() HashKey {
	h := fnv.New64a() // "hash/fnv" kütüphanesini import etmen gerek!
	h.Write([]byte(s.Value))
	return HashKey{Type: STRING_OBJ, Value: h.Sum64()}
}

// (object.go içindeki Integer struct'ının altına bu metodu ekle)
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: INTEGER_OBJ, Value: uint64(i.Value)}
}

// GoToHash: JSON'dan (Go interface{}) gelen veriyi HajiLang nesnesine çevirir.
func GoToHash(val interface{}) Object {
	switch v := val.(type) {
	case string:
		return &String{Value: v}
	case float64:
		// JSON'da sayılar float gelir, biz Integer'a çeviriyoruz
		return &Integer{Value: int64(v)}
	case int:
		return &Integer{Value: int64(v)}
	case bool:
		// Eğer Boolean yapın varsa onu dön, yoksa NULL dön
		if v {
			return TRUE
		}
		return FALSE
	case map[string]interface{}:
		// JSON Objesini (Map) -> Haji Hash'ine çevir
		pairs := make(map[HashKey]HashPair)
		for k, val := range v {
			key := &String{Value: k}
			value := GoToHash(val) // İç içe (Recursive)
			pairs[key.HashKey()] = HashPair{Key: key, Value: value}
		}
		return &Hash{Pairs: pairs}
	case []interface{}:
		// JSON Listesini (Array) -> Haji Array'ine çevir
		elements := []Object{}
		for _, i := range v {
			elements = append(elements, GoToHash(i))
		}
		return &Array{Elements: elements}
	case nil:
		return NULL
	}
	return NULL
}
