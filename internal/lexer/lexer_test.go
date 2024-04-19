package lexer

import (
	"testing"

	"github.com/Laellekoenig/monke/internal/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - wrong token type. Got %v but expected %v", i, token.Type, testToken.expectedType)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. Got %v but expected %v", i, token.Literal, testToken.expectedLiteral)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `let five = 5;
  let ten = 10;

  let add = fn(x, y) {
    x + y;
  }

  let result = add(five, ten);
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - wrong token type. Got %v but expected %v. Token literal '%v'", i, token.Type, testToken.expectedType, token.Literal)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. Got %v but expected %v", i, token.Literal, testToken.expectedLiteral)
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `let five = 5;
  let ten = 10;
  let add = fn(x, y) {
  x + y;
  };
  let result = add(five, ten);
  !-/*5;
  5 < 10 > 5;
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERIX, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - wrong token type. Got %v but expected %v. Token literal '%v'", i, token.Type, testToken.expectedType, token.Literal)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. Got %v but expected %v", i, token.Literal, testToken.expectedLiteral)
		}
	}
}

func TestNextToken4(t *testing.T) {
	input := `

  if (5 < 10) {
    return true;
  } else {
    return false;
  }
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - wrong token type. Got %v but expected %v. Token literal '%v'", i, token.Type, testToken.expectedType, token.Literal)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. Got %v but expected %v", i, token.Literal, testToken.expectedLiteral)
		}
	}
}

func TestNextToken5(t *testing.T) {
	input := `

    10 == 10;
    10 != 9;
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, testToken := range tests {
		token := l.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - wrong token type. Got %v but expected %v. Token literal '%v'", i, token.Type, testToken.expectedType, token.Literal)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. Got %v but expected %v", i, token.Literal, testToken.expectedLiteral)
		}
	}
}
