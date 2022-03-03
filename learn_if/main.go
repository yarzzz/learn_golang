package main

import (
	"fmt"
)

func main() {
	var a []int
	if len(a) == 0 || a[0] == 0 {
		fmt.Println(a)
	}
}
