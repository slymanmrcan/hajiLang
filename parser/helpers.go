package parser

import (
	"fmt"

	"github.com/slymanmrcan/hajilang/token"
)

// registerPrefix - Prefix parse fonksiyonunu kayıt eder
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix - Infix parse fonksiyonunu kayıt eder
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// expectPeek - Bir sonraki token beklenen tipte mi kontrol eder
// Eğer öyleyse token'ı ilerletir ve true döner
// Değilse hata ekler ve false döner
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// Errors - Parser hatalarını döndürür
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError - Beklenen token gelmediğinde hata ekler
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Satır %d: Beklenen token %s, ama gelen %s", p.curToken.Line, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// noPrefixParseFnError - Prefix parse fonksiyonu bulunamadığında hata ekler
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("Satır %d: '%s' (Tip: %s) ile cümleye başlanamaz.", p.curToken.Line, p.curToken.Literal, t)
	p.errors = append(p.errors, msg)
}
