package object

import (
	"bytes"
	"fmt"
	"strings"
)

// Array - Dizi objesi
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

// HashKey - Hash anahtarı için struct
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashPair - Hash key-value çifti
type HashPair struct {
	Key   Object
	Value Object
}

// Hash - Hash map objesi
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

// Hashable - Hash key olarak kullanılabilecek tiplerin interface'i
type Hashable interface {
	HashKey() HashKey
}