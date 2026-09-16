package display

import "github.com/charmbracelet/lipgloss"

var (
	WhitePieceStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	BlackPieceStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000"))

	NeutralStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a19e9e"))

	BoardBorderStyle = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#a19e9e")).
		Border(lipgloss.ASCIIBorder())
)