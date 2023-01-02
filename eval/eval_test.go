package eval

import (
	"github.com/chzyer/readline"
	"golisp/ast"
	"io"
	"os"
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
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "let*"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "a"}, {Type: ast.Number, Value: 123}}}, {Type: ast.Symbol, Value: "a"}}}, expected: "123"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "def!"}, {Type: ast.Symbol, Value: "a"}, {Type: ast.String, Value: "foo \n bar"}}}, expected: "\"foo \\n bar\""},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "do"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 3}}}, expected: "3"},
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "do"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}}}, expected: "3"},
	}

	for _, test := range tests {
		testEval(t, env, test.input, test.expected)
	}
}

func TestPrn(t *testing.T) {
	rescueStdout := readline.Stdout
	r, w, _ := os.Pipe()
	readline.Stdout = w

	test := ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "prn"}, {Type: ast.String, Value: "Hello"}}}
	Eval(&test, NewEnv(nil))

	err := w.Close()
	if err != nil {
		return
	}
	out, _ := io.ReadAll(r)
	readline.Stdout = rescueStdout

	if string(out) != "Hello\n" {
		t.Errorf("expected \"Hello\", got %q", string(out))
	}
}

func testEval(t *testing.T, env *Env, input ast.Node, expected string) {
	result := Eval(&input, env).PrStr(false)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
