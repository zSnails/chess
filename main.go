package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var blackKnight = NewKnight(BLACK)
var blackPawn = NewPawn(BLACK)
var blackQueen = NewQueen(BLACK)
var blackKing = NewKing(BLACK)
var blackBishop = NewBishop(BLACK)
var blackRook = NewRook(BLACK)

var whiteKnight = NewKnight(WHITE)
var whitePawn = NewPawn(WHITE)
var whiteQueen = NewQueen(WHITE)
var whiteKing = NewKing(WHITE)
var whiteBishop = NewBishop(WHITE)
var whiteRook = NewRook(WHITE)

var board = Board{
	{&whiteRook, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackRook},
	{&whiteKnight, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackKnight},
	{&whiteBishop, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackBishop},
	{&whiteKing, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackQueen},
	{&whiteQueen, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackKing},
	{&whiteBishop, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackBishop},
	{&whiteKnight, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackKnight},
	{&whiteRook, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackRook},
}

type game struct {
	board       *ebiten.Image
	selectStart image.Point
}

// Draw implements ebiten.Game.
func (g *game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	screen.DrawImage(g.board, &opts)

	for x, row := range board {
		for y, piece := range row {
			if piece != nil {
				opts.GeoM.Translate(float64(16+(x*32)), float64((y*32)+16))
				screen.DrawImage(piece.Sprite(), &opts)
				opts.GeoM.Reset()
			}
		}
	}
}

// Layout implements ebiten.Game.
func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	// return 640, 480
	return 288, 288
}

// Update implements ebiten.Game.
func (g *game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		cellX := (x - 16) / 32
		cellY := (y - 16) / 32
		g.selectStart = image.Pt(cellX, cellY)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		cellX := (x - 16) / 32
		cellY := (y - 16) / 32
		board.Move(g.selectStart.X, g.selectStart.Y, cellX, cellY)
	}
	return nil
}

func main() {
	boardImage, _, err := ebitenutil.NewImageFromFile("./assets/Classic/Board/Board - classic 1.png")
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(&game{
		board: boardImage,
	}); err != nil {
		panic(err)
	}
}
