package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// evalIndexExpression - Index expression'ını evaluate eder
// Örnek: arr[0], hash["key"], matrix[i][j]
func evalIndexExpression(node *ast.IndexExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}
	
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

// evalArrayIndexExpression - Array index işlemini evaluate eder
// Örnek: [1, 2, 3][0] -> 1
func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	// Index sınırları dışında mı?
	if idx < 0 || idx > max {
		return object.NULL // Hata vermek yerine NULL dön
	}

	return arrayObject.Elements[idx]
}

// evalHashIndexExpression - Hash index işlemini evaluate eder
// Örnek: {"name": "John"}["name"] -> "John"
func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	// Index, hash key olarak kullanılabilir mi?
	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}

	// Anahtar hash'te var mı?
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return object.NULL // Anahtar yoksa NULL dön
	}

	return pair.Value
}