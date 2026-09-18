package main

import (
	"os"
	"os/exec"
	"runtime"
	
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
	display.RunDisplay(display.PlayerVSPlayer)
	ClearTerminal()
}
