package board

import "math/rand"

func (b *Board) fill(holesCount int) {
	// generate holes
	holes := rand.Perm(b.BoardSize() * b.BoardSize())[:holesCount]

	for _, hole := range holes {
		b.fillHole(hole)
		b.fillHoleNeighbors(hole)
	}
}

func (b *Board) fillHole(hole int) {
	b.matrix[hole/b.BoardSize()+1][hole%b.BoardSize()+1] = Hole
}

func (b *Board) fillHoleNeighbors(hole int) {
	i := hole/b.BoardSize() + 1
	j := hole%b.BoardSize() + 1

	if b.matrix[i][j-1] != Hole {
		b.matrix[i][j-1]--
	}

	if b.matrix[i][j+1] != Hole {
		b.matrix[i][j+1]--
	}

	// top
	if b.matrix[i+1][j] != Hole {
		b.matrix[i+1][j]--
	}

	if b.matrix[i+1][j-1] != Hole {
		b.matrix[i+1][j-1]--
	}

	if b.matrix[i+1][j+1] != Hole {
		b.matrix[i+1][j+1]--
	}

	// bottom
	if b.matrix[i-1][j] != Hole {
		b.matrix[i-1][j]--
	}

	if b.matrix[i-1][j-1] != Hole {
		b.matrix[i-1][j-1]--
	}

	if b.matrix[i-1][j+1] != Hole {
		b.matrix[i-1][j+1]--
	}
}
