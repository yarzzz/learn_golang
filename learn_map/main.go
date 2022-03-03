package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := map[int]int{}
	a[-1] = -1
	b := a
	b[-2] = -2
	fmt.Println(a, b)
	for i := 0; i < 50; i++ {
		x := i
		wg.Add(1)
		go func() {
			b[x] = x
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < 50; i++ {
		x := i
		wg.Add(1)
		go func() {
			fmt.Println(x, b[x])
			wg.Done()
		}()
	}
	wg.Wait()

}
