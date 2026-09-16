package display

import "github.com/JKG-cpu/ChessAI/core"

type model struct {
	width  int
	height int

	board *core.Board
}

func initialModel() model {
	return model{board: core.NewGame()}
}