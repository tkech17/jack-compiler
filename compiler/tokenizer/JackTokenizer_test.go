package tokenizer_test

import (
	"github.com/tkech17/jack_compiler/compiler/tokenizer"
	"strings"
	"testing"
)

func TestGetTokensContent(t *testing.T) {
	var result string = tokenizer.GetTokensContent(mainJack)
	var expected string = strings.ReplaceAll(mainJackTXMLExpected, " ", "")
	expected = strings.ReplaceAll(expected, "\n", "")
	var actual string = strings.ReplaceAll(result, " ", "")
	actual = strings.ReplaceAll(actual, "\n", "")

	for i := range actual {
		var expectedChar uint8 = expected[i]
		var actualChar uint8 = actual[i]
		if expectedChar != actualChar {
			t.Fatal(i)
		}
	}
}

const mainJack = "// This file is part of www.nand2tetris.org\n" +
	"// and the book \"The Elements of Computing Systems\"\n" +
	"// by Nisan and Schocken, MIT Press.\n" +
	"// File name: projects/10/ArrayTest/Main.jack\n" +
	"\n" +
	"// (identical to projects/09/Average/Main.jack)\n" +
	"\n" +
	"/** Computes the average of a sequence of integers. */\n" +
	"class Main {\n" +
	"    function void main() {\n" +
	"        var Array a;\n" +
	"        var int length;\n" +
	"        var int i, sum;\n" +
	"\t" +
	"\n" +
	"\t" +
	"let length = Keyboard.readInt(\"HOW MANY NUMBERS? \");\n" +
	"\t" +
	"let a = Array.new(length);\n" +
	"\t" +
	"let i = 0;\n" +
	"\t" +
	"\n" +
	"\t" +
	"while (i < length) {\n" +
	"\t" +
	"    let a[i] = Keyboard.readInt(\"ENTER THE NEXT NUMBER: \");\n" +
	"\t" +
	"    let i = i + 1;\n" +
	"\t" +
	"}\n" +
	"\t" +
	"\n" +
	"\t" +
	"let i = 0;\n" +
	"\t" +
	"let sum = 0;\n" +
	"\t" +
	"\n" +
	"\t" +
	"while (i < length) {\n" +
	"\t" +
	"    let sum = sum + a[i];\n" +
	"\t" +
	"    let i = i + 1;\n" +
	"\t" +
	"}\n" +
	"\t" +
	"\n" +
	"\t" +
	"do Output.printString(\"THE AVERAGE IS: \");\n" +
	"\t" +
	"do Output.printInt(sum / length);\n" +
	"\t" +
	"do Output.println();\n" +
	"\t" +
	"\n" +
	"\t" +
	"return;\n" +
	"    }\n" +
	"}\n"

const mainJackTXMLExpected = "<tokens>\n" +
	"<keyword> class </keyword>\n" +
	"<identifier> Main </identifier>\n" +
	"<symbol> { </symbol>\n" +
	"<keyword> function </keyword>\n" +
	"<keyword> void </keyword>\n" +
	"<identifier> main </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> { </symbol>\n" +
	"<keyword> var </keyword>\n" +
	"<identifier> Array </identifier>\n" +
	"<identifier> a </identifier>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> var </keyword>\n" +
	"<keyword> int </keyword>\n" +
	"<identifier> length </identifier>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> var </keyword>\n" +
	"<keyword> int </keyword>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> , </symbol>\n" +
	"<identifier> sum </identifier>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> length </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<identifier> Keyboard </identifier>\n" +
	"<symbol> . </symbol>\n" +
	"<identifier> readInt </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<stringConstant> HOW MANY NUMBERS?  </stringConstant>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> a </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<identifier> Array </identifier>\n" +
	"<symbol> . </symbol>\n" +
	"<identifier> new </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<identifier> length </identifier>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<integerConstant> 0 </integerConstant>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> while </keyword>\n" +
	"<symbol> ( </symbol>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> &lt; </symbol>\n" +
	"<identifier> length </identifier>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> { </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> a </identifier>\n" +
	"<symbol> [ </symbol>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> ] </symbol>\n" +
	"<symbol> = </symbol>\n" +
	"<identifier> Keyboard </identifier>\n" +
	"<symbol> . </symbol>\n" +
	"<identifier> readInt </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<stringConstant> ENTER THE NEXT NUMBER:  </stringConstant>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> + </symbol>\n" +
	"<integerConstant> 1 </integerConstant>\n" +
	"<symbol> ; </symbol>\n" +
	"<symbol> } </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<integerConstant> 0 </integerConstant>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> sum </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<integerConstant> 0 </integerConstant>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> while </keyword>\n" +
	"<symbol> ( </symbol>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> &lt; </symbol>\n" +
	"<identifier> length </identifier>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> { </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> sum </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<identifier> sum </identifier>\n" +
	"<symbol> + </symbol>\n" +
	"<identifier> a </identifier>\n" +
	"<symbol> [ </symbol>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> ] </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> let </keyword>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> = </symbol>\n" +
	"<identifier> i </identifier>\n" +
	"<symbol> + </symbol>\n" +
	"<integerConstant> 1 </integerConstant>\n" +
	"<symbol> ; </symbol>\n" +
	"<symbol> } </symbol>\n" +
	"<keyword> do </keyword>\n" +
	"<identifier> Output </identifier>\n" +
	"<symbol> . </symbol>\n" +
	"<identifier> printString </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<stringConstant> THE AVERAGE IS:  </stringConstant>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> do </keyword>\n" +
	"<identifier> Output </identifier>\n" +
	"<symbol> . </symbol>\n" +
	"<identifier> printInt </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<identifier> sum </identifier>\n" +
	"<symbol> / </symbol>\n" +
	"<identifier> length </identifier>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> do </keyword>\n" +
	"<identifier> Output </identifier>\n" +
	"<symbol> . </symbol>\n" +
	"<identifier> println </identifier>\n" +
	"<symbol> ( </symbol>\n" +
	"<symbol> ) </symbol>\n" +
	"<symbol> ; </symbol>\n" +
	"<keyword> return </keyword>\n" +
	"<symbol> ; </symbol>\n" +
	"<symbol> } </symbol>\n" +
	"<symbol> } </symbol>\n" +
	"</tokens>\n"
