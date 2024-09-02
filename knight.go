package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Knight struct {
	side   Side
	sprite *ebiten.Image
}

func NewKnight(side Side) Knight {
	img := getSprite(side, "Knight")
	return Knight{
		side:   side,
		sprite: img,
	}
}

// Sprite implements Piece.
func (k *Knight) Sprite() *ebiten.Image {
	return k.sprite
}

// IsValidMove implements Piece.
func (k *Knight) IsValidMove(board *Board, x1, y1, x2, y2 int) bool {
	dx := int(math.Abs(float64(x1 - x2)))
	dy := int(math.Abs(float64(y1 - y2)))
	other := board.PieceAt(x2, y2)
	return (dx == 2 && dy == 1 || dx == 1 && dy == 2) && (other == nil || k.Side() != other.Side())
}

// Side implements Piece.
func (k *Knight) Side() Side {
	return k.side
}

var _ Piece = &Knight{}
