package main

import (
	"math"
)

type Bishop struct {
	side Side
}

func isValidBishopMove(b Piece, board *Board, x1, y1, x2, y2 int) bool {
	dx := int((float64(x2 - x1)))
	dy := int((float64(y2 - y1)))

	if math.Abs(float64(dx)) != math.Abs(float64(dy)) {
		return false
	}

	xs := sign(float64(dx))
	ys := sign(float64(dy))
	for change := 1; change < int(math.Abs(float64(dx))); change++ {
		piece := board.PieceAt(x1+(xs*change), y1+(ys*change))
		if piece != nil {
			return false
		}
	}

	other := board.PieceAt(x2, y2)
	return other == nil || b.Side() != other.Side()
}

// IsValidMove implements Piece.
func (b *Bishop) IsValidMove(board *Board, x1, y1, x2, y2 int) bool {
	return isValidBishopMove(b, board, x1, y1, x2, y2)
}

// Side implements Piece.
func (b *Bishop) Side() Side {
	return b.side
}

var _ Piece = &Bishop{}
