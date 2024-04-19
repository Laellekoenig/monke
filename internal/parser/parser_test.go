package parser

import (
	"testing"

	"github.com/Laellekoenig/monke/internal/ast"
	"github.com/Laellekoenig/monke/internal/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let foobar = 83838;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("Parse program returned nil.")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Program does not contain 3 statements. Got %d.", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Token literal is not 'let'. Got %s.", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Got %T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("Name not %s. Got %s.", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("TokenLiteral not %s. Got %s.", name, letStatement.Name.Value)
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser encountered %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
  return 5;
  return 10;
  return 993322;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements but got %d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		ret, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not *ast.ReturnStatement, got %T", statement)
			continue
		}

		if ret.TokenLiteral() != "return" {
			t.Errorf("return TokenLiteral is not 'return', got %s", ret.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program does not have right amount of statements, expected 1, got %d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression is not an ast.ExpressionStatementn, got %T", program.Statements[0])
	}

	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not an ast.Identifier, got %T", statement.Expression)
	}

	if identifier.Value != "foobar" {
		t.Errorf("identifier.Value is not foobar, got %s", identifier.Value)
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf("identifier.TokenLiteral is not foobar, got %s", identifier.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program does not have right amount of statements, expected 1, got %d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expression is not an ast.ExpressionStatementn, got %T", program.Statements[0])
	}

	identifier, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression is not an ast.Identifier, got %T", statement.Expression)
	}

	if identifier.Value != 5 {
		t.Errorf("identifier.Value is not foobar, got %d", identifier.Value)
	}

	if identifier.TokenLiteral() != "5" {
		t.Errorf("identifier.TokenLiteral is not foobar, got %s", identifier.TokenLiteral())
	}
}
