package display

import (
	"strconv"

	"github.com/JKG-cpu/ChessAI/core"
	tea "github.com/charmbracelet/bubbletea"
)

// Helpers
func SwitchTurn(m model) model {
	if m.turn == core.White {
		m.turn = core.Black
	} else {
		m.turn = core.White
	}
	return m
}

// Cursor Functions
func MoveCursor(cursor int, direction string) int {
	file := cursor % 8
	rank := cursor / 8

	switch direction {
	case "up":
		if rank < 7 { rank++ }
	case "down":
		if rank > 0 { rank-- }
	case "left":
		if file > 0 { file-- }
	case "right":
		if file < 7 { file++ }
	}

	return rank*8 + file
}

func MoveCursorToCol(cursor int, col string) int {
	if len(col) == 0 {
		return cursor
	}
	
	file := int(col[0] - 'a')
	rank := cursor / 8

	if file < 0 || file > 7 {
		return cursor
	}

	return rank*8 + file
}

func MoveCursorToRow(cursor int, col string) int {
	file := cursor % 8
	rank, _ := strconv.Atoi(col)

	return (rank - 1)*8 + file
}

func SelectSquareForLegalMoves(m model) (model, bool) {
	// Check if trying to move piece
	if m.selectedSq != -1 {
		for _, move := range m.selectedSqMoves {
			if move.To == m.cursorSq {
				m.board.Move(move)
				m = SwitchTurn(m)
				m = ClearSelection(m)

				// Check Game Over
				gameOver, gameMessage := core.IsGameOver(m.board, m.turn)

				m.gameOver = gameOver
				m.gameOverMessage = gameMessage

				return m, true
			}
		}
	}

	piece := m.board.Squares[m.cursorSq]
	pieceColor, _ := core.ConvertPieceToPieceType(piece)

	if pieceColor != m.turn || piece == core.Empty {
		return m, false
	}

	m.selectedSq = m.cursorSq
	m.selectedSqMoves = core.GetLegalMovesForSquare(m.board, m.selectedSq, m.turn)

	return m, false
}

func ClearSelection(m model) model {
	m.selectedSq = -1
	m.selectedSqMoves = nil

	return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		// Quitting
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Check for AI Move
		if m.IsAiTurn() {
			return m, nil
		}

		// Cursor Movement
		if !m.gameOver {
			switch msg.String() {
			case "up", "down", "left", "right":
				m.cursorSq = MoveCursor(m.cursorSq, msg.String())
			case "a", "b", "c", "d", "e", "f", "g", "h":
				m.cursorSq = MoveCursorToCol(m.cursorSq, msg.String())
			case "1", "2", "3", "4", "5", "6", "7", "8":
				m.cursorSq = MoveCursorToRow(m.cursorSq, msg.String())
			}

			// Piece Selection
			if msg.Type == tea.KeyEnter {
				var moveMade bool
				m, moveMade = SelectSquareForLegalMoves(m)

				if moveMade && m.IsAiTurn() && !m.gameOver {
					return m, AiMoveCmd(m.board, 4, m.turn)
				}
			}

			if msg.Type == tea.KeyEscape {
				m = ClearSelection(m)
			}
		}
	case AiMove:
		m.board.Move(msg.move)
		m = SwitchTurn(m)

		gameOver, gameMessage := core.IsGameOver(m.board, m.turn)
		m.gameOver = gameOver
		m.gameOverMessage = gameMessage

		if m.IsAiTurn() && !m.gameOver {
			return m, AiMoveCmd(m.board, 5, m.turn)
		}

		return m, nil
	}

	return m, nil
}