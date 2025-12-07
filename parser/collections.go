package parser

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/token"
)

// parseArrayLiteral - Array literal'i parse eder
// Örnek: [1, 2, 3], ["a", "b"], [true, false]
func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

// parseHashLiteral - Hash (map) literal'i parse eder
// Örnek: {"name": "John", "age": 30}
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	// } görene kadar key:value çiftlerini parse et
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		// : bekle
		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		hash.Pairs[key] = value

		// } değilse , bekle
		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	// Kapanış } bekle
	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return hash
}

// parseIndexExpression - Index expression'ı parse eder
// Örnek: arr[0], hash["key"], matrix[i][j]
func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	// ] bekle
	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}

// parseExpressionList - Virgülle ayrılmış expression listesi parse eder
// Array ve function arguments için kullanılır
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	// Eğer hemen closing token varsa boş liste dön
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	// İlk element'i parse et
	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	// , görüldükçe devam et
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // , 'yi geç
		p.nextToken() // sonraki expression'a geç
		list = append(list, p.parseExpression(LOWEST))
	}

	// Closing token'ı bekle (] veya ))
	if !p.expectPeek(end) {
		return nil
	}

	return list
}
