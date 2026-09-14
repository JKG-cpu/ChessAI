package core

import "fmt"

// Board Generation
func ToSquarePiece(color Color, piece PieceType) Piece {
	base := 1 + int(piece)
	if color == Black {
		base += 6
	}
	return Piece(base)
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
