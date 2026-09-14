package core

func KnightMoves(board *Board, sq int, color Color) uint64 {
	offsets := [8][2]int{
		{1, 2}, {2, 1}, {2, -1}, {1, -2},
		{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	}

	file := sq % 8
	rank := sq / 8

	var attacks uint64

	for _, offset := range offsets {
		newFile := file + offset[0]
		newRank := rank + offset[1]

		if newFile >= 0 && newFile < 8 && newRank >= 0 && newRank < 8 {
			newSq := newRank * 8 + newFile
			attacks |= 1 << newSq
		}
	}

	var ownOccupied uint64

	switch color {
	case White:
		ownOccupied = board.WhiteOccupied
	case Black:
		ownOccupied = board.BlackOccupied
	}

	return attacks & ^ownOccupied
}