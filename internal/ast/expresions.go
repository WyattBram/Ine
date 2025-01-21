package ast

type binaryExpresion struct {
	left  Expresion
	right Expresion
	// opperator string
}

func (b binaryExpresion) expr() {}

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
