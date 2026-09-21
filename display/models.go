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
	whiteIsAi bool
	blackIsAi bool
	turn core.Color

	cursorSq int
	selectedSq int
	selectedSqMoves []core.Move

	gameOver bool
	gameOverMessage string
}

func (m model) IsAiTurn() bool {
	if m.turn == core.White && m.whiteIsAi {
		return true
	}

	if m.turn == core.Black && m.blackIsAi {
		return true
	}

	return false
}

func initialModel(mode GameMode) model {
	if mode == PlayerVSAI {
		return model{
			board: core.NewGame(),
			mode: mode,
			blackIsAi: true,
			whiteIsAi: false,
			cursorSq: 35,
			selectedSq: -1,
		}
	}

	if mode == AIVSAI {
		return model{
			board: core.NewGame(),
			mode: mode,
			blackIsAi: true,
			whiteIsAi: true,
			cursorSq: -1,
			selectedSq: -1,
		}
	}

	return model{
		board: core.NewGame(),
		mode: mode,
		blackIsAi: false,
		whiteIsAi: false,
		cursorSq: 35,
		selectedSq: -1,
	}
}