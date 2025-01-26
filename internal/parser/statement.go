package parser

import "github.com/WyattBram/Ine/internal/ast"

func parseStatement(p *parser) ast.Statement {
	statementfn, exists := statement_Lookup[p.tokenType()]

	if exists {
		return statementfn(p)
	}

	expr := parseExpr(p, defaultBinding)

	return ast.ExpressionStatement{
		Expression: expr,
	}

}
