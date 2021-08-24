package graph

import (
	"bufio"
	"fmt"
	"os"
)

// SparseGraph...
type SparseGraph struct {
	v          int
	ecnt, vcnt int
	directed   bool
	adj        [][]int
}

func (g *SparseGraph) SearchR(v, w int) bool {
	return g.searchR(v, w, make(map[int]bool))
}

func (g *SparseGraph) searchR(v, w int, visited map[int]bool) bool {
	if v == w {
		return true
	}

	visited[v] = true
	for i := range g.adj[v] {
		neighbor := g.adj[v][i]
		if visited[neighbor] {
			continue
		}
		return g.searchR(neighbor, w, visited)
	}

	return false
}

func (g *SparseGraph) Directed() bool {
	return g.directed
}

func (g *SparseGraph) Insert(e *Edge) {
	g.adj[e.V] = append(g.adj[e.V], e.W)
	if !g.directed {
		g.adj[e.W] = append(g.adj[e.W], e.V)
	}
	g.ecnt++
}

func (g *SparseGraph) Remove(e *Edge) {}

func (g *SparseGraph) Contains(e *Edge) bool {
	return false
}

func (g *SparseGraph) Show() {
	for i := range g.adj {
		fmt.Printf("%d :", i)
		for j := range g.adj[i] {
			fmt.Printf(" %d", g.adj[i][j])
		}
		fmt.Println()
	}
}

func (g *SparseGraph) Ecnt() int {
	return g.ecnt
}

func (g *SparseGraph) Vcnt() int {
	return g.vcnt
}

func (g *SparseGraph) Scan() error {
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

func (g *SparseGraph) Count() int { return -1 }

func (g *SparseGraph) Connected() bool { return true }
