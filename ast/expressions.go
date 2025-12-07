package ast

import (
	"bytes"
	"strings"

	"github.com/slymanmrcan/hajilang/token"
)

// Identifier - Değişken veya fonksiyon isimleri
// Örnek: x, myVar, add
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// PrefixExpression - Prefix operatörlü ifadeler
// Örnek: -5, !true
type PrefixExpression struct {
	Token    token.Token // !, -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression - Infix operatörlü ifadeler
// Örnek: 5 + 5, 10 * 2, x == y
type InfixExpression struct {
	Token    token.Token // +, -, *, /, ==, !=, <, >
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

// IfExpression - If/else ifadesi
// Örnek: if (x < y) { ... } else { ... }
type IfExpression struct {
	Token       token.Token // 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement // else kısmı (opsiyonel)
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// CallExpression - Fonksiyon çağrısı
// Örnek: add(1, 2), puts("hello")
type CallExpression struct {
	Token     token.Token // '(' token
	Function  Expression  // Identifier veya FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	return ce.Function.String() + "(" + strings.Join(args, ", ") + ")"
}

// IndexExpression - Index erişimi
// Örnek: array[0], hash["key"]
type IndexExpression struct {
	Token token.Token // '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

// FunctionLiteral - Fonksiyon tanımlama (gelecek için hazır)
// Örnek: fn(x, y) { return x + y; }
type FunctionLiteral struct {
	Token      token.Token // 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}