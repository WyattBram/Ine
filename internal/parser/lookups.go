package parser

import (
	"github.com/WyattBram/Ine/internal/ast"
	"github.com/WyattBram/Ine/internal/lexer"
)

type bindingPower int

const (
	defaultBinding bindingPower = iota
	comma                       = iota
	assignment                  = iota
	logical                     = iota
	relational                  = iota
	additive                    = iota
	multiplicative              = iota
	unary                       = iota
	call                        = iota
	member                      = iota
	primary                     = iota
)

type statementHandler func(p *parser) ast.Statement
type nudHandler func(p *parser) ast.Expresion
type leftDenotedHandler func(p *parser, left ast.Expresion, bp bindingPower) ast.Statement

type statementLookup map[lexer.Type]statementHandler
type nudLookup map[lexer.Type]nudHandler
type leftDenotedLookup map[lexer.Type]leftDenotedHandler
type bindingPowerLookup map[lexer.Type]bindingPower

var bp_Lookup = bindingPowerLookup{}
var nud_Lookup = nudLookup{}
var led_Lookup = leftDenotedLookup{}
var statement_Lookup = statementLookup{}

func led(t lexer.Type, bp bindingPower, led_fn leftDenotedHandler) {
	bp_Lookup[t] = bp
	led_Lookup[t] = led_fn
}

func nud(t lexer.Type, bp bindingPower, nud_fn nudHandler) {
	bp_Lookup[t] = bp
	nud_Lookup[t] = nud_fn
}

func statement(t lexer.Type, stmt_fn statementHandler) {
	bp_Lookup[t] = defaultBinding
	statement_Lookup[t] = stmt_fn
}

func createTokenLookup() {
	nud(lexer.Number, primary, parsePrimary)
	nud(lexer.String, primary, parsePrimary)
	nud(lexer.Identifier, primary, parsePrimary)
}
