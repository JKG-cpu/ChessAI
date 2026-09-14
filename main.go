package main

import (
	"fmt"

	"github.com/JKG-cpu/ChessAI/core"
)

func main() {
	board := core.NewGame()

	core.PrintBitBoard(board.Pieces[core.Black][core.Knight])
	fmt.Println()
	core.PrintBitBoard(core.KnightMoves(board, 57, core.Black))
}
