package main

import (
	"fmt"
	"graph_theory/src/data_structures"
)

func main() {
	fmt.Println("Stack")
	var stack = new(data_structures.ArrStack)
	fmt.Println(stack)
	fmt.Println(stack.IsEmpty())
	for i := 0; i < 20; i++ {
		stack.Push(i)
	}
	fmt.Println(stack)
	fmt.Println(stack.IsEmpty())
	for i := 0; i < 20; i++ {
		_, _ = stack.Pop()
		if i == 17 {
			fmt.Println(stack)
		}
	}
	fmt.Println(stack)
	fmt.Println(stack.IsEmpty())

	fmt.Println("Queue")
	var queue = new(data_structures.ArrQueue)
	fmt.Println(queue)
	fmt.Println(queue.IsEmpty())
	for i := 0; i < 20; i++ {
		queue.Enqueue(i)
	}
	fmt.Println(queue)
	fmt.Println(queue.IsEmpty())
	for i := 0; i < 20; i++ {
		_, _ = queue.Dequeue()
		if i == 17 {
			fmt.Println(queue)
		}
	}
	fmt.Println(queue)
	fmt.Println(queue.IsEmpty())
}
