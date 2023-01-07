package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 5)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c <- i
		}()
	}
	wg.Wait()
	close(c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	for v := range c {
		fmt.Println(v)
	}
}
