package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/samber/lo"
	"github.com/xuender/poker/pb"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Poker struct {
	pokers      []pb.Poker
	images      []*ebiten.Image
	scene       pb.Scene
	sceneUpdate []func() error
	sceneDraw   []func(*ebiten.Image)
}

func NewPoker() *Poker {
	poker := &Poker{
		pokers: []pb.Poker{pb.Poker_heart2, pb.Poker_heartA},
		images: make([]*ebiten.Image, len(pb.Poker_name)),
	}

	poker.sceneUpdate = []func() error{
		nilFunc,
		nilFunc,
	}
	poker.sceneDraw = []func(*ebiten.Image){
		poker.loadDraw,
		poker.showDraw,
	}

	go poker.init()

	return poker
}

func nilFunc() error { return nil }

func (p *Poker) Update() error              { return p.sceneUpdate[p.scene]() }
func (p *Poker) Draw(screen *ebiten.Image)  { p.sceneDraw[p.scene](screen) }
func (p *Poker) Layout(_, _ int) (int, int) { return screenWidth, screenHeight }

func (p *Poker) loadDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "load...")
}

func (p *Poker) showPoker(screen *ebiten.Image, poker pb.Poker, x, y float64) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(x, y)
	screen.DrawImage(p.images[poker], op)
}

func (p *Poker) showDraw(screen *ebiten.Image) {
	// p.showPoker(screen, pb.Poker_heartA, 0, 0)
	// nolint: gomnd
	for i := 0; i < len(pb.Poker_name); i++ {
		x := i % 13 * 56
		y := i / 13 * 80

		p.showPoker(screen, pb.Poker(i), float64(x), float64(y))
	}
}

func (p *Poker) init() {
	for key := range pb.Poker_name {
		poker := pb.Poker(key)
		img, _ := lo.Must2(image.Decode(bytes.NewReader(poker.Image())))

		p.images[poker] = ebiten.NewImageFromImage(img)
	}

	p.scene = pb.Scene_show
}

func main() {
	flag.Usage = usage
	flag.Parse()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	// ebiten.SetWindowSize(400, 300)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("自由扑克")
	lo.Must0(ebiten.RunGame(NewPoker()))
}

func usage() {
	fmt.Fprintf(os.Stderr, "poker\n\n")
	fmt.Fprintf(os.Stderr, "显示扑克.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
