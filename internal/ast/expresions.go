package ast

type binaryExpresion struct {
	left      Expresion
	right     Expresion
	opperator string
}

func (n binaryExpresion) expr() {}

type NumericalLiteral struct {
	value float32
}

func (n NumericalLiteral) expr() {}

type Identifier struct {
	identifier string
}

func (n Identifier) expr() {}
