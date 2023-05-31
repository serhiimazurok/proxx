package render

import (
	"fmt"
	"github.com/serhiimazurok/proxx/internal/board"
	"github.com/serhiimazurok/proxx/pkg/terminal/border"
	"github.com/serhiimazurok/proxx/pkg/terminal/color"
)

type BoardRenderer struct {
	board *board.Board
}

func New(board *board.Board) *BoardRenderer {
	return &BoardRenderer{
		board: board,
	}
}

func (r *BoardRenderer) Render() {
	fmt.Printf("    X")
	for i := 1; i < r.board.Size()-1; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Printf("\n")

	for i := 0; i < r.board.Size(); i++ {
		for j := 0; j < r.board.Size(); j++ {
			if i == 0 && j == 0 {
				fmt.Printf("  Y %c", border.BorderTopLeft)
			} else if i == 0 && j == r.board.Size()-1 {
				fmt.Printf("%c", border.BorderTopRight)
			} else if i == r.board.Size()-1 && j == 0 {
				fmt.Printf("    %c", border.BorderBottomLeft)
			} else if i == r.board.Size()-1 && j == r.board.Size()-1 {
				fmt.Printf("%c", border.BorderBottomRight)
			} else if i == 0 {
				fmt.Printf("%c%c%c", border.BorderHorizontal, border.BorderHorizontal, border.BorderHorizontal)
			} else if i == r.board.Size()-1 {
				fmt.Printf("%c%c%c", border.BorderHorizontal, border.BorderHorizontal, border.BorderHorizontal)
			} else if j == 0 {
				fmt.Printf(" %2d %c", i, border.BorderVertical)
			} else if j == r.board.Size()-1 {
				fmt.Printf("%c", border.BorderVertical)
			} else {
				r.renderField(i, j)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("✔ %v/%v", r.board.OpenedCells(), r.board.BoardSize()*r.board.BoardSize()-r.board.HolesCount())

	fmt.Printf("\n\n")

	switch {
	case r.board.IsWin():
		fmt.Println(color.Colorize(color.Green, "You win!"))
	case r.board.IsGameOver():
		fmt.Println(color.Colorize(color.Red, "You lose!"))
	}
}

func (r *BoardRenderer) renderField(i, j int) {
	var field string
	switch r.board.GetMaskedCell(i, j) {
	case board.Empty:
		field = " . "
	case board.Hole:
		field = color.Colorize(color.Red, " * ")
	case board.EmptyOpened:
		field = "   "
	default:
		field = color.Colorize(color.Green, fmt.Sprintf(" %d ", r.board.GetMaskedCell(i, j)))
	}
	fmt.Print(field)
}
