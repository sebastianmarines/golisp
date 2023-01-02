package reader

import (
	"golisp/ast"
	"testing"
)
import "golisp/lexer"

// This test should traverse the AST
func TestReader(t *testing.T) {
	tests := []struct {
		str      string
		expected *ast.Node
	}{
		{"()", &ast.Node{Type: ast.List, Children: nil}},
		{"a", &ast.Node{Type: ast.Symbol, Value: "a"}},
		{"(a)", &ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "a"}}}},
		{"(a b)", &ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "a"}, {Type: ast.Symbol, Value: "b"}}}},
		{"(\"a\")", &ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.String, Value: "a"}}}},
		{"((34))", &ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.List, Children: []*ast.Node{{Type: ast.Number, Value: 34}}}}}},
		{"-34", &ast.Node{Type: ast.Number, Value: -34}},
		{"nil", &ast.Node{Type: ast.Nil, Value: nil}},
		{"true", &ast.Node{Type: ast.True, Value: true}},
		{"false", &ast.Node{Type: ast.False, Value: false}},
	}

	for _, test := range tests {
		testReader(t, test.str, test.expected)
	}
}

func TestMultiline(t *testing.T) {
	str := `
		(+ 1 2)
		(+ 3 4)
		(+
			5
			6
		)
	`
	tokens := lexer.Tokenize(str)
	reader := NewReader(tokens)

	node := reader.Read()

	if node.Type != ast.List {
		t.Errorf("expected list, got %v", node.Type)
	}

	if len(node.Children) != 3 {
		t.Errorf("expected 3 children, got %v", len(node.Children))
	}

	node = reader.Read()

	if node.Type != ast.List {
		t.Errorf("expected list, got %v", node.Type)
	}

	if len(node.Children) != 3 {
		t.Errorf("expected 3 children, got %v", len(node.Children))
	}

	node = reader.Read()

	if node.Type != ast.List {
		t.Errorf("expected list, got %v", node.Type)
	}

	if len(node.Children) != 3 {
		t.Errorf("expected 3 children, got %v", len(node.Children))
	}

	node = reader.Read()

	if node != nil {
		t.Errorf("expected nil, got %v", node)
	}
}

func testReader(t *testing.T, str string, expected *ast.Node) {
	tokens := lexer.Tokenize(str)
	reader := NewReader(tokens)
	node := reader.Read()

	compareAST(t, str, node, expected)
}

func compareAST(t *testing.T, str string, current, expected *ast.Node) {
	if current.Type != expected.Type {
		t.Errorf("[%v] expected type %v, got %v", str, expected.Type, current.Type)
	}

	if current.Value != expected.Value {
		t.Errorf("[%v] expected value %v, got %v", str, expected.Value, current.Value)
	}

	if len(current.Children) != len(expected.Children) {
		t.Errorf("[%v] expected %v children, got %v", str, len(expected.Children), len(current.Children))
	}

	for i, child := range current.Children {
		compareAST(t, str, child, expected.Children[i])
	}
}
