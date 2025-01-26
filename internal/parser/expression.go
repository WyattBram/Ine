package parser

import (
	"strconv"

	"github.com/WyattBram/Ine/internal/ast"
	"github.com/WyattBram/Ine/internal/lexer"
)

func parseExpr(p *parser, bp bindingPower) ast.Expresion {
	tokenType := p.tokenType()
	nud, exists := nud_Lookup[tokenType]

	if !exists {
		panic("Oh No")
	}

	left := nud(p)

	for bp_Lookup[p.tokenType()] > bp {
		tokenType = p.tokenType()
		led, exists := led_Lookup[tokenType]

		if !exists {
			panic("Oh No")
		}

		left = led(p, left, bp)
	}

	return left

}

func parsePrimary(p *parser) ast.Expresion {
	switch p.tokenType() {

	case lexer.Number:
		number, _ := strconv.ParseFloat(p.nextToken().Lexemes, 64)
		return ast.NumericalLiteral{
			Value: number,
		}
	case lexer.String:
		return ast.StringLiteral{
			Value: p.nextToken().Lexemes,
		}
	case lexer.Identifier:
		return ast.StringLiteral{
			Value: p.nextToken().Lexemes,
		}

	default:
		panic("WTF")
	}
}

func parseBinary(p *parser, left ast.Expresion, bp bindingPower) ast.Expresion {
	operator := p.nextToken()
	right := parseExpr(p, bp)

	return ast.BinaryExpresion{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
