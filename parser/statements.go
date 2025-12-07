package parser

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/token"
)

// parseStatement - Token tipine göre uygun statement'ı parse eder
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.HAJI:
		return p.parseHajiStatement()
	case token.KATI:
		return p.parseKatiStatement()
	case token.RETURN: // ← YENİ EKLEME!
		return p.parseReturnStatement()
	case token.FOR:
		return p.parseForStatement()
	case token.IDENT:
		// Assignment check: x = 5
		if p.peekToken.Type == token.ASSIGN {
			return p.parseAssignmentStatement()
		}
		return p.parseExpressionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseHajiStatement - haji statement'ını parse eder
func (p *Parser) parseHajiStatement() *ast.HajiStatement {
	stmt := &ast.HajiStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// parseKatiStatement - kati statement'ını parse eder
func (p *Parser) parseKatiStatement() *ast.KatiStatement {
	stmt := &ast.KatiStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// parseLetStatement - Let statement'ını parse eder
// Örnek: let x = 5;
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Değişken adını bekle
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// = işaretini bekle
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	// Değeri parse et
	stmt.Value = p.parseExpression(LOWEST)

	// Opsiyonel semicolon
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// parseAssignmentStatement - Assignment statement'ını parse eder
// Örnek: x = 10;
func (p *Parser) parseAssignmentStatement() *ast.LetStatement {
	// Assignment'ı LetStatement olarak kullanıyoruz
	stmt := &ast.LetStatement{Token: token.Token{Type: token.LET, Literal: "let"}}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// = işaretini bekle
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	// Yeni değeri parse et
	stmt.Value = p.parseExpression(LOWEST)

	// Opsiyonel semicolon
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement - Expression statement'ını parse eder
// Örnek: 5 + 5; veya fonksiyon();
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	// Opsiyonel semicolon
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement - Return statement'ını parse eder
// Örnek: return 5; return x + y;
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// Return değerini parse et
	stmt.ReturnValue = p.parseExpression(LOWEST)

	// Opsiyonel semicolon
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}
