package ast

import (
	"bytes"

	"github.com/slymanmrcan/hajilang/token"
)

// LetStatement - Let declaration statement
// Örnek: let x = 5;
type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}
type HajiStatement struct {
	Token token.Token // HAJI token
	Name  *Identifier
	Value Expression
}
type KatiStatement struct {
	Token token.Token // KATI token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " " + ls.Name.String() + " = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement - Expression statement
// Örnek: 5 + 5;
type ExpressionStatement struct {
	Token      token.Token // İlk token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement - Block statement
// Örnek: { let x = 5; return x; }
type BlockStatement struct {
	Token      token.Token // '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// ReturnStatement - Return statement (gelecek için hazır)
// Örnek: return 5;
type ReturnStatement struct {
	Token       token.Token // RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ForStatement - For loop statement (gelecek için hazır)
// Örnek: for (let i = 0; i < 10; i = i + 1) { ... }
type ForStatement struct {
	Token     token.Token // FOR token
	Init      Statement   // let i = 0
	Condition Expression  // i < 10
	Update    Statement   // i = i + 1
	Body      *BlockStatement
	Post      Statement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString("for (")
	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	}
	out.WriteString("; ")
	if fs.Condition != nil {
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")
	if fs.Update != nil {
		out.WriteString(fs.Update.String())
	}
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	return out.String()
}
func (hs *HajiStatement) statementNode()       {}
func (hs *HajiStatement) TokenLiteral() string { return hs.Token.Literal }
func (hs *HajiStatement) String() string {
	var out bytes.Buffer
	out.WriteString(hs.TokenLiteral() + " " + hs.Name.String() + " = ")
	if hs.Value != nil {
		out.WriteString(hs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
func (ks *KatiStatement) statementNode()       {}
func (ks *KatiStatement) TokenLiteral() string { return ks.Token.Literal }
func (ks *KatiStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ks.TokenLiteral() + " " + ks.Name.String() + " = ")
	if ks.Value != nil {
		out.WriteString(ks.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
