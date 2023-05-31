package proxx

import (
	"github.com/manifoldco/promptui"
	"github.com/serhiimazurok/proxx/internal/proxx/validation"
	"strconv"
)

func AskCoordinate(c string, size int) (int, error) {
	prompt := promptui.Prompt{
		Label:    c,
		Validate: validation.Coordinate(c, size),
	}

	result, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	x, err := strconv.ParseInt(result, 10, 64)

	if err != nil {
		return 0, err
	}

	return int(x), nil
}

func AskBoardSize() (int, error) {
	prompt := promptui.Prompt{
		Label:    "Board Size",
		Validate: validation.BoardSize,
		Default:  "8",
	}

	result, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	size, err := strconv.ParseInt(result, 10, 64)

	if err != nil {
		return 0, err
	}

	return int(size), nil
}

func AskHolesCount(size int) (int, error) {
	prompt := promptui.Prompt{
		Label:    "Holes Count",
		Validate: validation.HolesCount(size),
	}

	result, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	count, err := strconv.ParseInt(result, 10, 64)

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func AskToContinue() (bool, error) {
	prompt := promptui.Select{
		HideHelp: true,
		Label:    "Are you going to continue",
		Items:    []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return false, err
	}

	return result == "Yes", nil
}
