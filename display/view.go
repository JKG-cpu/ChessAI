package display

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JKG-cpu/ChessAI/core"
)

func RenderBoard(board *core.Board) string {
	var finalString strings.Builder

	leftSizePadding := strings.Repeat(" ", 4)

	// File Label
	fileLabel := "    A    B    C    D    E    F    G    H  "

	finalString.WriteString(leftSizePadding)
	finalString.WriteString(NeutralStyle.Render(fileLabel))
	finalString.WriteString("\n")

	// Actual Board + Squares
	for rank := 7; rank >= 0; rank-- {
		finalString.WriteString(leftSizePadding)
		finalString.WriteString(NeutralStyle.Render(strconv.Itoa(rank + 1)))
		finalString.WriteString(" ")

		for file := range 8 {
			sq := rank*8 + file
			piece := board.Squares[sq]

			var symbol string
			if piece == core.Empty {
				symbol = " "
			} else {
				pieceColor, pieceType := core.ConvertPieceToPieceType(piece)
				symbol = pieceType.GetPieceASCII()

				if pieceColor == core.White {
					symbol = WhitePieceStyle.Render(symbol)
				} else {
					symbol = BlackPieceStyle.Render(symbol)
				}
			}

			finalString.WriteString(NeutralStyle.Render("[ "))
			finalString.WriteString(symbol)
			finalString.WriteString(NeutralStyle.Render(" ]"))
		}

		finalString.WriteString(" ")
		finalString.WriteString(NeutralStyle.Render(strconv.Itoa(rank + 1)))
		finalString.WriteString("\n")
	}

	finalString.WriteString(leftSizePadding)
	finalString.WriteString(NeutralStyle.Render(fileLabel))

	return finalString.String()
	// return BoardBorderStyle.Render(finalString.String())
}

// func GetBoardSize(board string) (int, int) {

// }

func (m model) View() string {
	const minWidth = 50
	const minHeight = 15

	if m.width < minWidth || m.height < minHeight {
		return fmt.Sprintf(
			"Terminal too small.\nPlease resize to at least %dx%d (current: %dx%d)",
			minWidth, minHeight, m.width, m.height,
		)
	}

	return RenderBoard(m.board)
}
