package main

import "github.com/hajimehoshi/ebiten/v2"

type Side uint8

func (s Side) String() string {
	switch s {
	case WHITE:
		return "white"
	case BLACK:
		return "black"
	default:
		panic("invalid side")
	}
}

const (
	NONE Side = iota
	WHITE
	BLACK
)

type Piece interface {
	IsValidMove(board *Board, x1, y1, x2, y2 int) bool
	Sprite() *ebiten.Image
	Side() Side
}

type Board [8][8]Piece

func (b *Board) PieceAt(x, y int) Piece {
	if x < 0 || y < 0 || x > 7 || y > 7 {
		return nil
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
