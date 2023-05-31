package board

// Cell represents the cell
type Cell int8

const (
	Hole        Cell = -100
	Empty       Cell = 0
	EmptyOpened Cell = 100
)

// Board represents the board
type Board struct {
	holes       int
	openedCells int
	state       State
	matrix      [][]Cell
}

// New creates a new board
func New(size, holesCount int) *Board {
	// initialize the matrix
	// decided to create a matrix with a border of empty cells
	// to avoid checking the borders of the matrix and for easier rendering
	matrix := make([][]Cell, size+2)
	for i := 0; i < size+2; i++ {
		matrix[i] = make([]Cell, size+2)
	}

	board := &Board{
		openedCells: 0,
		holes:       holesCount,
		state:       Continue,
		matrix:      matrix,
	}

	board.fill(holesCount)

	return board
}

// HolesCount returns the number of holes
func (b *Board) HolesCount() int {
	return b.holes
}

// OpenedCells returns the number of opened cells
func (b *Board) OpenedCells() int {
	return b.openedCells
}

// Size returns the size of the board including the border
func (b *Board) Size() int {
	return len(b.matrix)
}

// BoardSize returns the size of the board
func (b *Board) BoardSize() int {
	return len(b.matrix) - 2
}

// GetMaskedCell returns the masked cell
func (b *Board) GetMaskedCell(i, j int) Cell {
	if b.matrix[i][j] <= 0 {
		// show holes if the game is over
		if b.matrix[i][j] == Hole && b.IsGameOver() {
			return Hole
		}
		return Empty
	}
	return b.matrix[i][j]
}

// Open opens the cell
func (b *Board) Open(i, j int) {
	if b.matrix[i][j] == Hole {
		b.FinishGame(Lose)
		return
	}

	b.openedCells += b.openRecursive(i, j)

	if b.OpenedCells() == b.BoardSize()*b.BoardSize()-b.HolesCount() {
		b.FinishGame(Win)
		return
	}
}

func (b *Board) openRecursive(i, j int) int {
	if !(i >= 1 && i <= b.BoardSize() && j >= 1 && j <= b.BoardSize()) {
		return 0
	}

	// in this case don't open the cell
	if b.matrix[i][j] == Hole || b.matrix[i][j] > 0 {
		return 0
	}

	// just need to open the cell and stop recursion
	if b.matrix[i][j] < 0 {
		b.matrix[i][j] = -b.matrix[i][j]
		return 1
	}

	var opened int

	if b.matrix[i][j] == Empty {
		b.matrix[i][j] = EmptyOpened
		opened++
	}

	opened += b.openRecursive(i-1, j)
	opened += b.openRecursive(i-1, j-1)
	opened += b.openRecursive(i-1, j+1)

	opened += b.openRecursive(i+1, j)
	opened += b.openRecursive(i+1, j-1)
	opened += b.openRecursive(i+1, j+1)

	opened += b.openRecursive(i, j-1)
	opened += b.openRecursive(i, j+1)

	return opened
}
