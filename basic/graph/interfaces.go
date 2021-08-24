package graph

type GraphType int

const (
	Dense  GraphType = 0
	Sparse GraphType = 1
)

type Graph interface {
	Insert(e *Edge)
	Directed() bool
	Remove(e *Edge)
	Contains(e *Edge) bool
	Scan() error
	Show()
	Vcnt() int
	SearchR(v, w int) bool
}
