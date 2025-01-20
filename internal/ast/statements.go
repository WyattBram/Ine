package ast

type Program struct {
	Statements []Statement
}

func (p Program) stmt() {}

type ExpressionStatement struct {
	Expression Expresion
}

func (p ExpressionStatement) stmt() {}
