package evaluator

import (
	"fmt"
	"hajilang/ast"
	"hajilang/object"
)

// --- EKSİK OLAN KISIM 1: GLOBAL TRUE/FALSE ---
var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// 1. Program
	case *ast.Program:
		return evalProgram(node, env)

	// 2. Expression Statement
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	// 3. Let Statement
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return nil

	// 4. Identifier (Değişken)
	case *ast.Identifier:
		return evalIdentifier(node, env)

	// --- EKSİK OLAN KISIM 2: IF VE BLOKLAR ---
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	// 5. Tamsayılar
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	// --- EKSİK OLAN KISIM 3: BOOLEAN ---
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	// 6. Prefix (-5, !true)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	// 7. Infix (5 + 5, 10 < 20)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		if result != nil && result.Type() == object.ERROR_OBJ {
			return result
		}
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil && result.Type() == object.ERROR_OBJ {
			return result
		}
	}
	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("tanımlanmamış değişken: " + node.Value)
	}
	return val
}

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
		return nil
	}
}

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

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("bilinmeyen operatör: %s%s", operator, right.Type())
	}
}

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

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("bilinmeyen operatör: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left, right)
	}
	// Boolean Karşılaştırma (==, !=)
	if operator == "==" {
		return nativeBoolToBooleanObject(left == right)
	}
	if operator == "!=" {
		return nativeBoolToBooleanObject(left != right)
	}
	if left.Type() != right.Type() {
		return newError("tür uyuşmazlığı: %s %s %s", left.Type(), operator, right.Type())
	}
	return newError("bilinmeyen operatör: %s %s %s", left.Type(), operator, right.Type())
}

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
	// --- EKSİK OLAN KISIM 4: KÜÇÜKTÜR/BÜYÜKTÜR ---
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

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
