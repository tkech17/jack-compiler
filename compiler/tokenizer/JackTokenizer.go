package tokenizer

import (
	"github.com/tkech17/jack_compiler/compiler/syntax"
	"github.com/tkech17/jack_compiler/utils/functional"
	"strings"
	"unicode"
)

func GetTokensContent(content string) string {
	var parsedTokens []string = getParsedTokens(content)
	var result string
	result += "<tokens>\n"
	for _, token := range parsedTokens {
		var taggedToken string = getTaggedToken(token)
		result += taggedToken
		result += "\n"
	}
	result += "</tokens>\n"
	return result
}

func getTaggedToken(token string) string {
	var tokenType string = getTokenType(token)
	if token == "&" {
		token = "&amp;"
	} else if token == "\"" {
		token = "&quot;"
	} else if token == ">" {
		token = "&gt;"
	} else if token == "<" {
		token = "&lt;"
	} else if isString(token[0]) {
		token = token[1 : len(token)-1]
	}
	var result string
	result += "<" + tokenType + ">"
	result += token + ""
	result += "</" + tokenType + ">"
	return result
}

func getTokenType(token string) string {
	var tokenType syntax.JackKeywordsAlias = syntax.GetTokenType(token)
	return syntax.TokenTypes[tokenType]
}

func getParsedTokens(content string) []string {
	var result []string
	var lines []string = strings.Split(content, "\n")
	lines = functional.MapString(lines, removeComments)
	lines = removeMultilineComments(lines)
	lines = functional.MapString(lines, strings.TrimSpace)
	lines = functional.Filter(lines, isNotEmpty)

	for _, line := range lines {
		var parsedLine []string = parseLine(line)
		result = append(result, parsedLine...)
	}
	return result
}

func parseLine(line string) []string {
	var tokens []string
	for i := 0; i < len(line); {
		var ch uint8 = line[i]
		if unicode.IsDigit(rune(ch)) {
			var number string
			number, i = getNumber(line, i)
			tokens = append(tokens, number)
		} else if isString(ch) {
			var stringLiteral string
			stringLiteral, i = getString(line, i)
			tokens = append(tokens, stringLiteral)
		} else if isIdentifier(ch) {
			var identifier string
			identifier, i = getIdentifier(line, i)
			tokens = append(tokens, identifier)
		} else if ch != ' ' {
			tokens = append(tokens, string(ch))
			i++
		} else {
			i++
		}
	}
	return tokens
}

func getIdentifier(line string, i int) (string, int) {
	var identifier string
	for (i < len(line)) && (isIdentifier(line[i]) || unicode.IsDigit(rune(line[i]))) {
		identifier += string(line[i])
		i++
	}
	return identifier, i
}

func isIdentifier(ch uint8) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func getString(line string, i int) (string, int) {
	var stringLiteral string
	stringLiteral += string(line[i])
	i++
	for line[i] != '"' {
		stringLiteral += string(line[i])
		i++
	}
	stringLiteral += string(line[i])
	i++
	return stringLiteral, i
}

func isString(u uint8) bool {
	return u == '"'
}

func getNumber(line string, i int) (string, int) {
	var number string
	for (i < len(line)) && unicode.IsDigit(rune(line[i])) {
		number += string(line[i])
		i++
	}
	return number, i
}

func removeMultilineComments(lines []string) []string {
	var result []string
	var isComment = false
	for _, line := range lines {
		var endOfComment int = strings.Index(line, "*/")
		if endOfComment != -1 {
			line = line[endOfComment+2:]
			isComment = false
		}

		var startOfComment int = strings.Index(line, "/*")
		if startOfComment != -1 {
			line = line[0:startOfComment]
			isComment = true
		}

		if !isComment {
			result = append(result, line)
		}
	}

	return result
}

func isNotEmpty(elem string) bool {
	return len(elem) != 0
}

func removeComments(str string) string {
	var indexOfComment = strings.Index(str, "//")
	if indexOfComment == -1 {
		return str
	}
	return str[0:indexOfComment]
}
