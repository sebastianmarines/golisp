package eval

import "golisp/ast"

func Eval(astNode *ast.Node, env *Env) *ast.Node {
	return evalAst(astNode, env)
}

func evalAst(astNode *ast.Node, env *Env) *ast.Node {
	switch astNode.Type {
	case ast.Symbol:
		return evalSymbol(astNode, env)
	case ast.List:
		return evalList(astNode, env)
	default:
		return astNode
	}
}

func evalSymbol(astNode *ast.Node, env *Env) *ast.Node {
	value, ok := env.Get(astNode.Value.(string))
	if !ok {
		panic("undefined symbol")
	}
	return &value
}

func evalList(astNode *ast.Node, env *Env) *ast.Node {
	if len(astNode.Children) == 0 {
		return astNode
	}
	first := astNode.Children[0]
	if first.Type == ast.Symbol {
		i, ok := env.Get(first.Value.(string))
		if !ok {
			panic("undefined symbol")
		}
		if i.Type == ast.InternalFunction {
			return evalFunction(astNode, env)
		}
	}
	return evalListChildren(astNode, env)
}

func evalListChildren(astNode *ast.Node, env *Env) *ast.Node {
	var children []*ast.Node
	for _, child := range astNode.Children {
		children = append(children, evalAst(child, env))
	}
	return &ast.Node{Type: ast.List, Children: children}
}

func evalFunction(astNode *ast.Node, env *Env) *ast.Node {
	first := astNode.Children[0]
	f, ok := env.Get(first.Value.(string))
	if !ok {
		panic("undefined symbol")
	}
	if f.Type != ast.InternalFunction {
		panic("not a function")
	}

	if first.Value == "def!" || first.Value == "let*" {
		return f.Value.(func(ast.Node, ast.Node, *Env) *ast.Node)(*astNode.Children[1], *astNode.Children[2], env)
	}

	if first.Value == "do" {
		return f.Value.(func(*Env, ...*ast.Node) *ast.Node)(env, astNode.Children[1:]...)
	}

	if first.Value == "prn" {
		return f.Value.(func(*Env, ...*ast.Node) *ast.Node)(env, astNode.Children[1:]...)
	}

	if first.Value == "=" {
		return f.Value.(func(*Env, ...*ast.Node) *ast.Node)(env, astNode.Children[1:]...)
	}

	var values []interface{}
	for _, child := range astNode.Children[1:] {
		child = evalAst(child, env)
		values = append(values, child.Value)
	}

	result := f.Value.(func(...interface{}) interface{})(values...)

	// Get the node type from the result
	var nodeType ast.NodeType
	switch result.(type) {
	case int:
		nodeType = ast.Number
	case float64:
		nodeType = ast.Number
	case string:
		nodeType = ast.String
	}

	return &ast.Node{Type: nodeType, Value: result}
}
