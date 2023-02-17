package main

import (
	"fmt"
	"runtime"
)

func listThreads() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	runtime.GOMAXPROCS(8)
	fmt.Println(listThreads())
}
