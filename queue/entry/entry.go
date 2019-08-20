package main

import (
	"fmt"
	"test/queue"
)

func main() {
	q := queue.Queue{1}
	fmt.Println(q)
	q.Push(2)
	fmt.Println(q)
	q.Push(3)
	fmt.Println(q)

	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)

	fmt.Println()

	fmt.Println(q.IsEmpty())

}
