package syntax

import (
	"github.com/tkech17/jack_compiler/utils/maths"
	"strings"
	"unicode"
)

type JackKeywordsAlias = int

type jackKeywordsList struct {
	Keyword    JackKeywordsAlias
	Symbol     JackKeywordsAlias
	Identifier JackKeywordsAlias
	NumberType JackKeywordsAlias
	StringType JackKeywordsAlias
	Other      JackKeywordsAlias
}

var Enum = &jackKeywordsList{
	Keyword:    0,
	Symbol:     1,
	Identifier: 2,
	NumberType: 3,
	StringType: 4,
	Other:      5,
}

var TokenTypes []string = []string{"keyword", "symbol", "identifier", "integerConstant", "stringConstant", "none"}
var ClassKeywords []string = []string{"static", "field"}
var FunctionTypes []string = []string{"constructor", "function", "method"}
var VariableTypes []string = []string{"int", "char", "boolean"}
var Statements []string = []string{"if", "let", "else", "while", "do", "return"}
var Operators []string = []string{"+", "-", "*", "/", "&", "|", "<", ">", "=", "~"}
var UnaryOperators []string = []string{"-", "~"}
var KeywordConstant []string = []string{"true", "false", "null", "this"}

var Keywords []string = []string{
	"class", "constructor", "function",
	"method", "field", "static", "var",
	"int", "char", "boolean", "void", "true",
	"false", "null", "this", "let", "do",
	"if", "else", "while", "return",
}

var Symbols []string = []string{
	"{", "}", "(", ")", "[", "]", ".",
	",", ";", "+", "-", "*", "/", "&",
	"|", "<", ">", "=", "~",
}

func GetTokenType(token string) JackKeywordsAlias {
	if contains(Symbols, token) {
		return Enum.Symbol
	} else if contains(Keywords, token) {
		return Enum.Keyword
	} else if strings.HasPrefix(token, "\"") {
		return Enum.StringType
	} else if maths.IsNumber(token) {
		return Enum.NumberType
	} else if token[0] == '_' || unicode.IsLetter(rune(token[0])) {
		return Enum.Identifier
	}
	return Enum.Other
}

func contains(elements []string, elem string) bool {
	for _, arrElem := range elements {
		if arrElem == elem {
			return true
		}
	}
	return false
}
