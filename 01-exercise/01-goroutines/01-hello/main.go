package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("goroutine-1")
	// goroutine with anonymous function
	go func() {
		fun("Goroutine2")
	}()
	// goroutine with function value call
	hello := fun
	go hello("goroutine3")
	// wait for goroutines to end
	time.Sleep(200 * time.Millisecond)
	fmt.Println("done..")
}
