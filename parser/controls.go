package parser

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/token"
)

// parseIfExpression - If/else if/else expression'ını parse eder
// Örnek: if x > 5 { return true; } else { return false; }
// Örnek: if path == "/" { ... } else if path == "/about" { ... }
func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	// if'ten sonra koşul - PARANTEZ OLMADAN da çalışır
	p.nextToken()

	// Eğer '(' varsa, parantezli syntax
	if p.curToken.Type == token.LPAREN {
		p.nextToken()
		expression.Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		// Parantezsiz syntax: if path == "/" {
		expression.Condition = p.parseExpression(LOWEST)
	}

	// { bekliyoruz
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// Consequence block'unu parse et
	expression.Consequence = p.parseBlockStatement()

	// ELSE var mı kontrol et
	if p.peekToken.Type == token.ELSE {
		p.nextToken() // else'i geç

		// else if mi kontrol et
		if p.peekToken.Type == token.IF {
			p.nextToken() // if'e geç
			// else if'i yeni bir IfExpression olarak parse et
			elseIfExp := p.parseIfExpression()
			// Bunu bir block içine koy
			expression.Alternative = &ast.BlockStatement{
				Token: p.curToken,
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      p.curToken,
						Expression: elseIfExp,
					},
				},
			}
		} else if p.peekToken.Type == token.LBRACE {
			// Normal else bloğu
			p.nextToken()
			expression.Alternative = p.parseBlockStatement()
		}
	}

	return expression
}

// parseBlockStatement - Block statement'ı parse eder
// Örnek: { let x = 5; return x; }
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	// } veya EOF'a kadar statement'ları parse et
	for p.curToken.Type != token.RBRACE && p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()

	// Init statement (haji i = 0 veya i = 0)
	if p.curToken.Type == token.HAJI {
		stmt.Init = p.parseHajiStatement()
	} else if p.curToken.Type == token.KATI {
		stmt.Init = p.parseKatiStatement()
	} else if p.curToken.Type == token.LET {
		stmt.Init = p.parseLetStatement()
	} else if p.curToken.Type == token.IDENT && p.peekToken.Type == token.ASSIGN {
		stmt.Init = p.parseAssignmentStatement()
	} else {
		stmt.Init = p.parseStatement()
	}

	// Semicolon sonrası (statement'lar zaten semicolon'u consume ediyor olabilir)
	// curToken şu an semicolon ise ilerle, değilse semicolon bekle
	if p.curToken.Type != token.SEMICOLON {
		if !p.expectPeek(token.SEMICOLON) {
			return nil
		}
	}

	p.nextToken()

	// Condition (i < 10)
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	p.nextToken()

	// Post statement (i = i + 1)
	if p.curToken.Type == token.IDENT && p.peekToken.Type == token.ASSIGN {
		stmt.Post = p.parseAssignmentStatement()
	} else {
		stmt.Post = p.parseStatement()
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}
