package eval

import "golisp/ast"

type Env struct {
	outer *Env
	store map[string]ast.Node
}

func NewEnv(outer *Env) *Env {
	env := &Env{outer: outer, store: make(map[string]ast.Node)}
	env.Set("+", ast.Node{Type: ast.Function, Value: plus})
	env.Set("-", ast.Node{Type: ast.Function, Value: minus})
	env.Set("*", ast.Node{Type: ast.Function, Value: multiply})
	env.Set("/", ast.Node{Type: ast.Function, Value: divide})
	env.Set("def!", ast.Node{Type: ast.Function, Value: define})
	return env
}

func (e *Env) Set(key string, value ast.Node) {
	e.store[key] = value
}

func (e *Env) Get(key string) (ast.Node, bool) {
	value, ok := e.store[key]
	if !ok && e.outer != nil {
		return e.outer.Get(key)
	}
	return value, ok
}
