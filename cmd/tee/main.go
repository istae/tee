package tee

import (
	"bufio"
	"fmt"
	"os"
)

const test = `
func lol() {
}
x = lol()
`

const test2 = `
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

	r := bufio.NewReader(os.Stdin)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(s)
	}
}
