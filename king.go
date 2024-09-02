package main

import "math"

type King struct {
	side Side
}

// IsValidMove implements Piece.
func (k *King) IsValidMove(board *Board, x1 int, y1 int, x2 int, y2 int) bool {
	// Valid King move, if the piece moves from (X1, Y1) to (X2, Y2), the move is valid if and only if |X2-X1|<=1 and |Y2-Y1|<=1.
	other := board.PieceAt(x2, y2)
	return math.Abs(float64(x2-x1)) <= 1 && math.Abs(float64(y2-y1)) <= 1 && (other == nil || k.Side() != other.Side())
}

// Side implements Piece.
func (k *King) Side() Side {
	return k.side
}

var _ Piece = &King{}
