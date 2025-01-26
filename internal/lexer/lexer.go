package lexer

import (
	"fmt"
	"os"
	"regexp"
	"unicode"
)

type Token struct {
	Lexemes    string
	Identifier Type
}

type Type int

const (
	Identifier    = iota
	Keyword       = iota
	Number        = iota
	Plus          = iota
	Minus         = iota
	Multiply      = iota
	Divide        = iota
	Mod           = iota
	Equals        = iota
	EOF           = iota
	String        = iota
	GreaterThan   = iota
	LessThan      = iota
	Not           = iota
	GreaterEqual  = iota
	LessEqual     = iota
	Equivelent    = iota
	NotEquivelent = iota
	And           = iota
	Or            = iota
	NULL          = iota
)

var Reserved = map[string]Type{
	"fn":    Keyword,
	"int":   Keyword,
	"str":   Keyword,
	"bool":  Keyword,
	"float": Keyword,
	"and":   And,
	"or":    Or,
}

type regex regexp.Regexp

var (
	WordRGX          = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)
	GreaterThanRGX   = regexp.MustCompile(`^>$`)
	GreaterEqualRGX  = regexp.MustCompile(`^>=$`)
	LessThanRGX      = regexp.MustCompile(`^<$`)
	LessThanEqualRGX = regexp.MustCompile(`^<=$`)
	EquivlentRGX     = regexp.MustCompile(`^==$`)
	NotEquivelentRGX = regexp.MustCompile(`^!=$`)
	EqualsRGX        = regexp.MustCompile(`^=$`)
	AddRGX           = regexp.MustCompile(`^\+$`)
	SubtractRGX      = regexp.MustCompile(`^-$`)
	MultiplyRGX      = regexp.MustCompile(`^\*$`)
	DivideRGX        = regexp.MustCompile(`^/$`)
	ModRGX           = regexp.MustCompile(`^%$`)
	NumberRGX        = regexp.MustCompile(`^\d*\.?\d*$`)
	NotRGX           = regexp.MustCompile(`^!$`)
)

func matchStringTo(s string) Type {

	if WordRGX.MatchString(s) {
		return defineWord(s)
	} else if NumberRGX.MatchString(s) {
		return Number
	} else if GreaterThanRGX.MatchString(s) {
		return GreaterThan
	} else if GreaterEqualRGX.MatchString(s) {
		return GreaterEqual
	} else if LessThanRGX.MatchString(s) {
		return LessThan
	} else if LessThanEqualRGX.MatchString(s) {
		return LessEqual
	} else if EquivlentRGX.MatchString(s) {
		return Equivelent
	} else if NotEquivelentRGX.MatchString(s) {
		return NotEquivelent
	} else if EqualsRGX.MatchString(s) {
		return Equals
	} else if AddRGX.MatchString(s) {
		return Plus
	} else if SubtractRGX.MatchString(s) {
		return Minus
	} else if MultiplyRGX.MatchString(s) {
		return Multiply
	} else if DivideRGX.MatchString(s) {
		return Divide
	} else if ModRGX.MatchString(s) {
		return Mod
	} else if NotRGX.MatchString(s) {
		return Not
	} else {
		return NULL
	}

}

func makeToken(val string, ident Type) Token {
	return Token{Lexemes: val, Identifier: ident}
}

func fileToWords(byteList []byte) []string {
	wordList := []string{}

	currWord := ""
	for byte := range byteList {
		currByte := byteList[byte]
		if unicode.IsSpace(rune(currByte)) {
			if len(currWord) == 0 {
				continue
			}
			wordList = append(wordList, currWord)
			currWord = ""
		} else {
			currWord += string(currByte)
		}
	}
	return wordList
}

func defineWord(word string) Type {

	typeReserved, reserved := Reserved[word]

	//we go reserved
	if reserved {
		return typeReserved
	}

	return Identifier

}

func Tokenizer() []Token {
	data, err := os.ReadFile("test.ine")
	if err != nil {
		fmt.Print("err")
	}

	wordList := fileToWords(data)
	tokenList := []Token{}

	for word := range wordList {
		tokenList = append(tokenList, makeToken(wordList[word], matchStringTo(wordList[word])))

	}

	tokenList = append(tokenList, makeToken("EOF", EOF))

	return tokenList

}
