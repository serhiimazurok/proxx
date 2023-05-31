package proxx

import (
	"fmt"
	"github.com/serhiimazurok/proxx/internal/board"
	"github.com/serhiimazurok/proxx/internal/board/render"
	"github.com/serhiimazurok/proxx/pkg/terminal"
	"github.com/serhiimazurok/proxx/pkg/terminal/color"
	"math/rand"
	"time"
)

func Run() {
	rand.NewSource(time.Now().Unix())

	for {
		err := Play()
		if err != nil {
			fmt.Println("something went wrong", err)
		}

		ok, err := AskToContinue()

		if err != nil {
			fmt.Println("something went wrong", err)
		}

		if !ok {
			break
		}
	}

	terminal.Clear()

	fmt.Println(color.Colorize(color.Green, "Bye!"))
}

func Play() error {
	terminal.Clear()

	size, err := AskBoardSize()
	if err != nil {
		return err
	}

	holes, err := AskHolesCount(size)
	if err != nil {
		return err
	}

	terminal.Clear()

	brd := board.New(size, holes)
	r := render.New(brd)

	for {
		r.Render()

		x, err := AskCoordinate("X", size)
		if err != nil {
			return err
		}

		y, err := AskCoordinate("Y", size)
		if err != nil {
			return err
		}

		brd.Open(y, x)

		terminal.Clear()

		// stop an interaction if game is finished
		if brd.IsGameOver() || brd.IsWin() {
			break
		}
	}

	// render board last time with win/lose message
	r.Render()

	return nil
}
