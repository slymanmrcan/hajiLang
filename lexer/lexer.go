package lexer

import "github.com/slymanmrcan/hajilang/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int // Satır numarası
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Bir sonraki karaktere bak (İlerletmeden)
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		// Eğer bir sonraki de '=' ise, bu '==' demektir
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '!':
		// Eğer bir sonraki '=' ise, bu '!=' demektir
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch) //
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch) //
	case ']':
		tok = newToken(token.RBRACKET, l.ch) //
	case '/':
		// Eğer bir sonraki de '/' ise bu bir yorumdur!
		if l.peekChar() == '/' {
			// Satır sonuna (\n) kadar her şeyi atla
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}
			// Yorum bitti, boşlukları atla ve sıradaki GERÇEK tokenı bul
			l.skipWhitespace()
			return l.NextToken()
		} else {
			// Değilse bu normal bölme işlemidir
			tok = newToken(token.SLASH, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '%':
		tok = newToken(token.PERCENT, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			numStr, isFloat := l.readNumber()
			if isFloat {
				tok.Type = token.FLOAT
				tok.Literal = numStr
			} else {
				tok.Type = token.INT
				tok.Literal = numStr
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	tok.Line = l.line
	l.readChar()
	return tok
}

func (l *Lexer) newTokenWithLine(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: l.line}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()

		// Eğer \ (ters taksim) görürsek, bu bir kaçış karakteridir.
		// Bir sonraki karakteri de stringin parçası olarak kabul et ve döngüye devam et.
		// Böylece string " işaretini görünce bitmez.
		if l.ch == '\\' {
			l.readChar() // \ işaretinden sonrakini oku (örn: " işaretini)
			continue
		}

		// Stringin sonuna geldik mi? (EOF veya " işareti)
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
func (l *Lexer) readNumber() (string, bool) {
	position := l.position
	isFloat := false

	for isDigit(l.ch) {
		l.readChar()
	}

	// Eğer nokta görürsek ve sonrası rakamsa, bu float'tır
	if l.ch == '.' && isDigit(l.peekChar()) {
		isFloat = true
		l.readChar() // '.' geç

		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position], isFloat
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
