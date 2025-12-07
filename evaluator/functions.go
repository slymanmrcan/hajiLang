package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// evalCallExpression - Function call expression'ını evaluate eder
// Örnek: add(1, 2), puts("hello")
func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if isError(function) {
		return function
	}

	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return applyFunction(function, args)
}

// applyFunction - Fonksiyonu argümanlarla çağırır
func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Builtin:
		// Builtin fonksiyon (puts, len, first vb.)
		return fn.Fn(args...)

	case *object.Function:
		// User-defined fonksiyon
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

// extendFunctionEnv - Fonksiyon için yeni environment oluşturur
// Parametreleri argümanlarla bağlar
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param.Value, args[paramIdx])
		}
	}

	return env
}

// unwrapReturnValue - ReturnValue'yu unwrap eder
// Fonksiyon içindeki return değerini çıkarır
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

// evalFunctionLiteral - Function literal'ı evaluate eder
// fn(x, y) { return x + y; } -> Function object
func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	return &object.Function{
		Parameters: node.Parameters,
		Body:       node.Body,
		Env:        env,
	}
}
