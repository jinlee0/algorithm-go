package main

import "fmt"

func main() {
	m := map[int]int{}
	m[1] = 1
	m[2]++
	fmt.Print(m[1])
	fmt.Print(m[2])
}
