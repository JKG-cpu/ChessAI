package main

import (
	"os"
	"os/exec"
	"runtime"
	
	// "fmt"
	// "github.com/JKG-cpu/ChessAI/core"
	"github.com/JKG-cpu/ChessAI/display"
)

func ClearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// board := core.NewGame()

	// // Manually clear f1 and g1 to test White kingside castling
	// board.Pieces[core.White][core.Knight] &= ^(uint64(1) << 6) // remove knight from g1
	// board.Pieces[core.White][core.Bishop] &= ^(uint64(1) << 5) // remove bishop from f1
	// board.Squares[5] = core.Empty
	// board.Squares[6] = core.Empty
	// core.RecomputeOccupied(board)

	// moves := core.GenerateAllLegalMoves(board, core.White)
	// fmt.Println(moves)
	display.RunDisplay(display.PlayerVSPlayer)
	// ClearTerminal()
}
