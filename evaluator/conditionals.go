package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// evalIfExpression - If expression'ını evaluate eder
// Örnek: if x > 5 { return true; } else { return false; }
func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return object.NULL
	}
}

// isTruthy - Bir değerin truthy olup olmadığını kontrol eder
// false ve nil dışında her şey truthy'dir
func isTruthy(obj object.Object) bool {
	switch obj {
	case nil:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
func evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
	// Yeni bir scope oluştur (döngü değişkeni için)
	forEnv := object.NewEnclosedEnvironment(env)

	// 1. Init çalıştır (i = 0)
	if node.Init != nil {
		Eval(node.Init, forEnv)
	}

	// 2. Döngü
	for {
		// Condition kontrolü (i < 10)
		if node.Condition != nil {
			condition := Eval(node.Condition, forEnv)
			if isError(condition) {
				return condition
			}

			if !isTruthy(condition) {
				break
			}
		}

		// Body çalıştır
		result := Eval(node.Body, forEnv)
		if isError(result) {
			return result
		}

		// Post çalıştır (i = i + 1)
		if node.Post != nil {
			Eval(node.Post, forEnv)
		}
	}

	return object.NULL
}
