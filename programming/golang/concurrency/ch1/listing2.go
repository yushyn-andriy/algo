package main

import (
	"fmt"
)

func showNumber(id, i int) {
	fmt.Println(id, i)
}

func main() {
	for i := 0; i < 10; i++ {
		go showNumber(i, i)
	}
	// runtime.GOMAXPROCS(2)
	// runtime.Gosched()

	fmt.Println("bye")
}
