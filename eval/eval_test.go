package eval

import (
	"golisp/ast"
	"testing"
)

func TestEval(t *testing.T) {
	env := NewEnv(nil)
	tests := []struct {
		input    ast.Node
		expected string
	}{
		{input: ast.Node{Type: ast.Number, Value: 123}, expected: "123"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}, expected: "3"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 3}}}}}, expected: "6"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.String, Value: "foo"}, {Type: ast.String, Value: "bar"}}}, expected: "\"foobar\""},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "*"}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 3}}}, expected: "6"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "/"}, {Type: ast.Number, Value: 6}, {Type: ast.Number, Value: 3}}}, expected: "2"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "-"}, {Type: ast.Number, Value: 6}, {Type: ast.Number, Value: 3}}}, expected: "3"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "def!"}, {Type: ast.Symbol, Value: "a"}, {Type: ast.Number, Value: 123}}}, expected: "123"},
		{input: ast.Node{Type: ast.Symbol, Value: "a"}, expected: "123"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "def!"}, {Type: ast.Symbol, Value: "a"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}}}, expected: "3"},
		{input: ast.Node{Type: ast.Symbol, Value: "a"}, expected: "3"},
	}

	for _, test := range tests {
		testEval(t, env, test.input, test.expected)
	}
}

func testEval(t *testing.T, env *Env, input ast.Node, expected string) {
	result := Eval(&input, env)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
