// Minimax will always be playing as Black on the chess board
package minimax

import (
	"math"

	"github.com/JKG-cpu/ChessAI/core"
)

func MiniMax(board *core.Board, depth int, maximizingPlayer bool) float64 {
	if depth == 0 {
		return core.GetScore(board)
	}

	color := core.Black
	if maximizingPlayer {
		color = core.White
	}

	legalMoves := core.GenerateAllLegalMoves(board, color)

	if len(legalMoves) == 0 {
		if core.IsKingInCheck(board, color) {
			if maximizingPlayer { 
				return math.Inf(-1)
			}
			return math.Inf(1)
		}
		return 0
	}

	if maximizingPlayer {
		best := math.Inf(-1)
		for _, move := range legalMoves {
			board.Move(move)
			score := MiniMax(board, depth-1, false)
			board.UndoMove(move)

			if score > best {
				best = score
			}
		}
		return best
	}

	best := math.Inf(1)
	for _, move := range legalMoves {
		board.Move(move)
		score := MiniMax(board, depth-1, true)
		board.UndoMove(move)

		if score < best {
			best = score
		}
	}
	return best
}

func GetBestMove(board *core.Board, depth int, color core.Color) core.Move {
	var bestMove core.Move
	bestScore := math.Inf(1)

	isMaximizing := true
	if color == core.White {
		isMaximizing = false
	}

	for _, move := range core.GenerateAllLegalMoves(board, color) {
		board.Move(move)

		score := MiniMax(board, depth - 1, isMaximizing)

		if score < bestScore {
			bestScore = score
			bestMove = move
		}

		board.UndoMove(move)
	}

	return bestMove
}