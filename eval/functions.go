package eval

import "golisp/ast"

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
