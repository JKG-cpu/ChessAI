package display

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JKG-cpu/ChessAI/core"
	"github.com/charmbracelet/lipgloss"
)

func RenderBoard(board *core.Board, cursorSq int, selectedMoves []core.Move) string {
	var finalString strings.Builder

	leftSizePadding := strings.Repeat(" ", 4)
	rightSizePadding := strings.Repeat(" ", 4)

	// File Label
	fileLabel := "   A  B  C  D  E  F  G  H  "

	finalString.WriteString(leftSizePadding)
	finalString.WriteString(NeutralStyle.Render(fileLabel))
	finalString.WriteString(rightSizePadding)
	finalString.WriteString("\n")

	// Actual Board + Squares
	for rank := 7; rank >= 0; rank-- {
		finalString.WriteString(leftSizePadding)
		finalString.WriteString(NeutralStyle.Render(strconv.Itoa(rank + 1)))
		finalString.WriteString(" ")

		for file := range 8 {
			sq := rank*8 + file
			isDarkSquare := (file + rank) % 2 == 0
			piece := board.Squares[sq]

			isLegalMoveSq := false
			for _, move := range selectedMoves {
				if move.To == sq {
					isLegalMoveSq = true
					break
				}
			}

			var symbol string
			if piece == core.Empty {
				if isLegalMoveSq {
					symbol = LegalMoveStyle.Render("*")
				} else {
					symbol = " "
				}
			} else {
				pieceColor, pieceType := core.ConvertPieceToPieceType(piece)
				raw := pieceType.GetPieceASCII()

				if isLegalMoveSq {
					symbol = LegalMoveStyle.Render(raw)
				} else if pieceColor == core.White {
					symbol = WhitePieceStyle.Render(raw)
				} else {
					symbol = BlackPieceStyle.Render(raw)
				}
			}

			bracketOpen := "[ "
			bracketClose := " ]"

			if sq == cursorSq {
				bracketOpen = HighlightedStyle.Render("[")
				bracketClose = HighlightedStyle.Render("]")
			} else {
				if isDarkSquare {
					bracketOpen = BlackPieceStyle.Render("[")
					bracketClose = BlackPieceStyle.Render("]")
				} else {
					bracketOpen = WhitePieceStyle.Render("[")
					bracketClose = WhitePieceStyle.Render("]")
				}
			}

			finalString.WriteString(bracketOpen)
			finalString.WriteString(symbol)
			finalString.WriteString(bracketClose)
		}

		finalString.WriteString(" ")
		finalString.WriteString(NeutralStyle.Render(strconv.Itoa(rank + 1)))
		finalString.WriteString(rightSizePadding)
		finalString.WriteString("\n")
	}

	finalString.WriteString(leftSizePadding)
	finalString.WriteString(NeutralStyle.Render(fileLabel))
	finalString.WriteString(rightSizePadding)

	return BoardBorderStyle.Render(finalString.String())
}

func GetBoardSize(board string) (int, int) {
	return lipgloss.Width(board), lipgloss.Height(board)
}

func GetSquare(cursor int) string {
	file := cursor % 8
	rank := cursor / 8

	stringFile := string(rune('a' + file))
	stringRank := strconv.Itoa(rank + 1)

	return stringFile + stringRank
}

func RenderStatusPanel(m model) string {
	var sb strings.Builder

	sb.WriteString("Turn: ")
	if m.turn == core.White {
		sb.WriteString("White")
	} else {
		sb.WriteString("Black")
	}
	sb.WriteString("\n\n")
	sb.WriteString("Cursor: ")
	sb.WriteString(GetSquare(m.cursorSq))
	sb.WriteString("\n\n")

	switch m.mode {
	case PlayerVSPlayer:
		sb.WriteString("Player VS Player")
	case PlayerVSAI:
		sb.WriteString("Player VS AI")
	case AIVSAI:
		sb.WriteString("AI VS AI")
	}

	if m.gameOver {
		sb.WriteString("\n\n")
		if m.turn == core.White {
			sb.WriteString("Black Won!!!")
		} else {
			sb.WriteString("White Won!!!")
		}
	}

	return sb.String()
}

func (m model) View() string {
	const statusPanelVerticalOverhead = 2
	
	boardString := RenderBoard(m.board, m.cursorSq, m.selectedSqMoves)
	_, boardHeight := GetBoardSize(boardString)

	statusString := StatusPanelStyle.
		Height(boardHeight - statusPanelVerticalOverhead).
		Render(RenderStatusPanel(m))

	combined := lipgloss.JoinHorizontal(lipgloss.Top, boardString, statusString)

	requiredWidth, requiredHeight := GetBoardSize(combined)

	if m.width < requiredWidth || m.height < requiredHeight {
		return lipgloss.Place(
			m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			fmt.Sprintf(
				"Terminal too small.\nPlease resize to at least %dx%d (current: %dx%d)",
				requiredWidth, requiredHeight, m.width, m.height,
			),
		)
	}

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		combined,
	)
}