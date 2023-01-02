package eval

import (
	"github.com/chzyer/readline"
	"golisp/ast"
	"io"
	"os"
	"testing"
)

func TestArithmetic(t *testing.T) {
	env := NewEnv(nil)
	tests := []struct {
		input    ast.Node
		expected string
	}{
		// 123
		{input: ast.Node{Type: ast.Number, Value: 123}, expected: "123"},
		// (+ 1 2)
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}, expected: "3"},
		// (+ 1 (+ 2 3))
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 3}}}}}, expected: "6"},
		// (+ "foo" "bar")
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.String, Value: "foo"}, {Type: ast.String, Value: "bar"}}}, expected: "\"foobar\""},
		// (* 2 3)
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "*"}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 3}}}, expected: "6"},
		// (/ 6 3)
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "/"}, {Type: ast.Number, Value: 6}, {Type: ast.Number, Value: 3}}}, expected: "2"},
		// (- 6 3)
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "-"}, {Type: ast.Number, Value: 6}, {Type: ast.Number, Value: 3}}}, expected: "3"},
	}

	for _, test := range tests {
		testEval(t, env, test.input, test.expected)
	}
}

func TestInternalFunctions(t *testing.T) {
	env := NewEnv(nil)
	tests := []struct {
		input    ast.Node
		expected string
	}{
		// (def! a 123) => 123
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "def!"}, {Type: ast.Symbol, Value: "a"}, {Type: ast.Number, Value: 123}}}, expected: "123"},
		{input: ast.Node{Type: ast.Symbol, Value: "a"}, expected: "123"},

		// (def! b (+ 1 2)) => 3
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "def!"}, {Type: ast.Symbol, Value: "b"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}}}, expected: "3"},
		{input: ast.Node{Type: ast.Symbol, Value: "b"}, expected: "3"},

		// (let* (a 123) a) => 123
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "let*"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "a"}, {Type: ast.Number, Value: 123}}}, {Type: ast.Symbol, Value: "a"}}}, expected: "123"},

		// (def! c "foo \n bar") => "foo \n bar"
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "def!"}, {Type: ast.Symbol, Value: "c"}, {Type: ast.String, Value: "foo \n bar"}}}, expected: "\"foo \\n bar\""},

		// (do 1 2 3) => 3
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "do"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 3}}}, expected: "3"},

		// (do 1 2 (+ 1 2)) => 3
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "do"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}}}, expected: "3"},
	}

	for _, test := range tests {
		testEval(t, env, test.input, test.expected)
	}
}

func TestComparison(t *testing.T) {
	env := NewEnv(nil)
	tests := []struct {
		input    ast.Node
		expected string
	}{
		// (= 1 1) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "="}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, expected: "true"},
		// (= 1 2) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "="}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}, expected: "false"},
		// (= (1 1) (1 1)) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "="}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}}}, expected: "true"},
		// (= (1 1) (1 2)) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "="}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}}}, expected: "false"},
		// (= (+ 1 1) 2) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "="}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 2}}}, expected: "true"},
		// (= (+ 1 1) 3) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "="}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 3}}}, expected: "false"},

		// (> 1 1) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, expected: "false"},
		// (> 1 2) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}, expected: "false"},
		// (> 2 1) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">"}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 1}}}, expected: "true"},
		// (> (+ 1 1) 2) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 2}}}, expected: "false"},
		// (> (+ 1 1) 1) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 1}}}, expected: "true"},

		// (< 1 1) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "<"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, expected: "false"},
		// (< 1 2) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "<"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}, expected: "true"},
		// (< 2 1) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "<"}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 1}}}, expected: "false"},
		// (< (+ 1 1) 2) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "<"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 2}}}, expected: "false"},
		// (< (+ 1 1) 1) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "<"}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 1}}}, expected: "false"},

		// (>= 1 1) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">="}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, expected: "true"},
		// (>= 1 2) => false
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">="}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 2}}}, expected: "false"},
		// (>= 2 1) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">="}, {Type: ast.Number, Value: 2}, {Type: ast.Number, Value: 1}}}, expected: "true"},
		// (>= (+ 1 1) 2) => true
		{input: ast.Node{Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: ">="}, {Type: ast.List, Children: []*ast.Node{{Type: ast.Symbol, Value: "+"}, {Type: ast.Number, Value: 1}, {Type: ast.Number, Value: 1}}}, {Type: ast.Number, Value: 2}}}, expected: "true"},
		// (>= (+ 1 1) 1) => true

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
		t.Errorf("[%v] expected %v, got %v", input.PrStr(false), expected, result)
	}
}
