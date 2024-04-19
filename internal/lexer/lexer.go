package lexer

import (
	"github.com/Laellekoenig/monke/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peakChar() byte {
  if l.readPosition >= len(l.input) {
    return 0
  }

  return l.input[l.readPosition]
}

func createToken(t token.TokenType, l byte) *token.Token {
	return &token.Token{
		Type:    t,
		Literal: string(l),
	}
}

func (l *Lexer) NextToken() *token.Token {
	var t *token.Token

	l.eatWhitespace()

	switch l.char {
	case '=':
    if l.peakChar() == '=' {
      t = &token.Token{Type: token.EQ, Literal: "=="}
      l.readChar()
    } else {
		  t = createToken(token.ASSIGN, l.char)
    }
	case '+':
		t = createToken(token.PLUS, l.char)
	case ',':
		t = createToken(token.COMMA, l.char)
	case ';':
		t = createToken(token.SEMICOLON, l.char)
	case '(':
		t = createToken(token.LPAREN, l.char)
	case ')':
		t = createToken(token.RPAREN, l.char)
	case '{':
		t = createToken(token.LBRACE, l.char)
	case '}':
		t = createToken(token.RBRACE, l.char)
	case '!':
    if l.peakChar() == '=' {
      t = &token.Token{Type: token.NEQ, Literal: "!="}
      l.readChar()
    } else {
		  t = createToken(token.BANG, l.char)
    }
	case '-':
		t = createToken(token.MINUS, l.char)
	case '/':
		t = createToken(token.SLASH, l.char)
	case '*':
		t = createToken(token.ASTERIX, l.char)
	case '<':
		t = createToken(token.LT, l.char)
	case '>':
		t = createToken(token.GT, l.char)
	case 0:
		t = createToken(token.EOF, l.char)
	default:
		if isLetter(l.char) {
			literal := l.readWhile(isLetter)
			tokenType := token.LookupKeyword(&literal)
			return &token.Token{Type: tokenType, Literal: literal}

		} else if isNumber(l.char) {
			number := l.readWhile(isNumber)
			return &token.Token{Type: token.INT, Literal: number}

		} else {
			t = createToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return t
}

func (l *Lexer) eatWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return ('a' <= char && 'z' >= char) || ('A' <= char && 'z' >= char) || '_' == char
}

func isNumber(char byte) bool {
	return '0' <= char && '9' >= char
}

func (l *Lexer) readWhile(f func(byte) bool) string {
	start := l.position

	for f(l.char) {
		l.readChar()
	}

	return l.input[start:l.position]
}
