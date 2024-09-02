package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	boardSprite    *ebiten.Image
	selectStart    image.Point
	cellUnderMouse image.Point
	dragged        Piece
}

// Draw implements ebiten.Game.
func (g *game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	screen.DrawImage(g.boardSprite, &opts)
	vector.DrawFilledRect(screen, float32(g.cellUnderMouse.X), float32(g.cellUnderMouse.Y), 32, 32, color.RGBA{A: 0x30}, false)
	for x, row := range board {
		for y, piece := range row {
			if piece != nil {
				opts.GeoM.Translate(float64(16+(x*32)), float64((y*32)+16))
				screen.DrawImage(piece.Sprite(), &opts)
				opts.GeoM.Reset()
			}
		}
	}

	if g.dragged != nil {
		x, y := ebiten.CursorPosition()
		opts.GeoM.Translate(float64(x)-16, float64(y)-16)
		opts.ColorScale.SetA(0.5)
		screen.DrawImage(g.dragged.Sprite(), &opts)
		opts.GeoM.Reset()
		opts.ColorScale.Reset()
	}

}

// Layout implements ebiten.Game.
func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	// return 640, 480
	return 288, 288
}

// Update implements ebiten.Game.
func (g *game) Update() error {
	x, y := ebiten.CursorPosition()
	cellX := max(0, min((x-16)/32, 7))
	cellY := max(0, min((y-16)/32, 7))

	xcoord, ycoord := 16+(cellX*32), (cellY*32)+16
	g.cellUnderMouse = image.Pt(xcoord, ycoord)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.dragged = board.PieceAt(cellX, cellY)
		g.selectStart = image.Pt(cellX, cellY)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.dragged = nil
		board.Move(g.selectStart.X, g.selectStart.Y, cellX, cellY)
	}
	return nil
}

func main() {
	boardSprite, _, err := ebitenutil.NewImageFromFile("./assets/Classic/Board/Board - classic 1.png")
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGameWithOptions(&game{
		boardSprite: boardSprite,
	}, &ebiten.RunGameOptions{
		X11ClassName:    "chess",
		X11InstanceName: "Chess",
	}); err != nil {
		panic(err)
	}
}
