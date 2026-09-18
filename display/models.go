package display

import "github.com/JKG-cpu/ChessAI/core"

type GameMode int

const (
	PlayerVSPlayer = iota
	PlayerVSAI
	AIVSAI
)

type model struct {
	width  int
	height int

	board *core.Board

	mode GameMode
	turn core.Color

	cursorSq int
	selectedSq int
	selectedSqMoves []core.Move

	gameOver bool
	gameOverMessage string
}

func initialModel(mode GameMode) model {
	return model{
		board: core.NewGame(),
		mode: mode,
		cursorSq: 35,
		selectedSq: -1,
	}
}