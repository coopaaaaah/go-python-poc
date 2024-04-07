// libadd.go
package main

import "C"

import (
	"fmt"
	"time"
)

//export add
func add(left, right int) int {
	time.Sleep(500 * time.Millisecond)
	return left + right
}

//export async_add
func async_add() {
	for i := 0; i < 5; i++ {
		go add(1, 2)
		fmt.Println(fmt.Sprintf("Async Iteration %d", i))
	}
}

//export slow_add
func slow_add() {
	for i := 0; i < 5; i++ {
		add(1, 2)
		fmt.Println(fmt.Sprintf("Slow Iteration %d", i))
	}
}

func main() {
}
