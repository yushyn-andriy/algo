package zigzagconversion

import (
	"sort"
	"strings"
)

type Direction uint8
type Key [2]uint16

const (
	Down Direction = iota
	RightUp
)

type Directions struct {
	arr     []Direction
	current int
}

func (d *Directions) Reset() {
	d.current = 0
}

func (d *Directions) Current() Direction {
	return d.arr[d.current]
}

func (d *Directions) Next() Direction {
	d.current += 1
	if d.current >= len(d.arr) {
		d.current = 0
	}
	return d.arr[d.current]
}

func sortedKeys(keys []Key) {
	sort.Slice(keys, func(i, j int) bool {
		k1, k2 := keys[i], keys[j]
		if k1[0] == k2[0] {
			return k1[1] < k2[1]
		}
		return k1[0] < k2[0]
	})
}

// O(nlog(n)) time | O(n) space
// O(n) - map construction
// O(n) - key slice construction
// O(nlog(n)) - quicksorting of keys using standard lib
// O(n) - write result using builder
// where n is the length of the string and m is the numRows
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	directions := Directions{[]Direction{
		Down,
		RightUp,
	}, 0}
	lettersMap := make(map[Key]rune, len(s))

	// O(n)
	var row, coll uint16
	var currentDirection Direction
	for i, ch := range s {
		lettersMap[Key{row, coll}] = ch

		currentDirection = directions.Current()
		if currentDirection == Down {
			row += 1
		} else if currentDirection == RightUp {
			row -= 1
			coll += 1
		}

		if (i+1)%(numRows-1) == 0 {
			directions.Next()
		}
	}

	// O(n)
	keys := make([]Key, 0, len(lettersMap))
	for k := range lettersMap {
		keys = append(keys, k)
	}

	// O(nlog(n))
	sortedKeys(keys)

	builder := strings.Builder{}
	builder.Grow(len(s))

	// O(n)
	for _, k := range keys {
		builder.WriteRune(lettersMap[k])
	}
	return builder.String()
}
