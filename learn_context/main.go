package main

import (
	"context"
	"fmt"
	"time"
)

func Work(ctx context.Context, name string, i int) {
	if i > 0 {
		go Work(ctx, name, i-1)
	}
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, i, "stopped")
			return
		case <-ticker.C:
			fmt.Println(name, i, "running")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Work(ctx, "work", 3)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
