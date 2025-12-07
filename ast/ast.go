package ast

import (
	"bytes"
)

// Node - Tüm AST node'larının implement etmesi gereken ana interface
type Node interface {
	TokenLiteral() string
	String() string // Hata ayıklama için yazdırılabilir hali
}

// Statement - Statement node'larının interface'i
type Statement interface {
	Node
	statementNode()
}

// Expression - Expression node'larının interface'i
type Expression interface {
	Node
	expressionNode()
}

// Program - Programın tamamını temsil eden root node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
