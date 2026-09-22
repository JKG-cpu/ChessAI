package core

import (
	"fmt"
)

// Piece Conversion
func ConvertPieceToPieceType(piece Piece) (Color, PieceType) {
	var color Color
	base := int(piece)

	if piece >= BlackPawn {
		color = Black
		base -= 6
	} else {
		color = White
	}

	pieceType := PieceType(base - 1)

	return color, pieceType
}

func ToSquarePiece(color Color, piece PieceType) Piece {
	base := 1 + int(piece)
	if color == Black {
		base += 6
	}
	return Piece(base)
}

// Move Generation
func GetSpecificPieceMove(board *Board, pieceType PieceType, sq int, color Color) uint64 {
	switch pieceType {
	case Pawn:
		return PawnMoves(board, sq, color)
	case Rook:
		return RookMoves(board, sq, color)
	case Knight:
		return KnightMoves(board, sq, color)
	case Bishop:
		return BishopMoves(board, sq, color)
	case Queen:
		return QueenMoves(board, sq, color)
	case King:
		return KingMoves(board, sq, color)
	}
	return 0
}

func GenerateAllMoves(board *Board, color Color) []Move {
	var moves []Move

	for sq := range 64 {
		piece := board.Squares[sq]
	
		if piece == Empty { continue }
	
		pieceColor, pieceType := ConvertPieceToPieceType(piece)
		if pieceColor != color { continue }

		pieceMoves := GetSpecificPieceMove(
			board, pieceType, sq, color,
		)

		for destSq := range 64 {
			isEnPassant := pieceType == Pawn && destSq == board.enPassantTarget
			
			captured := board.Squares[destSq]

			if isEnPassant {
				enemyColor := White
				if pieceColor == White {
					enemyColor = Black
				}
				captured = ToSquarePiece(enemyColor, Pawn)
			}
			
			if (pieceMoves>>destSq) & 1 == 1 {
				m := Move{
					From: sq,
					To: destSq,
					Piece: piece,
					Captured: captured,
					PrevEnPassantCapture: board.enPassantTarget,
					isEnPassant: isEnPassant,

					PrevWhiteCanCastleKingSide:  board.WhiteCanCastleKingSide,
					PrevWhiteCanCastleQueenSide: board.WhiteCanCastleQueenSide,
					PrevBlackCanCastleKingSide:  board.BlackCanCastleKingSide,
					PrevBlackCanCastleQueenSide: board.BlackCanCastleQueenSide,
				}
				moves = append(moves, m)
			}
		}
	}


	// Castling
	kingSq := 4
	if color == Black {
		kingSq = 60
	}

	castleMoves := CastlingMoves(board, color)

	for destSq := range 64 {
		if (castleMoves>>destSq)&1 == 1 {
			m := Move{
				From:     kingSq,
				To:       destSq,
				Piece:    ToSquarePiece(color, King),
				Captured: Empty,
				isCastle: true,

				PrevEnPassantCapture:       board.enPassantTarget,
				PrevWhiteCanCastleKingSide:  board.WhiteCanCastleKingSide,
				PrevWhiteCanCastleQueenSide: board.WhiteCanCastleQueenSide,
				PrevBlackCanCastleKingSide:  board.BlackCanCastleKingSide,
				PrevBlackCanCastleQueenSide: board.BlackCanCastleQueenSide,
			}
			moves = append(moves, m)
		}
	}

	return moves
}

func GenerateAllLegalMoves(board *Board, color Color) []Move {
	var moves []Move
	allMoves := GenerateAllMoves(board, color)

	for _, m := range allMoves {
		board.Move(m)

		if !IsKingInCheck(board, color) {
			moves = append(moves, m)
		}

		board.UndoMove(m)
	}
	
	return moves
}

func GetLegalMovesForSquare(board *Board, sq int, color Color) []Move {
	allLegal := GenerateAllLegalMoves(board, color)

	var result []Move

	for _, m := range allLegal {
		if m.From == sq {
			result = append(result, m)
		}
	}

	return result
}

// Checks
func IsGameOver(board *Board, color Color) (bool, string) {
	legalMoves := GenerateAllLegalMoves(board, color)

	if len(legalMoves) > 0 {
		return false, ""
	}

	if IsKingInCheck(board, color) {
		return true, "Checkmate"
	}

	return true, "Stalemate"
}

