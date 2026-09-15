package main

import (
	"fmt"
	"github.com/JKG-cpu/ChessAI/core"
)

func main() {
	board := core.NewGame()

	moves := core.GenerateAllMoves(board, core.White)

	fmt.Println(len(moves))
}
