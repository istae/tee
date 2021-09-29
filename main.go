package main

import (
	"log"
	"tee/lexer"
	"tee/parser"
)

const test = `
x = 100.0 + 5.0 * 2.0 + 3.0 / 12.0
y = 12 - 10 * 2
z = x * y
if x > 10.0 {
	x = 123 * 12
}

for x < 1500 {
	x = x + 1
	if x == 1500 {
		x = 0
	}
}
`

const test1 = `
x = 3 + 5 * 2 / 1 + 1 * 2
x = x + 2
x = "asd" + 2
x = x < 2

if x < 2 {

	if y < 3 {
		for x < 4 {
			x = z + 12
			//jkjkj
		}
	}

	x = 12 + 4
}

y = x * x

for x < 4 {

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

	values := parser.NewEval().Eval(b)

	for _, v := range values {
		v.Print()
	}
}
