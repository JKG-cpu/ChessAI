package core

func GetOccupiedSides(board *Board, color Color) (uint64, uint64) {
	var ownOccupied uint64
	var enemyOccupied uint64

	switch color {
	case White:
		ownOccupied = board.WhiteOccupied
		enemyOccupied = board.BlackOccupied
	case Black:
		ownOccupied = board.BlackOccupied
		enemyOccupied = board.WhiteOccupied
	}

	return ownOccupied, enemyOccupied
}

// Gets the attacks for Knight + King
// Doesn't matter because it doesn't keep going in direction (x, y)
func GetSimpleMoves(offsets [][2]int, file int, rank int) uint64 {
	var moves uint64

	for _, offset := range offsets {
		newFile := file + offset[0]
		newRank := rank + offset[1]

		if newFile >= 0 && newFile < 8 && newRank >= 0 && newRank < 8 {
			newSq := newRank * 8 + newFile
			moves |= 1 << newSq
		}
	}

	return moves
}

// Gets the attacks for the rest of the pieces
// Because they can keep going diagonally / straight until blocked
func GetComplexMoves(offsets [][2]int, file int, rank int, occupied uint64, enemyOccupied uint64) uint64 {
	var moves uint64

	for _, offset := range offsets {
		currFile := file
		currRank := rank
		for {
			currFile += offset[0]
			currRank += offset[1]

			if currFile < 0 || currFile > 7 || currRank < 0 || currRank > 7 {
				break
			}

			sq := currRank * 8 + currFile

			// Check for Enemy / Friendly
			if (enemyOccupied>>sq) & 1 == 1 {
				moves |= 1 << sq
				break
			}

			if (occupied>>sq) & 1 == 1 {
				break
			}

			moves |= 1 << sq
		}
	}

	return moves
}

func RookMoves(board *Board, sq int, color Color) uint64 {
	offsets := [][2]int {
			{-1, 0},
		{0, -1}, {0, 1},
			{1, 0},
	}

	file := sq % 8
	rank := sq / 8

	ownOccupied, enemyOccupied := GetOccupiedSides(board, color)

	return GetComplexMoves(offsets, file, rank, ownOccupied, enemyOccupied)
}

func KnightMoves(board *Board, sq int, color Color) uint64 {
	offsets := [][2]int{
		{1, 2}, {2, 1}, {2, -1}, {1, -2},
		{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	}

	file := sq % 8
	rank := sq / 8

	moves := GetSimpleMoves(offsets, file, rank)

	var ownOccupied uint64

	switch color {
	case White:
		ownOccupied = board.WhiteOccupied
	case Black:
		ownOccupied = board.BlackOccupied
	}

	return moves & ^ownOccupied
}

func BishopMoves(board *Board, sq int, color Color) uint64 {
	offsets := [][2]int{
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	file := sq % 8
	rank := sq / 8

	ownOccupied, enemyOccupied := GetOccupiedSides(board, color)

	return GetComplexMoves(offsets, file, rank, ownOccupied, enemyOccupied)
}

func QueenMoves(board *Board, sq int, color Color) uint64 {
	offsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	file := sq % 8
	rank := sq / 8

	ownOccupied, enemyOccupied := GetOccupiedSides(board, color)

	return GetComplexMoves(offsets, file, rank, ownOccupied, enemyOccupied)
}

func KingMoves(board *Board, sq int, color Color) uint64 {
	offsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	file := sq % 8
	rank := sq / 8

	moves := GetSimpleMoves(offsets, file, rank)

	var ownOccupied uint64

	switch color {
	case White:
		ownOccupied = board.WhiteOccupied
	case Black:
		ownOccupied = board.BlackOccupied
	}
	
	return moves & ^ownOccupied
}

func PawnMoves(board *Board, sq int, color Color) uint64 {
	var direction int
	if color == White {
		direction = 1
	} else {
		direction = -1
	}

	file := sq % 8
	rank := sq / 8

	onHomeRank := (color == White && rank == 1) || (color == Black && rank == 6)
	ownOccupied, enemyOccupied := GetOccupiedSides(board, color)
	allOccupied := ownOccupied | enemyOccupied

	var moves uint64

	// One Jump
	singleRank := rank + direction
	singleSq := -1
	if singleRank >= 0 && singleRank <= 7 {
		singleSq = singleRank*8 + file
		if (allOccupied>>singleSq)&1 == 0 {
			moves |= 1 << singleSq

			// Two Jumps
			if onHomeRank {
				doubleRank := rank + 2*direction
				doubleSq := doubleRank*8 + file
				if (allOccupied>>doubleSq)&1 == 0 {
					moves |= 1 << doubleSq
				}
			}
		}
	}

	// Capturing
	captureFiles := []int{file - 1, file + 1}
	for _, cf := range captureFiles {
		if cf < 0 || cf > 7 {
			continue
		}
		captureRank := rank + direction
		if captureRank < 0 || captureRank > 7 {
			continue
		}
		captureSq := captureRank*8 + cf

		isNormalCapture := (enemyOccupied>>captureSq)&1 == 1
		isEnPassantCapture := captureSq == board.enPassantTarget

		if isNormalCapture || isEnPassantCapture {
			moves |= 1 << captureSq
		}
	}

	return moves
}
