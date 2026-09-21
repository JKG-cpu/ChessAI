package display

import (
	"github.com/JKG-cpu/ChessAI/core"
	"github.com/JKG-cpu/ChessAI/minimax"
	tea "github.com/charmbracelet/bubbletea"
)

type GameMode int

const (
	PlayerVSPlayer = iota
	PlayerVSAI
	AIVSAI
)

type AiMove struct {
	move core.Move
}

func AiMoveCmd(board *core.Board, depth int, color core.Color) tea.Cmd {
	boardCopy := *board

	return func() tea.Msg {
		move := minimax.GetBestMove(&boardCopy, depth, color)
		return AiMove{move: move}
	}
}

type model struct {
	width  int
	height int

	board *core.Board

	mode GameMode
	player1Color core.Color
	player2Color core.Color
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
		player1Color: core.White,
		player2Color: core.Black,
		selectedSq: -1,
	}
}