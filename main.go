package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"golisp/reader"
)

func Eval(ast string) string {
	return ast
}

func Print(str string) string {
	return str
}

func Rep(str string) string {
	newReader := reader.Read(str)
	return Print(Eval(newReader.String()))
}

func main() {
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
		fmt.Println(Rep(line))
	}
}
