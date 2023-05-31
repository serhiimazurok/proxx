package board

// State represents the state of the game
type State int

const (
	Continue State = iota
	Win
	Lose
)

// FinishGame finishes the game
func (b *Board) FinishGame(s State) {
	if s != Continue {
		b.state = s
	}
}

// IsGameOver returns true if the game is over
func (b *Board) IsGameOver() bool {
	return b.state == Lose
}

// IsWin returns true if the game is won
func (b *Board) IsWin() bool {
	return b.state == Win
}
