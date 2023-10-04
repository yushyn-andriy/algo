package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("*", rand.Intn(10) + 1))
}
