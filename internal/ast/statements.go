package ast

type Program struct {
	Statements []Statement
}

func (p Program) stmt() {}

type ExpressionStatement struct {
	Expression Expresion
}

func (p ExpressionStatement) stmt() {}

type VariableDeclerationStatement struct {
	VariableIdentifier string
	AssignedValue      Expresion
}

func (p VariableDeclerationStatement) stmt() {
}
