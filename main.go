package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
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

type game struct {
	boardSprite         *ebiten.Image
	selectStart         image.Point
	cellUnderMouse      image.Point
	cellUnderMouseColor color.Color
	dragged             Piece
	audioCtx            *audio.Context
	moveSound           *audio.Player
	takeSound           *audio.Player

	board Board
}

// Draw implements ebiten.Game.
func (g *game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	screen.DrawImage(g.boardSprite, &opts)

	vector.DrawFilledRect(screen, float32(g.cellUnderMouse.X), float32(g.cellUnderMouse.Y), 32, 32, g.cellUnderMouseColor, false)

	for x, row := range g.board {
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
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return io.EOF
	}
	x, y := ebiten.CursorPosition()
	cellX := max(0, min((x-16)/32, 7))
	cellY := max(0, min((y-16)/32, 7))

	xcoord, ycoord := 16+(cellX*32), (cellY*32)+16
	g.cellUnderMouse = image.Pt(xcoord, ycoord)
	g.cellUnderMouseColor = color.RGBA{A: 0x30}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.dragged = g.board.PieceAt(cellX, cellY)
		g.selectStart = image.Pt(cellX, cellY)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.dragged = nil
		moveState := g.board.Move(g.selectStart.X, g.selectStart.Y, cellX, cellY)
		if g.board.PieceAt(g.selectStart.X, g.selectStart.Y) != nil && moveState == Still {
			g.cellUnderMouseColor = color.RGBA{R: 0xFF, A: 0x30}
		}

		// play the sound
		switch moveState {
		case Took:
			g.takeSound.SetPosition(0)
			g.takeSound.Play()
		case Moved:
			g.moveSound.SetPosition(time.Millisecond * 33)
			g.moveSound.Play()
		case Still:
		default:
			panic(fmt.Sprintf("unexpected main.State: %#v", moveState))
		}
	}
	return nil
}

func main() {
	boardSprite, _, err := ebitenutil.NewImageFromFile("./assets/Classic/Board/Board - classic 1.png")
	if err != nil {
		panic(err)
	}

	audioCtx := audio.NewContext(24000)
	f, err := os.Open("./assets/audio/clank-2.raw")
	if err != nil {
		panic(err)
	}

	moveSound, err := audioCtx.NewPlayer(f)
	if err != nil {
		panic(err)
	}
	defer moveSound.Close()

	f, err = os.Open("./assets/audio/take.raw")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	takeSound, err := audioCtx.NewPlayer(f)
	if err != nil {
		panic(err)
	}
	defer takeSound.Close()

	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGameWithOptions(&game{
		boardSprite: boardSprite,
		audioCtx:    audioCtx,
		moveSound:   moveSound,
		takeSound:   takeSound,
		board: Board{
			{&whiteRook, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackRook},
			{&whiteKnight, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackKnight},
			{&whiteBishop, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackBishop},
			{&whiteKing, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackQueen},
			{&whiteQueen, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackKing},
			{&whiteBishop, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackBishop},
			{&whiteKnight, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackKnight},
			{&whiteRook, &whitePawn, nil, nil, nil, nil, &blackPawn, &blackRook},
		},
	}, &ebiten.RunGameOptions{
		X11ClassName:    "chess",
		X11InstanceName: "Chess",
	}); err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		panic(err)
	}
}
