package lexer

import (
	"fmt"
	"os"
	"unicode"
)

type Type int

const (
	Identifier = iota
	Keyword    = iota
	Number     = iota
	Operator   = iota
	Equals     = iota
	EOF        = iota
)

type token struct {
	Value      string
	Identifier int16
}

var Reserved = map[string]Type{
	"test": Keyword,
}

func tellOperator(r rune) bool {
	switch r {
	case '+', '-', '*', '/':
		return true
	}
	return false
}

func makeToken(val string, ident int16) token {
	return token{Value: val, Identifier: ident}
}

func Tokenizer() []token {
	data, err := os.ReadFile("test.ine")
	if err != nil {
		fmt.Print("err")
	}

	tokens := []token{}
	Word := ""

	for i := range data {
		currCharcter := rune(data[i])

		if unicode.IsLetter(currCharcter) || len(Word) != 0 {
			if !unicode.IsLetter(currCharcter) {
				_, ok := Reserved[Word]
				if ok {
					tokens = append(tokens, makeToken(Word, Keyword))
				} else {
					tokens = append(tokens, makeToken(Word, Identifier))
				}
				Word = ""

			} else {
				Word += string(currCharcter)
				continue
			}
		}

		if unicode.IsDigit(currCharcter) {
			tokens = append(tokens, makeToken(string(currCharcter), Number))
		} else if tellOperator(currCharcter) {
			tokens = append(tokens, makeToken(string(currCharcter), Operator))
		} else if currCharcter == '=' {
			tokens = append(tokens, makeToken(string(currCharcter), Equals))
		}
	}

	tokens = append(tokens, makeToken("EOF", EOF))

	return tokens

}
