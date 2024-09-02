package main

type Pawn struct {
	side Side
}

// Side implements Piece.
func (p *Pawn) Side() Side {
	return p.side
}

// IsValidMove implements Piece.
func (p *Pawn) IsValidMove(board *Board, x1 int, y1 int, x2 int, y2 int) bool {
	positiveSlope := y2-y1 == x2-x1
	negativeSlope := y2-y1 == x1-x2

	if p.Side() == WHITE && (y2-y1) != 1 || p.Side() == BLACK && (y2-y1) != -1 {
		return false
	}

	other := board.PieceAt(x2, y2)
	if positiveSlope || negativeSlope {
		return other == nil || p.Side() != other.Side()
	}

	return other == nil && x1 == x2
}

var _ Piece = &Pawn{}
