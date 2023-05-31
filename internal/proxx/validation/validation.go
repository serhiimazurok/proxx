package validation

import (
	"errors"
	"fmt"
	"strconv"
)

const BoardSizeMin = 5
const BoardSizeMax = 40

// BoardSize validate board size.
func BoardSize(input string) error {
	size, err := strconv.ParseInt(input, 10, 64)
	if err != nil || size < BoardSizeMin || size > BoardSizeMax {
		return errors.New(fmt.Sprintf("Invalid board size. Plase enter number between %d and %d", BoardSizeMin, BoardSizeMax))
	}
	return nil
}

// HolesCount validate holes count based on board size.
func HolesCount(size int) func(input string) error {
	min := size
	max := size*size - 1
	return func(input string) error {
		count, err := strconv.ParseInt(input, 10, 64)
		if err != nil || count < int64(min) || count > int64(max) {
			return errors.New(fmt.Sprintf("Invalid holes count. Plase enter number between %d and %d", min, max))
		}
		return nil
	}
}

// Coordinate validate coordinates.
func Coordinate(c string, size int) func(input string) error {
	min := 1
	max := size
	return func(input string) error {
		count, err := strconv.ParseInt(input, 10, 64)
		if err != nil || count < int64(min) || count > int64(max) {
			return errors.New(fmt.Sprintf("Invalid %s. Plase enter number between %d and %d", c, min, max))
		}
		return nil
	}
}
