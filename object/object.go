package object

import "fmt"

// Değişken Tipleri
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	ERROR_OBJ   = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string // Ekrana yazılacak hali
}

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

// --- BOOLEAN EKLEME ---
const BOOLEAN_OBJ = "BOOLEAN"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
