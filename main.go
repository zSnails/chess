package main

import (
	"fmt"
)

type Color uint8

const (
	None Color = iota
	White
	Black
)

type Piece interface {
	IsValidMove(board *Board, x1, y1, x2, y2 int) bool
	CanTake(other Piece) bool
	Side() Color
}

type Board [8][8]Piece

func (b *Board) PieceAt(x, y int) Piece {
	if x > 7 || y > 7 {
		panic("unreachable: invalid position")
	}
	return b[x][y]
}

func (c *Board) CanMove(x1, y1, x2, y2 int) bool {
	if x2 > 7 || y2 > 7 || x1 < 0 || x2 < 0 { // out of bounds check
		return false
	}

	piece := c.PieceAt(x1, y1)
	fmt.Printf("piece: %+v\n", piece)
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

var rook = Rook{
	side: White,
}

var enemy = Pawn{
	side: Black,
}

var canvas = Board{
	{&enemy, &rook, nil, nil, nil, &rook, nil, nil},
	{&enemy, &enemy, nil, nil, nil, nil, nil, nil},
	{&enemy, nil, nil, nil, nil, nil, nil, nil},
	{&enemy, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
	{nil, nil, nil, nil, nil, nil, nil, nil},
}

func main() {
	fmt.Printf("canvas.Move(0, 5, 2, 5): %v\n", canvas.Move(0, 5, 2, 5))
	fmt.Printf("canvas.Move(2, 5, 2, 0): %v\n", canvas.Move(2, 5, 2, 0))
}
