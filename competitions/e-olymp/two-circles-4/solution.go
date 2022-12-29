package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

const numbers = 6

var printf = fmt.Printf
var debug = flag.Bool("d", false, "debug mode")

type circle struct {
	x float64
	y float64
	r float64
}

func Norma(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(
		math.Pow(x2-x1, 2) +
			math.Pow(y2-y1, 2),
	)
}

func CountIntersection(c1 circle, c2 circle) int {
	n := Norma(c1.x, c2.x, c1.y, c2.y)
	sumr := c1.r + c2.r
	s := ""
	res := 0

	switch true {
	case n > sumr:
		s = "n > sumr"
		res = 0
	case n == sumr:
		s = "n == sumr"
		res = 1
	case n == 0.0 && c1.r == c2.r:
		s = "n == 0.0 && c1.r == c2.r"
		res = -1
	case n == 0.0 && c1.r != c2.r:
		s = "n == 0.0 && c1.r != c2.r"
		res = 0
	case (c1.r-n == c2.r) || (c2.r-n == c1.r):
		s = "(c1.r - n == c2.r) || (c2.r - n == c1.r)"
		res = 1
	/*case (n - 2 * c1.r > 0) || (n - 2 * c2.r > 0):
	printf("(n - 2 * c1.r > 0) || (n - 2 * c2.r > 0)\n")
	return 0*/
	case n < c1.r && c1.r+n < c2.r:
		s = "n < c1.r && c1.r + n < c2.r"
		res = 0

	case n < c2.r && c2.r > c1.r+n:
		s = "n < c2.r && c2.r > c1.r + n"
		res = 0
	case n < c1.r && c1.r > c2.r+n:
		s = "n < c2.r && c2.r > c1.r + n"
		res = 0

	case n < c1.r && c1.r+n > c2.r:
		s = "n < c1.r && c1.r + n > c2.r"
		res = 2
	case n < c2.r && c2.r+n > c1.r:
		s = "n < c2.r && c2.r + n > c1.r"
		res = 2
	case n < sumr && (n >= c1.r && n >= c2.r):
		s = "n < sumr && (n >= c1.r && n >= c2.r)"
		res = 2
	}

	if *debug {
		printf("%f %f %f %s\n", n, c1.r, c2.r, s)
	}
	return res
}

func main() {
	flag.Parse()
	var c1, c2 circle

	_, err := fmt.Scanf(strings.TrimRight(strings.Repeat("%f ", numbers), " "),
		&c1.x, &c1.y, &c1.r,
		&c2.x, &c2.y, &c2.r,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "circles: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(CountIntersection(c1, c2))
}
