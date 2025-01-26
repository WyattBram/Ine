package ast

import "github.com/WyattBram/Ine/internal/lexer"

type BinaryExpresion struct {
	Left     Expresion
	Operator lexer.Token
	Right    Expresion
}

func (b BinaryExpresion) expr() {}

type StringLiteral struct {
	Value string
}

func (s StringLiteral) expr() {}

type NumericalLiteral struct {
	Value float64
}

func (n NumericalLiteral) expr() {}

type MathmaticalLiteral struct {
	Value string
}

func (m MathmaticalLiteral) expr() {}

type Identifier struct {
	identifier string
}

func (i Identifier) expr() {}
