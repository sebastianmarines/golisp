package ast

import (
	"testing"
)

func TestNode_String(t *testing.T) {
	tests := []struct {
		ast      Node
		expected string
	}{
		{Node{Symbol, "a", []*Node{}}, "a"},
		{Node{List, nil, []*Node{}}, "()"},
		{Node{List, nil, []*Node{{Symbol, "a", []*Node{}}}}, "(a)"},
		{Node{List, nil, []*Node{{Symbol, "a", []*Node{}}, {Symbol, "b", []*Node{}}}}, "(a b)"},
		{Node{List, nil, []*Node{{String, "a", []*Node{}}}}, "(\"a\")"},
		{Node{List, nil, []*Node{{List, nil, []*Node{{Symbol, "34", []*Node{}}}}}}, "((34))"},
		{Node{Number, 34, []*Node{}}, "34"},
		{Node{Nil, nil, []*Node{}}, "nil"},
		{Node{True, true, []*Node{}}, "true"},
		{Node{False, false, []*Node{}}, "false"},
		{Node{String, "a\nb", []*Node{}}, "\"a\\nb\""},
	}

	for _, test := range tests {
		if current := test.ast.PrStr(false); current != test.expected {
			t.Errorf("expected %s, got %s", test.expected, current)
		}
	}
}

func TestNode_String_WithReadably(t *testing.T) {
	tests := []struct {
		ast      Node
		expected string
	}{
		{Node{List, nil, []*Node{{String, "a", []*Node{}}}}, "(\"a\")"},
		{Node{String, "a\nb", []*Node{}}, "a\nb"},
	}

	for _, test := range tests {
		if current := test.ast.PrStr(true); current != test.expected {
			t.Errorf("expected %s, got %s", test.expected, current)
		}
	}
}
