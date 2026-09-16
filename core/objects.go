package core

type Board struct {
	Pieces [2][6]uint64

	Squares [64]Piece

	WhiteOccupied uint64
	BlackOccupied uint64
	AllOccupied   uint64
}

func (board *Board) Move(move Move) {
	color, pieceType := ConvertPieceToPieceType(move.Piece)

	board.Pieces[color][pieceType] &= ^(uint64(1) << move.From)
	board.Pieces[color][pieceType] |= uint64(1) << move.To

	if move.Captured != Empty {
		capColor, capType := ConvertPieceToPieceType(move.Captured)
		board.Pieces[capColor][capType] &= ^(uint64(1) << move.To)
	}

	board.Squares[move.From] = Empty
	board.Squares[move.To] = move.Piece

	RecomputeOccupied(board)
}

func (board *Board) UndoMove(move Move) {
	color, pieceType := ConvertPieceToPieceType(move.Piece)

	board.Pieces[color][pieceType] |= uint64(1) << move.From
	board.Pieces[color][pieceType] &= ^(uint64(1) << move.To)

	if move.Captured != Empty {
		capColor, capType := ConvertPieceToPieceType(move.Captured)
		board.Pieces[capColor][capType] |= uint64(1) << move.To
	}

	board.Squares[move.From] = move.Piece
	board.Squares[move.To] = move.Captured

	RecomputeOccupied(board)
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

func (p PieceType) GetPieceASCII() string {
	switch p {
	case Pawn:
		return "P"
	case Knight:
		return "N"
	case Bishop:
		return "B"
	case Rook:
		return "R"
	case Queen:
		return "Q"
	case King:
		return "K"
	}
	return " "
}

type Color uint8

const (
	White Color = iota
	Black
)
