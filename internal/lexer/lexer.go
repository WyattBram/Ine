package lexer

import (
	"fmt"
	"os"
)

type token struct {
	Value      string
	Identifier int16
}

func Tokenizer() string {
	data, err := os.ReadFile("test.ine")
	if err != nil {
		return "Problem"
	}

	tokens := []string{}

	for i := range data {
		if string(data[i]) == "\n" {
			continue
		}
		tokens = append(tokens, string(data[i]))
	}

	fmt.Println(tokens)

	return "This works"
}
