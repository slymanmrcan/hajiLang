package object

import (
	"fmt"
	"strconv"
)

// Bütün gömülü fonksiyonları burada tanımlıyoruz
var Builtins = map[string]*Builtin{
	// --- GENEL FONKSİYONLAR ---

	// len(nesne) -> Nesnenin uzunluğunu döner
	"len": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},

	// puts(değer) -> Ekrana yazdırır (Console log gibi)
	"puts": &Builtin{
		Fn: func(args ...Object) Object {
			for _, arg := range args {
				// İŞTE BU SATIR EKRANA YAZAR
				// Eğer arg.Inspect() string tırnaklarıyla geliyorsa direkt Value yazdırabiliriz
				if str, ok := arg.(*String); ok {
					fmt.Println(str.Value)
				} else {
					fmt.Println(arg.Inspect())
				}
			}
			// puts bir değer dönmez, işi biter. NULL döner.
			return NULL
		},
	},

	// --- DİZİ (ARRAY) İŞLEMLERİ ---

	// first(dizi) -> İlk elemanı döner
	"first": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},

	// last(dizi) -> Son elemanı döner
	"last": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},

	// rest(dizi) -> İlk eleman hariç kalanı döner (cdr gibi)
	"rest": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &Array{Elements: newElements}
			}
			return NULL
		},
	},

	// push(dizi, eleman) -> Diziye eleman ekler ve yeni diziyi döner
	"push": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			newElements := make([]Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &Array{Elements: newElements}
		},
	},

	// --- TÜR DÖNÜŞÜMLERİ (API İÇİN KRİTİK) ---

	// to_int("5") -> 5 (String'i sayıya çevirir)
	"to_int": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *String:
				val, err := strconv.ParseInt(arg.Value, 0, 64)
				if err != nil {
					return newError("could not convert string to int")
				}
				return &Integer{Value: val}
			case *Integer:
				return arg
			default:
				return newError("argument to `to_int` not supported, got %s", args[0].Type())
			}
		},
	},

	// to_str(5) -> "5" (Sayıyı stringe çevirir)
	"to_str": &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &String{Value: args[0].Inspect()}
		},
	},
}

// Yardımcı hata fonksiyonu
func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
