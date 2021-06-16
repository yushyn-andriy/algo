package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vaucher struct {
	Days, Price int
}

func Permutations(x []Vaucher) int {
	p := [][]Vaucher{}
	chosen := make([]bool, len(x))
	max := 0
	permutations(x, []Vaucher{}, &p, chosen, &max)
	return max
}

func permutations(x, c []Vaucher, r *[][]Vaucher, chosen []bool, max *int) {
	if len(x) == len(c) {
		curr := 0
		for i, v := range c {
			if v.Days-i > 0 {
				curr += (v.Days - i) * v.Price
			}
		}
		if curr > *max {
			*max = curr
		}
		return
	}

	for i := 0; i < len(x); i++ {
		if chosen[i] {
			continue
		}
		c = append(c, x[i])
		chosen[i] = true
		permutations(x, c, r, chosen, max)
		c = c[:len(c)-1]
		chosen[i] = false
	}
}

func main() {
	var countLines, line int
	var vauchers = []Vaucher{}

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		if line == 0 {
			fmt.Sscanf(in.Text(), "%d", &countLines)
		} else {
			v := Vaucher{}
			fmt.Sscanf(in.Text(), "%d %d", &v.Days, &v.Price)
			vauchers = append(vauchers, v)
		}
		line++
	}
	fmt.Println(Permutations(vauchers))
}
