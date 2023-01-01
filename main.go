package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"golisp/eval"
	"golisp/reader"
)

func Print(str string) string {
	return str
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
		fmt.Println(Rep(line, env))
	}
}
