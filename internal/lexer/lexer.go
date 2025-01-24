package lexer

import (
	"fmt"
	"os"
	"unicode"
)

type Type int

const (
	Identifier   = iota
	Keyword      = iota
	Number       = iota
	Plus         = iota
	Minus        = iota
	Multiply     = iota
	Divide       = iota
	Mod          = iota
	Equals       = iota
	EOF          = iota
	String       = iota
	GreaterThan  = iota
	LessThan     = iota
	Not          = iota
	GreaterEqual = iota
	LessEqual    = iota
	Equivelent   = iota
	NULL         = iota
)

type Token struct {
	Lexemes    string
	Identifier Type
}

var Reserved = map[string]Type{
	"fn":    Keyword,
	"int":   Keyword,
	"str":   Keyword,
	"bool":  Keyword,
	"float": Keyword,
}

func tellOperator(r rune) bool {
	switch r {
	case '+', '-', '*', '/', '%':
		return true
	}
	return false
}

func tellCompare(s string) bool {
	switch s {
	case ">", "<", "<=", ">=", "==", "=":
		return true
	}
	return false
}

func whichOperator(r rune) Type {
	switch r {
	case '+':
		return Plus

	case '-':
		return Minus

	case '*':
		return Multiply

	case '/':
		return Divide

	case '%':
		return Mod
	}
	return NULL
}

func whichCompare(s string) Type {
	switch s {
	case ">":
		return GreaterThan

	case "<":
		return LessThan

	case ">=":
		return GreaterEqual

	case "<=":
		return LessEqual
	}
	return NULL
}

func makeToken(val string, ident Type) Token {
	return Token{Lexemes: val, Identifier: ident}
}

func Tokenizer() []Token {
	data, err := os.ReadFile("test.ine")
	if err != nil {
		fmt.Print("err")
	}

	tokens := []Token{}
	Word := ""
	Pass := false

	for i := range data {
		if Pass {
			Pass = false
			continue
		}
		currCharcter := rune(data[i])

		if currCharcter == '=' && len(Word) == 0 {
			if i < len(data)-1 && data[i+1] == '=' {
				tokens = append(tokens, makeToken("==", Equivelent))
				Pass = true
			} else {
				tokens = append(tokens, makeToken(string(currCharcter), Equals))
			}

			fmt.Print(Word)
		} else if unicode.IsLetter(currCharcter) || len(Word) != 0 || tellCompare(string(currCharcter)) {
			if !unicode.IsLetter(currCharcter) && !tellCompare(string(currCharcter)) {
				_, ok := Reserved[Word]
				if ok {
					tokens = append(tokens, makeToken(Word, Keyword))
				} else if tellCompare(Word) {
					tokens = append(tokens, makeToken(Word, whichCompare(Word)))
				} else {
					tokens = append(tokens, makeToken(Word, Identifier))
				}

				Word = ""
			} else {
				Word += string(currCharcter)
				continue
			}
		} else if unicode.IsDigit(currCharcter) {
			tokens = append(tokens, makeToken(string(currCharcter), Number))
		} else if tellOperator(currCharcter) {
			tokens = append(tokens, makeToken(string(currCharcter), whichOperator(currCharcter)))
		}
	}

	tokens = append(tokens, makeToken("EOF", EOF))

	return tokens

}
