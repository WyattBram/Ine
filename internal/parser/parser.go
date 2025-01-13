package parser

import (
	"fmt"
	"github.com/WyattBram/Ine/internal/lexer"
)

func Parse() {
	tokens := lexer.Tokenizer()

	for x := range tokens {
		fmt.Printf("%+v\n", tokens[x])
	}
}
