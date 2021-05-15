package trees

//  Tree ...
type Tree struct {
	Key, Value int
	Parent     *Tree
	Left       *Tree
	Right      *Tree
}
