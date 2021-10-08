package main

import (
	"log"
	"tee/eval"
	"tee/lexer"
	"tee/parser"
)

const test = `
y = 1.12
func lol(a, b) {
	func x(h,j) {
		if h / 2 > "1" {
			print(h/2, j, "\n")
			y = y*h * j
		}	
	}
	x(a,b)
}

lol(3,4)

x = y * 2

i = 0
for 1 > 0 {
	i = i + 1
	print("hello\n")
	if i > 100 {
		break
	}
}
print(y)
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
