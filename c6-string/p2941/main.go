package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cAlphabet = [...]string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	sum := 0
	for _, c := range cAlphabet {
		sum += strings.Count(input, c)
		input = strings.Replace(input, c, " ", -1)
	}
	input = strings.Replace(input, " ", "", -1)
	fmt.Println(sum + len(input))
}
