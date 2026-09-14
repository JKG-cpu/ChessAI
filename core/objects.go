package core

type Board struct {
	Pieces [2][6]uint64

	Squares [64]Piece

	WhiteOccupied uint64
	BlackOccupied uint64
	AllOccupied   uint64
}

type Move struct {
	From     int
	To       int
	Piece    Piece
	Captured Piece
}

type Piece uint8

const (
	Empty Piece = iota

	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing

	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

type PieceType uint8

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

type Color uint8

const (
	White Color = iota
	Black
)
