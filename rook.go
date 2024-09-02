package main

import (
	"math"
)

type Rook struct {
	side Color
}

// CanTake implements Piece.
func (r *Rook) CanTake(other Piece) bool {
	return r.Side() != other.Side()
}

func sign(val float64) int {
	if val == 0.0 {
		return 0
	}

	if math.Signbit(val) {
		return -1
	}
	return 1
}

// IsValidMove implements Piece.
func (r *Rook) IsValidMove(board *Board, x1 int, y1 int, x2 int, y2 int) bool {

	if x1 != x2 && y1 != y2 {
		return false
	}

	dirX := sign(float64(x2 - x1))
	dirY := sign(float64(y2 - y1))

	for x := x1 + dirX; x > -1 && x < 8 && x-x2 != 0; x += dirX {
		found := board.PieceAt(x, y1)
		if found != nil {
			return false
		}
	}

	for y := y1 + dirY; y > -1 && y < 8 && y-y2 != 0; y += dirY {
		found := board.PieceAt(x1, y)
		if found != nil {
			return false
		}
	}

	other := board.PieceAt(x2, y2)
	if other != nil {
		return r.CanTake(other)
	}

	return true
}

// Side implements Piece.
func (r *Rook) Side() Color {
	return r.side
}

var _ Piece = &Rook{}