func IsSquareAttacked(board *Board, sq int, byColor Color) bool {
	for i := range 64 {
		curPiece := board.Squares[i]

		if curPiece == Empty { continue }
		
		pieceColor, pieceType := ConvertPieceToPieceType(curPiece)

		if pieceColor != byColor {
			continue
		}

		moves := GetSpecificPieceMove(board, pieceType, i, byColor)

		if (moves>>sq) & 1 == 1 {
			return true
		}
	}

	return false
}

func IsKingInCheck(board *Board, color Color) bool {
    kingPiece := ToSquarePiece(color, King)

    var kingSq int
    for i := range 64 {
        if board.Squares[i] == kingPiece {
            kingSq = i
            break
        }
    }

    enemyColor := White
    if color == White {
        enemyColor = Black
    }

    return IsSquareAttacked(board, kingSq, enemyColor)
}

// Board Generation
func RecomputeOccupied(board *Board) {
	board.WhiteOccupied = board.Pieces[White][Rook] |
		board.Pieces[White][Knight] |
		board.Pieces[White][Bishop] |
		board.Pieces[White][Queen] |
		board.Pieces[White][King] |
		board.Pieces[White][Pawn]

	board.BlackOccupied = board.Pieces[Black][Rook] |
		board.Pieces[Black][Knight] |
		board.Pieces[Black][Bishop] |
		board.Pieces[Black][Queen] |
		board.Pieces[Black][King] |
		board.Pieces[Black][Pawn]

	board.AllOccupied = board.WhiteOccupied | board.BlackOccupied
}

func FillSquares(board *Board) {
	colors := []Color{White, Black}
	pieces := []PieceType{Pawn, Rook, Knight, Bishop, Queen, King}

	for _, color := range colors {
		for _, piece := range pieces {
			bits := board.Pieces[color][piece]
			for sq := range 64 {
				if (bits>>sq)&1 == 1 {
					board.Squares[sq] = ToSquarePiece(color, piece)
				}
			}
		}
	}
}

func NewGame() *Board {
	board := &Board{}

	board.Pieces[White][Rook] = (1 << 0) | (1 << 7)
	board.Pieces[White][Knight] = (1 << 1) | (1 << 6)
	board.Pieces[White][Bishop] = (1 << 2) | (1 << 5)
	board.Pieces[White][Queen] = 1 << 3
	board.Pieces[White][King] = 1 << 4
	board.Pieces[White][Pawn] = 0xFF << 8

	board.Pieces[Black][Rook] = (1 << 56) | (1 << 63)
	board.Pieces[Black][Knight] = (1 << 57) | (1 << 62)
	board.Pieces[Black][Bishop] = (1 << 58) | (1 << 61)
	board.Pieces[Black][Queen] = 1 << 59
	board.Pieces[Black][King] = 1 << 60
	board.Pieces[Black][Pawn] = 0xFF << 48

	FillSquares(board)

	board.WhiteCanCastleKingSide = true
	board.WhiteCanCastleQueenSide = true
	board.BlackCanCastleKingSide = true
	board.BlackCanCastleQueenSide = true

	RecomputeOccupied(board)

	return board
}

// (Bit)Board Display
func PrintBitBoard(bits uint64) {
	for rank := 7; rank >= 0; rank-- {
		for file := range 8 {
			sq := rank*8 + file
			if (bits>>sq)&1 == 1 {
				fmt.Print("1 ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

// Board Scoring (MiniMax)
func GetPieceScore(piece PieceType) int {
	switch piece {
	case Pawn:
		return PawnValue
	case Knight:
		return KnighValue
	case Bishop:
		return BishopValue
	case Rook:
		return RookValue
	case Queen:
		return QueenValue
	case King:
		return KingValue
	}
	return 0
}

func GetScore(board *Board) float64 {
	var whiteScore int
	var blackScore int

	for i, piece := range board.Squares {
		if piece == Empty {
			continue
		}
		
		pieceColor, pieceType := ConvertPieceToPieceType(piece)

		if pieceColor == White {
			if pieceType == Knight {
				whiteScore += KnightPositionScore(i, pieceColor)
			}
			whiteScore += GetPieceScore(pieceType)
		} else {
			if pieceType == Knight {
				blackScore += KnightPositionScore(i, pieceColor)
			}
			blackScore += GetPieceScore(pieceType)
		}
	}

	return float64(whiteScore - blackScore)
}
