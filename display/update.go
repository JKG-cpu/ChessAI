package display

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

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

		// Cursor Movement
		switch msg.String() {
		case "up", "down", "left", "right":
			m.cursorSq = MoveCursor(m.cursorSq, msg.String())
		case "a", "b", "c", "d", "e", "f", "g", "h":
			m.cursorSq = MoveCursorToCol(m.cursorSq, msg.String())
		case "1", "2", "3", "4", "5", "6", "7", "8":
			m.cursorSq = MoveCursorToRow(m.cursorSq, msg.String())
		}
	}

	return m, nil
}