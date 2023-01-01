package main

import (
	. "fmt"
)

func main() {
	var H, M int
	Scan(&H, &M)
	M -= 45
	if M < 0 {
		H--
		M += 60
	}
	if H < 0 {
		H += 24
	}
	Print(H, M)
}
