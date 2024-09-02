package main

type Queen struct {
	side Side
}

// IsValidMove implements Piece.
func (q *Queen) IsValidMove(board *Board, x1 int, y1 int, x2 int, y2 int) bool {
	return isValidBishopMove(q, board, x1, y1, x2, y2) || isValidRookMove(q, board, x1, y1, x2, y2)
}

// Side implements Piece.
func (q *Queen) Side() Side {
	return q.side
}

var _ Piece = &Queen{}

// Valid Queen move, a queen's move is valid if it is either a valid bishop or rook move.
