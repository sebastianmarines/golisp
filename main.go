package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"golisp/ast"
	"golisp/eval"
	"golisp/reader"
)

func Print(n *ast.Node) string {
	return n.PrStr(false)
}

func Rep(str string, env *eval.Env) string {
	newReader := reader.Read(str)
	return Print(eval.Eval(newReader, env))
}

func main() {
	env := eval.NewEnv(nil)
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer func(rl *readline.Instance) {
		err := rl.Close()
		if err != nil {

		}
	}(rl)

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		res := Rep(line, env)

		_, err = fmt.Fprintln(rl.Stdout(), res)
		if err != nil {
			return
		}
	}
}
