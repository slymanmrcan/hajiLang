package evaluator

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/object"
)

// Global boolean singleton'ları
var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

// Eval - Ana evaluation fonksiyonu, AST node'larını evaluate eder
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Program
	case *ast.Program:
		return evalProgram(node, env)

	// Statements
	case *ast.ExpressionStatement:
		result := Eval(node.Expression, env)
		return result

	case *ast.LetStatement:
		return evalLetStatement(node, env)
	case *ast.HajiStatement: // ← YENİ!
		return evalHajiStatement(node, env)

	case *ast.KatiStatement: // ← YENİ!
		return evalKatiStatement(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ReturnStatement: // ← return !
		return evalReturnStatement(node, env)
	case *ast.ForStatement: // ← YENİ!
		return evalForStatement(node, env)
	// Expressions
	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)

	case *ast.InfixExpression:
		return evalInfixExpression(node, env)

	case *ast.CallExpression:
		return evalCallExpression(node, env)

	case *ast.IndexExpression:
		return evalIndexExpression(node, env)

	// Literals
	case *ast.IntegerLiteral:
		return evalIntegerLiteral(node)
	case *ast.FloatLiteral: // ← YENİ!
		return evalFloatLiteral(node)
	case *ast.Boolean:
		return evalBoolean(node)

	case *ast.StringLiteral:
		return evalStringLiteral(node)

	case *ast.ArrayLiteral:
		return evalArrayLiteral(node, env)

	case *ast.HashLiteral:
		return evalHashLiteral(node, env)

	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)
	}

	return nil
}

// evalProgram - Program node'unu evaluate eder
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		// Sadece RETURN veya ERROR varsa dur
		if result != nil {
			switch result.(type) {
			case *object.ReturnValue:
				return result.(*object.ReturnValue).Value
			case *object.Error:
				return result
			}
		}
		// Diğer durumlarda (NULL, Integer, vb.) devam et
	}

	return result
}
