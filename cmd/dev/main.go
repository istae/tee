package main

import (
	"log"
	"tee/eval"
	"tee/lexer"
	"tee/parser"
)

const test = `
y = 2 
if y > 1 { y = 3 }
x = y * 2.0 + 3.0 / 12.0
`

const test1 = `
x = 100.0 + 5.0 * 2.0 + 3.0 / 12.0
y = 12 - 10 * 2
z = x * y
if x > 10.0 {
	x = 123 * 12
}

func lol() {

}

for x < 1500 {
	x = x + 1
	if x > 1500 {
		x = 0  // this is a comment
	}
}
`

func main() {

	l := lexer.NewLexer()

	t, err := l.Read(test)
	if err != nil {
		log.Fatal(err)
	}

	p := parser.NewParser()

	b, err := p.AST(test, t)
	if err != nil {
		log.Fatal(err)
	}

	_ = eval.NewEval().Eval(b)

	// for _, v := range values {
	// 	v.Print()
	// }
}
