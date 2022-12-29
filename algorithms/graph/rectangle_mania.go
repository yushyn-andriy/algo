package graph

import "fmt"

type Direction int

const (
	None Direction = iota - 1
	Up
	Down
	Left
	Right
)

// o(n^2) time | o(n^2) space - where n is the number of coordinates
func RectangleMania(coords [][]int) int {
	coordsTable := getCoordsTable(coords)
	return getRectangleCount(coords, coordsTable)
}

type CoordsTable map[string]map[Direction][][]int

func getCoordsTable(coords [][]int) CoordsTable {
	table := CoordsTable{}
	for _, coord1 := range coords {
		directions := map[Direction][][]int{
			Up:    {},
			Right: {},
			Down:  {},
			Left:  {},
		}
		for _, coord2 := range coords {
			coord2Direction := getCoordDirection(coord1, coord2)
			if coord2Direction != None {
				directions[coord2Direction] = append(directions[coord2Direction], coord2)
			}
		}
		table[coordToString(coord1)] = directions
	}
	return table
}

func getCoordDirection(coord1, coord2 []int) Direction {
	if coord2[1] == coord1[1] {
		if coord2[0] > coord1[0] {
			return Right
		} else if coord2[0] < coord1[0] {
			return Left
		}
	} else if coord2[0] == coord1[0] {
		if coord2[1] > coord1[1] {
			return Up
		} else if coord2[1] < coord1[1] {
			return Down
		}
	}
	return None
}

func getRectangleCount(coords [][]int, coordsTable CoordsTable) int {
	count := 0
	for _, coord := range coords {
		count += clockWiseCountRectangles(coord, coordsTable, Up, coord)
	}
	return count
}

func clockWiseCountRectangles(coord []int, coordsTable CoordsTable, direction Direction, origin []int) int {
	if direction == Left {
		for _, element := range coordsTable[coordToString(coord)][Left] {
			if element[0] == origin[0] && element[1] == origin[1] {
				return 1
			}
		}
		return 0
	}
	rectangleCount := 0
	nextDirection := direction.NextClockwise()
	for _, nextCoord := range coordsTable[coordToString(coord)][direction] {
		rectangleCount += clockWiseCountRectangles(nextCoord, coordsTable, nextDirection, origin)
	}
	return rectangleCount
}

func (d Direction) NextClockwise() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	return None
}

func coordToString(coord []int) string {
	return fmt.Sprintf("%d-%d", coord[0], coord[1])
}
