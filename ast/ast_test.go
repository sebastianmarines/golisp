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
	}

	for _, test := range tests {
		testString(t, test.ast, test.expected)
	}
}

func testString(t *testing.T, ast Node, expected string) {
	if current := ast.String(); current != expected {
		t.Errorf("expected %v, got %v", expected, current)
	}
}
