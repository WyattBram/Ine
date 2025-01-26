package main

import (
	"github.com/WyattBram/Ine/internal/lexer"
	"github.com/WyattBram/Ine/internal/parser"
	"github.com/sanity-io/litter"
)

func main() {
	ast := parser.Parse(lexer.Tokenizer())

	litter.Dump(ast)
}
