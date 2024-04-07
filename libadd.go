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
	var execution_start = time.Now()
	for i := 0; i < 5; i++ {
		go add(1, 2)
		fmt.Println(fmt.Sprintf("Async Iteration %d - %s", i, time.Since(execution_start)))
	}
}

//export slow_add
func slow_add() {
	var execution_start = time.Now()
	for i := 0; i < 5; i++ {
		add(1, 2)
		fmt.Println(fmt.Sprintf("Slow Iteration %d - %s", i, time.Since(execution_start)))
	}
}

func main() {
}
