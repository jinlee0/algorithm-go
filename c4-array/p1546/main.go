package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	N := scanInt()
	arr := make([]int, N)
	var sum, max int
	for _, score := range arr {
		score = scanInt()
		if score > max {
			max = score
		}
		sum += score
	}
	fmt.Println(float64(sum) / float64(N) / float64(max) * 100)
}

func scanInt() int {
	var value int
	fmt.Fscan(reader, &value)
	return value
}
