package parser

import (
	"strconv"

	"github.com/WyattBram/Ine/internal/ast"
	"github.com/WyattBram/Ine/internal/lexer"
)

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
