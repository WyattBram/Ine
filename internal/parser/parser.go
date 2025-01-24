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

	p := makeParser(t)

	for p.containsTokens() {
		//Body := append(Body, parseStatement(p))
	}

	return ast.Program{
		Statements: Body,
	}
}

func makeParser(t []lexer.Token) *parser {
	createTokenLookup()
	p := parser{
		tokens:   t,
		position: 0,
	}
	return &p
}

func (p *parser) tokenType() lexer.Type {
	return p.currToken().Identifier
}

func (p *parser) currToken() lexer.Token {
	return p.tokens[p.position]
}

func (p *parser) nextToken() lexer.Token {
	curr := p.currToken()
	p.position += 1
	return curr
}

func (p *parser) containsTokens() bool {
	return p.position < len(p.tokens) && p.currToken().Identifier != lexer.EOF
}
