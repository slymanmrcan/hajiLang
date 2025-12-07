package object

// ObjectType - Object tiplerini tanımlayan string sabitleri
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	NULL_OBJ         = "NULL"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	BUILTIN_OBJ      = "BUILTIN"
	FUNCTION_OBJ     = "FUNCTION" // Gelecek için hazır
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FLOAT_OBJ        = "FLOAT"
)

// Object - Tüm object tiplerinin implement etmesi gereken ana interface
type Object interface {
	Type() ObjectType
	Inspect() string // Ekrana yazılacak/debug için string hali
}
