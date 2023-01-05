package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	row := 0
	col := 0
	max := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			scanner.Scan()
			input, _ := strconv.Atoi(scanner.Text())
			if input > max {
				row, col, max = i, j, input
			}
		}
	}

	fmt.Println(max)
	fmt.Println(row+1, col+1)

}
