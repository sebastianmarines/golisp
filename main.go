package main

import (
	"fmt"
	"github.com/chzyer/readline"
)

func Read(str string) string {
	return str
}

func Eval(ast string) string {
	return ast
}

func Print(str string) string {
	return str
}

func Rep(str string) string {
	return Print(Eval(Read(str)))
}

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		fmt.Println(Rep(line))
	}
}
