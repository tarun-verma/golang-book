package echo

import (
	"fmt"
	"os"
	"strings"
)

func argsWithoutJoin() {
	var s string

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}

	fmt.Println(s)
}

func argsWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	argsWithoutJoin()
	argsWithJoin()
}
