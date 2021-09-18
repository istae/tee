package main

import (
	"log"
	"tee/lexer"
	"tee/parser"
)

const test = `
z = 12 / 3.0
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
