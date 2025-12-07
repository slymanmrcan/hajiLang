package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// evalIntegerLiteral - Integer literal'i evaluate eder
// Örnek: 5, 42, 1000
func evalIntegerLiteral(node *ast.IntegerLiteral) object.Object {
	return &object.Integer{Value: node.Value}
}

// evalBoolean - Boolean literal'i evaluate eder
// Örnek: true, false
func evalBoolean(node *ast.Boolean) object.Object {
	return nativeBoolToBooleanObject(node.Value)
}

// evalStringLiteral - String literal'i evaluate eder
// Örnek: "hello", "world"
func evalStringLiteral(node *ast.StringLiteral) object.Object {
	return &object.String{Value: node.Value}
}

// evalArrayLiteral - Array literal'i evaluate eder
// Örnek: [1, 2, 3], ["a", "b"]
func evalArrayLiteral(node *ast.ArrayLiteral, env *object.Environment) object.Object {
	elements := evalExpressions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}

// evalHashLiteral - Hash literal'i evaluate eder
// Örnek: {"name": "John", "age": 30}
func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		// Anahtar, hash oluşturulabilir bir tip mi? (String, Int, Bool)
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

// evalExpressions - Expression listesini evaluate eder
// Function arguments ve array elements için kullanılır
func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalFloatLiteral(node *ast.FloatLiteral) object.Object {
	return &object.Float{Value: node.Value}
}
