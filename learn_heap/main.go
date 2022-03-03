package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type myQueue []int

func (q myQueue) Len() int {
	return len(q)
}

func (q myQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q myQueue) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q *myQueue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *myQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

var count = 10000000

func useSort() {
	rand.Seed(0)
	q := &myQueue{}
	start := time.Now()
	for i := 0; i < count; i++ {
		q.Push(rand.Int())
		sort.Sort(q)
	}
	for len(*q) > 0 {
		*q = (*q)[1:]
		sort.Sort(q)
	}
	fmt.Println(time.Since(start))
}

func useHeap() {
	rand.Seed(0)
	q := &myQueue{}
	last := 0
	start := time.Now()
	for i := 0; i < count; i++ {
		heap.Push(q, rand.Int())
	}
	for q.Len() > 0 {
		t := (*q)[0]
		x := heap.Pop(q).(int)
		if x != t {
			fmt.Println(t, x)
		}
		if x < last {
			fmt.Println(x)
		}
		last = x
	}
	fmt.Println(time.Since(start))
}

func main() {
	useHeap()
	// useSort()
}
