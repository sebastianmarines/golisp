package lexer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		str      string
		expected []Token
	}{
		{"()", []Token{{LeftParen, "("}, {RightParen, ")"}}},
		{"a", []Token{{Symbol, "a"}}},
		{"(a b)", []Token{{LeftParen, "("}, {Symbol, "a"}, {Symbol, "b"}, {RightParen, ")"}}},
		{"( a)", []Token{{LeftParen, "("}, {Symbol, "a"}, {RightParen, ")"}}},
		{"(\"a string\")", []Token{{LeftParen, "("}, {String, "a string"}, {RightParen, ")"}}},
		{"(a \"a string\")", []Token{{LeftParen, "("}, {Symbol, "a"}, {String, "a string"}, {RightParen, ")"}}},
		{"\"hello \\\"world\\\"\"", []Token{{String, "hello \"world\""}}},
		{"123", []Token{{Integer, "123"}}},
		{"(123)", []Token{{LeftParen, "("}, {Integer, "123"}, {RightParen, ")"}}},
		{"-123", []Token{{Integer, "-123"}}},
		{"\"hello \nworld\"", []Token{{String, "hello \nworld"}}},
		{"(- 2 1)", []Token{{LeftParen, "("}, {Symbol, "-"}, {Integer, "2"}, {Integer, "1"}, {RightParen, ")"}}},
	}

	for _, test := range tests {
		testTokens(t, test.str, test.expected)
	}
}

func TestTokenizeWithComments(t *testing.T) {
	tests := []struct {
		str      string
		expected []Token
	}{
		{"; a comment ()", nil},
		{"(a ; a comment\n)", []Token{{LeftParen, "("}, {Symbol, "a"}, {RightParen, ")"}}},
		{"(a ; a\nb)", []Token{{LeftParen, "("}, {Symbol, "a"}, {Symbol, "b"}, {RightParen, ")"}}},
		{"(a ; a comment\nb ; another comment\n)", []Token{{LeftParen, "("}, {Symbol, "a"}, {Symbol, "b"}, {RightParen, ")"}}},
		{"(a ; a comment \n b)", []Token{{LeftParen, "("}, {Symbol, "a"}, {Symbol, "b"}, {RightParen, ")"}}},
	}

	for _, test := range tests {
		testTokens(t, test.str, test.expected)
	}
}

func testTokens(t *testing.T, str string, expected []Token) {
	tokens := Tokenize(str)
	if len(tokens) != len(expected) {
		t.Errorf("[%q] expected %v tokens, got %v", str, len(expected), len(tokens))
	}

	for i, token := range tokens {
		if token != expected[i] {
			t.Errorf("[%v] expected token %v to be %v, got %v", str, i, expected[i], token)
		}
	}
}
