package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// evalIdentifier - Identifier'ı evaluate eder
// Önce environment'ta arar, sonra builtin fonksiyonlara bakar
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	// 1. Environment'ta değişken var mı?
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	// 2. Builtin fonksiyon mu? (puts, len, first vb.)
	if builtin, ok := object.Builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

// evalPrefixExpression - Prefix expression'ı evaluate eder
// Örnek: -5, !true
func evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) object.Object {
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}

	switch node.Operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("bilinmeyen operatör: %s%s", node.Operator, right.Type())
	}
}

// evalBangOperatorExpression - ! operatörünü evaluate eder
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case nil:
		return TRUE
	default:
		return FALSE
	}
}

// evalMinusPrefixOperatorExpression - - operatörünü evaluate eder
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("bilinmeyen operatör: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// evalInfixExpression - Infix expression'ı evaluate eder
// Örnek: 5 + 5, "hello" + "world", x == y
func evalInfixExpression(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}

	// STRING İŞLEMLERİ
	if left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ {
		return evalStringInfixExpression(node.Operator, left, right)
	}

	// INTEGER İŞLEMLERİ
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(node.Operator, left, right)
	}
	if left.Type() == object.FLOAT_OBJ || right.Type() == object.FLOAT_OBJ {
		return evalFloatInfixExpression(node.Operator, left, right, env)
	}
	// BOOLEAN KARŞILAŞTIRMA
	if node.Operator == "==" {
		return nativeBoolToBooleanObject(left == right)
	}
	if node.Operator == "!=" {
		return nativeBoolToBooleanObject(left != right)
	}

	// MANTIKSAL OPERATÖRLER (&& ve ||)
	if node.Operator == "&&" {
		return nativeBoolToBooleanObject(isTruthy(left) && isTruthy(right))
	}
	if node.Operator == "||" {
		return nativeBoolToBooleanObject(isTruthy(left) || isTruthy(right))
	}

	// TÜR UYUŞMAZLIĞI
	if left.Type() != right.Type() {
		return newError("tür uyuşmazlığı: %s %s %s", left.Type(), node.Operator, right.Type())
	}

	return newError("bilinmeyen operatör: %s %s %s", left.Type(), node.Operator, right.Type())
}

// evalStringInfixExpression - String operasyonları
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		// String birleştirme
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("bilinmeyen operatör: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evalIntegerInfixExpression - Integer operasyonları
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("sıfıra bölünemez!")
		}
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		if rightVal == 0 {
			return newError("sıfıra mod alınamaz!")
		}
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("bilinmeyen operatör: %s %s %s", left.Type(), operator, right.Type())
	}
}

// evalFloatInfixExpression - Float operasyonları (mixed type da destekler)
func evalFloatInfixExpression(operator string, left, right object.Object, env *object.Environment) object.Object {
	var leftVal, rightVal float64

	// Sol tarafı float'a çevir
	switch l := left.(type) {
	case *object.Float:
		leftVal = l.Value
	case *object.Integer:
		leftVal = float64(l.Value)
	default:
		return newError("bilinmeyen operatör: %s %s %s", left.Type(), operator, right.Type())
	}

	// Sağ tarafı float'a çevir
	switch r := right.(type) {
	case *object.Float:
		rightVal = r.Value
	case *object.Integer:
		rightVal = float64(r.Value)
	default:
		return newError("bilinmeyen operatör: %s %s %s", left.Type(), operator, right.Type())
	}

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("sıfıra bölünemez!")
		}
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("bilinmeyen operatör: %s %s %s", left.Type(), operator, right.Type())
	}
}
