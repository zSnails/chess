package main

import "fmt"

type Side uint8

const (
	NONE Side = iota
	WHITE
	BLACK
)

type Piece interface {
	IsValidMove(board *Board, x1, y1, x2, y2 int) bool
	Side() Side
}

type Board [8][8]Piece

func (b *Board) PieceAt(x, y int) Piece {
	if x < 0 || y < 0 || x > 7 || y > 7 {
		panic("unreachable: invalid position")
	}
	return b[x][y]
}

func (c *Board) CanMove(x1, y1, x2, y2 int) bool {
	if x2 > 7 || y2 > 7 || x1 < 0 || x2 < 0 { // out of bounds check
		return false
	}

	piece := c.PieceAt(x1, y1)
	return piece != nil && piece.IsValidMove(c, x1, y1, x2, y2)
}

func (c *Board) Move(x1, y1, x2, y2 int) bool {
	if c.CanMove(x1, y1, x2, y2) {
		c[x2][y2] = c[x1][y1]
		c[x1][y1] = nil
		return true
	}

	return false
}

var blackRook = Rook{
	side: BLACK,
}

var whiteRook = Rook{
	side: WHITE,
}

var pawn = Pawn{
	side: BLACK,
}

var knight = Knight{
	side: WHITE,
}

var bishop = Bishop{
	side: WHITE,
}

var king = King{
	side: WHITE,
}

var queen = Queen{
	side: WHITE,
}

var board = Board{
	{&knight, &pawn, nil, nil, nil, nil, nil, nil},
	{&pawn, &pawn, &whiteRook, nil, nil, nil, nil, nil},
	{&bishop, nil, nil, nil, nil, nil, nil, nil},
	{nil, &pawn, nil, nil, nil, nil, nil, nil},
	{&queen, nil, nil, nil, nil, nil, &pawn, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
}

func main() {
	fmt.Printf("board.Move(0, 0, 7, 7): %v\n", board.Move(4, 0, 2, 2))
}
