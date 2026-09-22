// Minimax will always be playing as Black on the chess board
package minimax

import (
	"math"

	"github.com/JKG-cpu/ChessAI/core"
)

func MiniMax(board *core.Board, depth int, maximizingPlayer bool, alpha float64, beta float64) float64 {
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
			score := MiniMax(board, depth-1, false, alpha, beta)
			board.UndoMove(move)

			if score > best {
				best = score
			}
			if best > alpha {
				alpha = best
			}
			if alpha >= beta {
				break
			}
		}
		return best
	}

	best := math.Inf(1)
	for _, move := range legalMoves {
		board.Move(move)
		score := MiniMax(board, depth-1, true, alpha, beta)
		board.UndoMove(move)

		if score < best {
			best = score
		}
		if best < beta {
			beta = best
		}
		if alpha >= beta {
			break
		}
	}
	return best
}

func GetBestMove(board *core.Board, depth int, color core.Color) core.Move {
	legalMoves := core.GenerateAllLegalMoves(board, color)

	if len(legalMoves) == 0 {
		return core.Move{}
	}
	
	var bestMove core.Move
	alpha := math.Inf(-1)
	beta := math.Inf(1)

	if color == core.White {
		bestScore := math.Inf(-1)

		isMaximizing := false

		for _, move := range core.GenerateAllLegalMoves(board, color) {
			board.Move(move)

			score := MiniMax(board, depth - 1, isMaximizing, alpha, beta)
			
			board.UndoMove(move)

			if score > bestScore {
				bestScore = score
				bestMove = move
			}

			if bestScore > alpha {
				alpha = bestScore
			}
		}
	} else {
		bestScore := math.Inf(1)

		isMaximizing := true

		for _, move := range core.GenerateAllLegalMoves(board, color) {
			board.Move(move)

			score := MiniMax(board, depth - 1, isMaximizing, alpha, beta)
			
			board.UndoMove(move)

			if score < bestScore {
				bestScore = score
				bestMove = move
			}

			if bestScore < beta {
				beta = bestScore
			}
		}
	}

	return bestMove
}