package main

import (
	"fmt"
)

func Numbers(s string) int {
	return len(s)
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Printf("%d", Numbers(s))
}
