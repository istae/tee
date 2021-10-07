package tee

import (
	"bufio"
	"fmt"
	"os"
)

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
