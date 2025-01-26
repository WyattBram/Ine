package main

import (
	"github.com/WyattBram/Ine/internal/lexer"
	"github.com/WyattBram/Ine/internal/parser"
	"github.com/sanity-io/litter"
)

func main() {
	//parser.TestParse()

	ast := parser.Parse(lexer.Tokenizer())
	litter.Dump(ast)
}
