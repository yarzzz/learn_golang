package main

import (
	"log"
	"os"
)

func one(x int, c chan int) {
	c <- x
	// log.Println(x)
}

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go one(i, c)
	}

	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)

	log.Println(os.Args[1])
}
