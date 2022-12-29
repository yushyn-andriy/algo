package main

import (
	"fmt"

	"github.com/baybaraandrey/algo/basic/graph"
)

func main() {
	fmt.Println("Dense")
	g, _ := graph.NewGraph(10, false, graph.Dense)
	graph.RandGraph(g, 5)
	g.Show()

	fmt.Println("Sparse")
	g1, _ := graph.NewGraph(10, false, graph.Sparse)
	graph.RandGraph(g1, 15)
	g1.Show()
	fmt.Println(g1.SearchR(0, 5))
	fmt.Println(g1.SearchR(2, 8))
}
