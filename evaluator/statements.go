package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// evalLetStatement - Let statement'ını evaluate eder
// Örnek: let x = 5;
func evalLetStatement(node *ast.LetStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	env.Set(node.Name.Value, val)
	return nil
}

// evalBlockStatement - Block statement'ı evaluate eder
// Örnek: { let x = 5; return x; }
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		// SADECE ERROR veya RETURN durumunda dur
		if result != nil {
			switch result.(type) {
			case *object.ReturnValue:
				return result
			case *object.Error:
				return result
			}
		}
		// NULL veya başka değerlerde devam et
	}

	return result
}

// evalReturnStatement - Return statement'ını evaluate eder
// Örnek: return 5;
func evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) object.Object {
	val := Eval(node.ReturnValue, env)
	if isError(val) {
		return val
	}
	return &object.ReturnValue{Value: val}
}

// evalHajiStatement - haji statement'ını evaluate eder
func evalHajiStatement(node *ast.HajiStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	// Eğer bu değişken daha önce kati olarak tanımlandıysa HATA!
	if env.IsConst(node.Name.Value) {
		return newError("'%s' sabittir, değiştirilemez!", node.Name.Value)
	}

	env.Set(node.Name.Value, val)
	return nil
}

// evalKatiStatement - kati statement'ını evaluate eder
func evalKatiStatement(node *ast.KatiStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	// Eğer bu isim zaten kullanılıyorsa HATA!
	if _, exists := env.Get(node.Name.Value); exists {
		return newError("'%s' zaten tanımlanmış!", node.Name.Value)
	}

	env.SetConst(node.Name.Value, val)
	return nil
}
