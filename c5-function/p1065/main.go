package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 고루틴 안쓰는 게 훨 빠르다.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := scanner.Text()
	solutionNoGoroutine(n)
	//solutionMonoChannel(n)
}

func solutionNoGoroutine(n string) {
	num, _ := strconv.Atoi(n)
	sum := 0
	for i := 1; i <= num; i++ {
		if isHansooNoGoroutine(i) {
			sum++
		}
	}
	fmt.Println(sum)
}

func isHansooNoGoroutine(n int) bool {
	strN := strconv.Itoa(n)
	if len(strN) == 1 {
		return true
	}
	nums := []rune(strN)
	prev := nums[0]
	diff := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-prev != diff {
			return false
		}
		prev = nums[i]
	}
	return true
}
