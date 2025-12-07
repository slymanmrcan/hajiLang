package object

import (
	"fmt"
	"hash/fnv"
)

// Integer - Tamsayı objesi
type Integer struct {
	Value int64
}

type Float struct {
	Value float64
}

func (f *Float) Type() ObjectType { return FLOAT_OBJ }
func (f *Float) Inspect() string  { return fmt.Sprintf("%g", f.Value) }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Integer'ı hash key olarak kullanabilmek için
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: INTEGER_OBJ, Value: uint64(i.Value)}
}

// Boolean - Boolean objesi
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Boolean'ı hash key olarak kullanabilmek için
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: BOOLEAN_OBJ, Value: value}
}

// Global Boolean singleton'ları
var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

// String - String objesi
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// String'i hash key olarak kullanabilmek için
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: STRING_OBJ, Value: h.Sum64()}
}

// Null - Null objesi (değer yok)
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// Global Null singleton
var NULL = &Null{}

func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%f", f.Value)))
	return HashKey{Type: FLOAT_OBJ, Value: h.Sum64()}
}
