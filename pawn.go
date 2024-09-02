package main

type Pawn struct {
	side Color
}

// Side implements Piece.
func (p *Pawn) Side() Color {
	return p.side
}

// CanTake implements Piece.
func (p *Pawn) CanTake(other Piece) bool {
	if other == nil {
		return false // XXX: this domain logic shouldn't be here I think
	}
	return other.Side() != p.Side()
}

// IsValidMove implements Piece.
func (p *Pawn) IsValidMove(board *Board, x1 int, y1 int, x2 int, y2 int) bool {
	positiveSlope := y2-y1 == x2-x1
	negativeSlope := y2-y1 == x1-x2

	if p.Side() == White && (y2-y1) != 1 {
		return false
	} else if p.Side() == Black && (y2-y1) != -1 {
		return false
	}

	other := board.PieceAt(x2, y2)
	if positiveSlope || negativeSlope {
		return p.CanTake(other)
	}

	return other == nil && x1 == x2
}

var _ Piece = &Pawn{}
