package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Queen struct {
	side   Side
	sprite *ebiten.Image
}

func NewQueen(side Side) Queen {
	img := getSprite(side, "Queen")
	return Queen{
		side:   side,
		sprite: img,
	}
}

// Sprite implements Piece.
func (q *Queen) Sprite() *ebiten.Image {
	return q.sprite
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
