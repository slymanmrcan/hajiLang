package runtime

import (
	"strconv"

	"github.com/slymanmrcan/hajilang/object"
)

func RegisterUtils(env *object.Environment) {
	// String to Integer (ID karşılaştırması için şart)
	env.Set("to_int", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "wrong number of args"}
			}
			str, ok := args[0].(*object.String)
			if !ok {
				return &object.Error{Message: "argument must be string"}
			}

			val, _ := strconv.Atoi(str.Value)
			return &object.Integer{Value: int64(val)}
		},
	})

	// Integer to String
	env.Set("to_str", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Argüman verilmemişse boş string dön
			if len(args) == 0 {
				return &object.String{Value: ""}
			}
			// Her şeyi stringe çevir (Inspect metodu sayesinde)
			return &object.String{Value: args[0].Inspect()}
		},
	})

	// Array Push (Listeye ekleme yapmak için)
	env.Set("push", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "wrong number of args"}
			}
			arr, ok := args[0].(*object.Array)
			if !ok {
				return &object.Error{Message: "first argument must be array"}
			}

			// Orijinal diziyi bozmamak için kopyasını oluşturup ekliyoruz (Immutability)
			newElements := make([]object.Object, len(arr.Elements), len(arr.Elements)+1)
			copy(newElements, arr.Elements)
			newElements = append(newElements, args[1])

			return &object.Array{Elements: newElements}
		},
	})
}
