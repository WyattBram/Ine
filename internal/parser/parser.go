package parser

import (
	//"fmt"
	"fmt"

	"github.com/WyattBram/Ine/internal/ast"
	"github.com/WyattBram/Ine/internal/lexer"
)

type parser struct {
	tokens   []lexer.Token
	position int
	//tokens := lexer.Tokenizer()
}

func TestParse() {
	par := parser{
		tokens:   lexer.Tokenizer(),
		position: 0,
	}

	for x := range par.tokens {
		fmt.Print(par.tokens[x])
	}
}

func Parse(t []lexer.Token) ast.Program {
	Body := make([]ast.Statement, 0)

	return ast.Program{
		Statements: Body,
	}
}

func currToken(p *parser) lexer.Token {
	return p.tokens[p.position]
}

func nextToken(p *parser) lexer.Token {
	curr := currToken(p)
	p.position += 1
	return curr
}

func containsTokens(p *parser) bool {
	return p.position < len(p.tokens) && currToken(p).Identifier != lexer.EOF
}
