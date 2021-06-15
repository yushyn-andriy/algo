package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var printf = fmt.Printf

func main() {
	var i, j int32
	var result []string
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Sscanf(input.Text(), "%d %d", &i, &j)
		var a, b int32 = i, j
		var c int32
		if i >= j {
			a = j
			b = i
		} else {
			a = i
			b = j
		}
		for ; a <= b; a++ {
			var counter int32 = 0
			var n int32 = a
			for {
				counter++
				if n == 1 {
					break
				}
				if n%2 == 0 {
					n = n / 2
				} else {
					n = 3*n + 1
				}
			}
			if counter > c {
				c = counter
			}
		}
		result = append(result, fmt.Sprintf("%d %d %d", i, j, c))
	}
	fmt.Println(strings.Join(result, "\n"))
}
