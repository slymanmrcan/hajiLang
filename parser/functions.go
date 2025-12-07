package parser

import (
	"github.com/slymanmrcan/hajilang/ast"
	"github.com/slymanmrcan/hajilang/token"
)

// parseCallExpression - Function call expression'ını parse eder
// Örnek: add(1, 2), myFunc(), nested(foo(x))
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

// parseFunctionLiteral - Function literal'ı parse eder
// Örnek: fn(x, y) { return x + y; }
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// ( bekliyoruz
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parametreleri parse et
	lit.Parameters = p.parseFunctionParameters()

	// { bekliyoruz
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// Body'yi parse et
	lit.Body = p.parseBlockStatement()

	return lit
}

// parseFunctionParameters - Fonksiyon parametrelerini parse eder
// Örnek: (x, y, z)
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	// Boş parametre listesi: fn() { ... }
	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	// İlk parametre
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	// Virgülle ayrılmış diğer parametreler
	for p.peekToken.Type == token.COMMA {
		p.nextToken() // virgülü geç
		p.nextToken() // sonraki parametreye geç
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	// ) bekliyoruz
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}
