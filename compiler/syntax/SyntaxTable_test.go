package syntax_test

import (
	"github.com/tkech17/jack_compiler/compiler/syntax"
	"testing"
)

func TestGetTokenType(t *testing.T) {
	if syntax.GetTokenType("\"") != syntax.Enum.StringType {
		t.Fatal("[\"] is string type")
	}
	if syntax.GetTokenType("123") != syntax.Enum.NumberType {
		t.Fatal("[123] is number type")
	}
	if syntax.GetTokenType("_bla") != syntax.Enum.Identifier {
		t.Fatal("[_] is identifier")
	}
	if syntax.GetTokenType("+") != syntax.Enum.Symbol {
		t.Fatal("[+] is symbol")
	}
	if syntax.GetTokenType("static") != syntax.Enum.Keyword {
		t.Fatal("[static] is keyword")
	}
}
