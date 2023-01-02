package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	sum := 0
	for i := 0; i <= n; i += 3 {
		if n-i < 3 && n-i != 0 {
			sum = -1
			break
		}
		if (n-i)%5 == 0 {
			sum += (n - i) / 5
			break
		}
		sum++
	}
	fmt.Println(sum)
}
