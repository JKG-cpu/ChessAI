package main

import (
	"github.com/JKG-cpu/ChessAI/core"
)

func main() {
	board := core.NewGame()

	core.PrintBitBoard(board.BlackOccupied)
}
