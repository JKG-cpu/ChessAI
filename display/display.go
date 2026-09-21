package display

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	if m.IsAiTurn() {
		return AiMoveCmd(m.board, 4, m.turn)
	}
	return nil
}

func RunDisplay(mode GameMode) {
	p := tea.NewProgram(initialModel(mode), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
