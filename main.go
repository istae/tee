package main

import (
	"log"
	"tee/lexer"
	"tee/parser"
)

const test = `
x = 3 + 5 * 2 / 1 + 1 * 2
x = x + 2
x = "asd" + 2
x = x < 2

if x < 2 {
	x = 12 + 4
}

y = x * x

`

func main() {

	l := lexer.NewLexer()

	t, err := l.Read(test)
	if err != nil {
		log.Fatal(err)
	}

	p := parser.NewParser()

	err = p.AST(test, t)
	if err != nil {
		log.Fatal(err)
	}
}
