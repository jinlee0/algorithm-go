package main

import (
	"fmt"
	"strconv"
)

func solutionMonoChannel(n string) {
	c := make(chan bool)
	num, _ := strconv.Atoi(n)
	for i := 1; i <= num; i++ {
		go isHansooMonoChannel(i, c)
	}
	sum := 0
	for i := 1; i <= num; i++ {
		if <-c {
			sum++
		}
	}
	fmt.Println(sum)
}

func isHansooMonoChannel(n int, c chan bool) {
	strN := strconv.Itoa(n)
	if len(strN) == 1 {
		c <- true
		return
	}
	nums := []rune(strN)
	prev := nums[0]
	diff := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-prev != diff {
			c <- false
			return
		}
		prev = nums[i]
	}
	c <- true
}
