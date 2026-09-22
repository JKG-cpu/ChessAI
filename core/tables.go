package core

var KnightTable = [64]int{
	-4, -3, -2, -2, -2, -2, -3, -4,
	-3, -2, 0, 1, 1, 0, -2, -3,
	-2, 1, 2, 3, 3, 2, 1, -2,
	-2, 1, 3, 4, 4, 3, 1, -2,
	-2, 1, 3, 4, 4, 3, 1, -2,
	-2, 1, 2, 3, 3, 2, 1, -2,
	-3, -2, 0, 1, 1, 0, -2, -3,
	-4, -3, -2, -2, -2, -2, -3, -4,
}

func MirrorSquare(sq int, color Color) int {
	if color == Black {
		file := sq % 8
		rank := sq / 8

		return (7 - rank) * 8 + file	
	}
	return sq
}

// Callables
func KnightPositionScore(sq int, color Color) int {
	sq = MirrorSquare(sq, color)
	return KnightTable[sq]
}
