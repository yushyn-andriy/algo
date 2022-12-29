package graph

import (
	"bufio"
	"fmt"
	"os"
)

type DenseGraph struct {
	v          int
	ecnt, vcnt int
	directed   bool
	adj        [][]bool
}

type Edge struct {
	V, W int
}

func (g *DenseGraph) SearchR(v, w int) bool {
	return false
}

func (g *DenseGraph) V() int {
	return -1
}

func (g *DenseGraph) E() int {
	return -1
}

func (g *DenseGraph) Directed() bool {
	return g.directed
}

func (g *DenseGraph) Insert(e *Edge) {
	if !g.adj[e.V][e.W] {
		g.ecnt++
	}
	g.adj[e.V][e.W] = true
	if !g.directed {
		g.adj[e.W][e.V] = true
	}
}

func (g *DenseGraph) Remove(e *Edge) {
	if g.adj[e.V][e.W] {
		g.ecnt--
	}
	g.adj[e.V][e.W] = false
	if g.directed {
		g.adj[e.W][e.V] = false
	}
}

func (g *DenseGraph) Contains(e *Edge) bool {
	return g.adj[e.V][e.W]
}

func (g *DenseGraph) Show() {
	for i := range g.adj {
		fmt.Printf("%d :", i)
		for j := range g.adj[i] {
			if g.adj[i][j] {
				fmt.Printf(" %d", j)
			}
		}
		fmt.Println()
	}
}

func (g *DenseGraph) ShowAdj() {
	fmt.Print("____")
	for i := range g.adj {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	for i := range g.adj {
		fmt.Printf("%2d|", i)
		for j := range g.adj[i] {
			if g.adj[i][j] {
				fmt.Printf(" %2d", 1)
			} else {
				fmt.Printf(" %2d", 0)
			}
		}
		fmt.Println()
	}
}

func (g *DenseGraph) Ecnt() int {
	return g.ecnt
}

func (g *DenseGraph) Vcnt() int {
	return g.vcnt
}

func (g *DenseGraph) Scan() error {
	scanner := bufio.NewScanner(os.Stdin)
	v, w := 0, 0
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &v, &w)
		if err != nil {
			return err
		}
		g.Insert(&Edge{v, w})
	}
	return nil
}

func (g *DenseGraph) Count() int { return -1 }

func (g *DenseGraph) Connected() bool { return true }
