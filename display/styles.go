package display

import "github.com/charmbracelet/lipgloss"

var (
	WhitePieceStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	BlackPieceStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#b6a0a0")).
		Bold(true)

	NeutralStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a19e9e")).
		Bold(true)

	HighlightedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1eff00")).
		Bold(true)

	LegalMoveStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a19e9e"))

	BoardBorderStyle = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#a19e9e")).
		Border(lipgloss.ASCIIBorder())

	StatusPanelStyle = lipgloss.NewStyle().
		Border(lipgloss.ASCIIBorder()).
		BorderForeground(lipgloss.Color("#a19e9e")).
		Padding(1, 2).
		Width(20).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center)
)