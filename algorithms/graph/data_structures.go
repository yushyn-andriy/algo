package graph

// ========================================================
// Graph data structures                                  =
// ========================================================

// Edge ...
type Edge struct {
	// Start point, end point
	A, B int

	// Distance between A and B
	W float64
}
