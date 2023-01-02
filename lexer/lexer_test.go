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
		{"(\"a string\")", []Token{{LeftParen, "("}, {String, "a string"}, {RightParen, ")"}}},
		{"(a \"a string\")", []Token{{LeftParen, "("}, {Symbol, "a"}, {String, "a string"}, {RightParen, ")"}}},
		{"\"hello \\\"world\\\"\"", []Token{{String, "hello \"world\""}}},
		{"123", []Token{{Integer, "123"}}},
		{"(123)", []Token{{LeftParen, "("}, {Integer, "123"}, {RightParen, ")"}}},
		{"-123", []Token{{Integer, "-123"}}},
		{"\"hello \nworld\"", []Token{{String, "hello \nworld"}}},
	}

	for _, test := range tests {
		testTokens(t, test.str, test.expected)
	}
}

func testTokens(t *testing.T, str string, expected []Token) {
	tokens := Tokenize(str)
	if len(tokens) != len(expected) {
		t.Errorf("[%v] expected %v tokens, got %v", tokens, len(expected), len(tokens))
	}

	for i, token := range tokens {
		if token != expected[i] {
			t.Errorf("[%v] expected token %v to be %v, got %v", tokens, i, expected[i], token)
		}
	}
}
