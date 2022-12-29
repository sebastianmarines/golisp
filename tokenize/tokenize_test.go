package tokenize

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		str      string
		expected []Token
	}{
		{"()", []Token{{"(", "("}, {")", ")"}}},
		{"a", []Token{{"symbol", "a"}}},
		{"(a)", []Token{{"(", "("}, {"symbol", "a"}, {")", ")"}}},
		{"(a b)", []Token{{"(", "("}, {"symbol", "a"}, {"symbol", "b"}, {")", ")"}}},
		{"(\"a\")", []Token{{"(", "("}, {"string", "a"}, {")", ")"}}},
		{"((34))", []Token{{"(", "("}, {"(", "("}, {"symbol", "34"}, {")", ")"}, {")", ")"}}},
	}

	for _, test := range tests {
		testTokens(t, test.str, test.expected)
	}
}

func testTokens(t *testing.T, str string, expected []Token) {
	tokens := Tokenize(str)
	if len(tokens) != len(expected) {
		t.Errorf("expected %v tokens, got %v", len(expected), len(tokens))
	}

	for i, token := range tokens {
		if token != expected[i] {
			t.Errorf("expected token %v to be %v, got %v", i, expected[i], token)
		}
	}
}
