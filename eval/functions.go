package eval

import (
	"fmt"
	"github.com/chzyer/readline"
	"golisp/ast"
)

func plus(values ...interface{}) interface{} {
	result := values[0]
	for _, value := range values[1:] {
		switch value.(type) {
		case int:
			result = result.(int) + value.(int)
		case float64:
			result = result.(float64) + value.(float64)
		case string:
			result = result.(string) + value.(string)
		}
	}
	return result
}

func minus(values ...interface{}) interface{} {
	result := values[0]
	for _, value := range values[1:] {
		switch value.(type) {
		case int:
			result = result.(int) - value.(int)
		case float64:
			result = result.(float64) - value.(float64)
		default:
			panic("invalid type")
		}
	}
	return result
}

func multiply(values ...interface{}) interface{} {
	result := values[0]
	for _, value := range values[1:] {
		switch value.(type) {
		case int:
			result = result.(int) * value.(int)
		case float64:
			result = result.(float64) * value.(float64)
		default:
			panic("invalid type")
		}
	}
	return result
}

func divide(values ...interface{}) interface{} {
	result := values[0]
	for _, value := range values[1:] {
		switch value.(type) {
		case int:
			result = result.(int) / value.(int)
		case float64:
			result = result.(float64) / value.(float64)
		default:
			panic("invalid type")
		}
	}
	return result
}

func define(name ast.Node, value ast.Node, env *Env) *ast.Node {
	newValue := evalAst(&value, env)
	env.Set(name.Value.(string), *newValue)
	return newValue
}

func let(values ast.Node, rest ast.Node, outer *Env) *ast.Node {
	env := NewEnv(outer)
	for i := 0; i < len(values.Children); i += 2 {
		env.Set(values.Children[i].Value.(string), *evalAst(values.Children[i+1], env))
	}
	return evalAst(&rest, env)
}

func do(outer *Env, rest ...*ast.Node) *ast.Node {
	var result *ast.Node
	for _, node := range rest {
		result = evalAst(node, outer)
	}
	return result
}

func prn(outer *Env, nodes ...*ast.Node) *ast.Node {
	for _, n := range nodes {
		result := evalAst(n, outer)
		_, err := fmt.Fprintf(readline.Stdout, "%v\n", result.Value)
		if err != nil {
			panic(err)
		}
	}
	return &ast.Node{Type: ast.Nil}
}

func equals(outer *Env, values ...*ast.Node) *ast.Node {
	first := evalAst(values[0], outer)
	for _, value := range values[1:] {
		second := evalAst(value, outer)
		if first.Type != second.Type {
			return &ast.Node{Type: ast.False}
		}
		switch first.Type {
		case ast.Number:
			if first.Value.(int) != second.Value.(int) {
				return &ast.Node{Type: ast.False}
			}
		case ast.String:
			if first.Value.(string) != second.Value.(string) {
				return &ast.Node{Type: ast.False}
			}
		case ast.Nil:
			if first.Value != second.Value {
				return &ast.Node{Type: ast.False}
			}
		case ast.List:
			if len(first.Children) != len(second.Children) {
				return &ast.Node{Type: ast.False}
			}
			for i := 0; i < len(first.Children); i++ {
				if first.Children[i].Value != second.Children[i].Value {
					return &ast.Node{Type: ast.False}
				}
			}
		default:
			panic("invalid type")
		}
	}
	return &ast.Node{Type: ast.True}
}

func greaterThan(outer *Env, values ...*ast.Node) *ast.Node {
	first := evalAst(values[0], outer)
	for _, value := range values[1:] {
		second := evalAst(value, outer)
		if first.Type != second.Type {
			panic("invalid type")
		}
		switch first.Type {
		case ast.Number:
			if first.Value.(int) <= second.Value.(int) {
				return &ast.Node{Type: ast.False}
			}
		case ast.String:
			if first.Value.(string) <= second.Value.(string) {
				return &ast.Node{Type: ast.False}
			}
		default:
			panic("invalid type")
		}
	}
	return &ast.Node{Type: ast.True}
}

func lessThan(outer *Env, values ...*ast.Node) *ast.Node {
	first := evalAst(values[0], outer)
	for _, value := range values[1:] {
		second := evalAst(value, outer)
		if first.Type != second.Type {
			panic("invalid type")
		}
		switch first.Type {
		case ast.Number:
			if first.Value.(int) >= second.Value.(int) {
				return &ast.Node{Type: ast.False}
			}
		case ast.String:
			if first.Value.(string) >= second.Value.(string) {
				return &ast.Node{Type: ast.False}
			}
		default:
			panic("invalid type")
		}
	}
	return &ast.Node{Type: ast.True}
}
