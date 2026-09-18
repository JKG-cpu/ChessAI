package core

type Board struct {
	Pieces [2][6]uint64

	Squares [64]Piece

	WhiteOccupied uint64
	BlackOccupied uint64
	AllOccupied   uint64

	moveHistory []Move

	enPassantTarget int
	prevEnPassantTarget int
}

func (board *Board) Move(move Move) {
	color, pieceType := ConvertPieceToPieceType(move.Piece)

	if pieceType == Pawn {
		diff := move.To - move.From
		if diff == 16 || diff == -16 {
			board.enPassantTarget = (move.From + move.To) / 2
		} else {
			board.enPassantTarget = -1
		}
	} else {
		board.enPassantTarget = -1
	}

	board.Pieces[color][pieceType] &= ^(uint64(1) << move.From)
	board.Pieces[color][pieceType] |= uint64(1) << move.To

	if move.isEnPassant {
		capturedPawnSq := move.From/8 * 8 + move.To % 8
		capColor, capPiece := ConvertPieceToPieceType(move.Captured)
		board.Pieces[capColor][capPiece] &= ^(uint64(1) << uint64(capturedPawnSq))
		board.Squares[capturedPawnSq] = Empty
	} else if move.Captured != Empty {
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

	if move.isEnPassant {
		capturedPawnSq := move.From/8*8 + move.To%8
		capColor, capType := ConvertPieceToPieceType(move.Captured)
		board.Pieces[capColor][capType] |= uint64(1) << capturedPawnSq
		board.Squares[capturedPawnSq] = move.Captured
		board.Squares[move.To] = Empty
	} else if move.Captured != Empty {
		capColor, capType := ConvertPieceToPieceType(move.Captured)
		board.Pieces[capColor][capType] |= uint64(1) << move.To
		board.Squares[move.To] = move.Captured
	} else {
		board.Squares[move.To] = Empty
	}

	board.Squares[move.From] = move.Piece
	board.enPassantTarget = move.PrevEnPassantCapture

	RecomputeOccupied(board)
}

type Move struct {
	From     int
	To       int
	Piece    Piece
	Captured Piece
	PrevEnPassantCapture int
	isEnPassant bool
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
