package core

type Board struct {
	Pieces [2][6]uint64

	Squares [64]Piece

	WhiteOccupied uint64
	BlackOccupied uint64
	AllOccupied   uint64

	moveHistory []Move

	enPassantTarget     int
	prevEnPassantTarget int

	WhiteCanCastleKingSide  bool
	WhiteCanCastleQueenSide bool
	BlackCanCastleKingSide  bool
	BlackCanCastleQueenSide bool
}

func (board *Board) Move(move Move) {
	color, pieceType := ConvertPieceToPieceType(move.Piece)

	// Snapshot castling rights before mutating anything
	// move.PrevWhiteCanCastleKingSide = board.WhiteCanCastleKingSide
	// move.PrevWhiteCanCastleQueenSide = board.WhiteCanCastleQueenSide
	// move.PrevBlackCanCastleKingSide = board.BlackCanCastleKingSide
	// move.PrevBlackCanCastleQueenSide = board.BlackCanCastleQueenSide

	// En passant target tracking
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

	// Castling rights updates
	if move.Piece == WhiteKing {
		board.WhiteCanCastleKingSide = false
		board.WhiteCanCastleQueenSide = false
	}
	if move.Piece == BlackKing {
		board.BlackCanCastleKingSide = false
		board.BlackCanCastleQueenSide = false
	}
	if move.From == 0 || move.To == 0 {
		board.WhiteCanCastleQueenSide = false
	}
	if move.From == 7 || move.To == 7 {
		board.WhiteCanCastleKingSide = false
	}
	if move.From == 56 || move.To == 56 {
		board.BlackCanCastleQueenSide = false
	}
	if move.From == 63 || move.To == 63 {
		board.BlackCanCastleKingSide = false
	}

	// Move the king (or any piece)
	board.Pieces[color][pieceType] &= ^(uint64(1) << move.From)
	board.Pieces[color][pieceType] |= uint64(1) << move.To

	// Castling: also move the rook
	if move.isCastle {
		var rookFrom, rookTo int
		switch move.To {
		case 6:
			rookFrom, rookTo = 7, 5
		case 2:
			rookFrom, rookTo = 0, 3
		case 62:
			rookFrom, rookTo = 63, 61
		case 58:
			rookFrom, rookTo = 56, 59
		}
		board.Pieces[color][Rook] &= ^(uint64(1) << rookFrom)
		board.Pieces[color][Rook] |= uint64(1) << rookTo
		board.Squares[rookFrom] = Empty
		board.Squares[rookTo] = ToSquarePiece(color, Rook)
	}

	if move.isEnPassant {
		capturedPawnSq := move.From/8*8 + move.To%8
		capColor, capType := ConvertPieceToPieceType(move.Captured)
		board.Pieces[capColor][capType] &= ^(uint64(1) << uint64(capturedPawnSq))
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

	if move.isCastle {
		var rookFrom, rookTo int
		switch move.To {
		case 6:
			rookFrom, rookTo = 7, 5
		case 2:
			rookFrom, rookTo = 0, 3
		case 62:
			rookFrom, rookTo = 63, 61
		case 58:
			rookFrom, rookTo = 56, 59
		}
		board.Pieces[color][Rook] |= uint64(1) << rookFrom
		board.Pieces[color][Rook] &= ^(uint64(1) << rookTo)
		board.Squares[rookFrom] = ToSquarePiece(color, Rook)
		board.Squares[rookTo] = Empty
	}

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
	board.WhiteCanCastleKingSide = move.PrevWhiteCanCastleKingSide
	board.WhiteCanCastleQueenSide = move.PrevWhiteCanCastleQueenSide
	board.BlackCanCastleKingSide = move.PrevBlackCanCastleKingSide
	board.BlackCanCastleQueenSide = move.PrevBlackCanCastleQueenSide

	RecomputeOccupied(board)
}

type Move struct {
	From                 int
	To                   int
	Piece                Piece
	Captured             Piece
	PrevEnPassantCapture int
	isEnPassant          bool
	isCastle             bool

	PrevWhiteCanCastleKingSide  bool
	PrevWhiteCanCastleQueenSide bool
	PrevBlackCanCastleKingSide  bool
	PrevBlackCanCastleQueenSide bool
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
