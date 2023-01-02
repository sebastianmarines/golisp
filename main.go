package main

import (
	_ "embed"
	"fmt"
	"github.com/chzyer/readline"
	"golisp/ast"
	"golisp/eval"
	"golisp/lexer"
	"golisp/reader"
	"os"
)

//go:embed core.golisp
var core []byte

func Print(n *ast.Node) string {
	return n.PrStr(false)
}

func Rep(str string, env *eval.Env) string {
	newReader := reader.Read(str)
	return Print(eval.Eval(newReader, env))
}

func Repl(env *eval.Env) {
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

func evalMultiline(str string, env *eval.Env) {
	tokens := lexer.Tokenize(str)
	tokenReader := reader.NewReader(tokens)

	for node := tokenReader.Read(); node != nil; node = tokenReader.Read() {
		eval.Eval(node, env)
	}
}

func main() {
	env := eval.NewEnv(nil)

	evalMultiline(string(core), env)

	if len(os.Args) > 1 {
		file, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		fileString := string(file)
		evalMultiline(fileString, env)

	} else {
		Repl(env)
	}
}
