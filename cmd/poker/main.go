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
	"github.com/xuender/poker/desktop"
	"github.com/xuender/poker/pb"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

type Poker struct {
	pokers      []pb.Poker
	images      []*ebiten.Image
	scene       pb.Scene
	sceneUpdate []func() error
	sceneDraw   []func(*ebiten.Image)
	backs       *desktop.List
	my          *desktop.List
	out         *desktop.List
	do          bool
}

func NewPoker() *Poker {
	poker := &Poker{
		pokers: []pb.Poker{pb.Poker_heart2, pb.Poker_heartA},
		images: make([]*ebiten.Image, len(pb.Poker_name)),
	}

	poker.sceneUpdate = []func() error{
		nilFunc,
		poker.update,
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
	// ebitenutil.DebugPrintAt(screen, strconv.Itoa(len(p.backs)), 400, 250)
	for _, img := range p.backs.Images() {
		p.showPoker(screen, img.Poker, img.X, img.Y)
	}

	for _, img := range p.out.Images() {
		p.showPoker(screen, img.Poker, img.X, img.Y)
	}

	for _, img := range p.my.Images() {
		p.showPoker(screen, img.Poker, img.X, img.Y)
	}
}

func (p *Poker) update() error {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		p.do = false

		return nil
	}

	if p.do {
		return nil
	}

	pox, poy := ebiten.CursorPosition()

	if poker := p.backs.Click(pox, poy); poker != pb.Poker_back {
		p.my.Add(poker)
	}

	if poker := p.my.Click(pox, poy); poker != pb.Poker_back {
		p.out.Add(poker)
	}

	if poker := p.out.Click(pox, poy); poker != pb.Poker_back {
		p.my.Add(poker)
	}

	// if len(p.backs) > 0 && x > 300 && x < 400+len(p.backs) && y > 200 && y < 300 {
	// 	p.my = append(p.my, p.backs[0])
	// 	p.backs = p.backs[1:]
	// }

	// if len(p.my) > 0 {
	// }

	p.do = true

	return nil
}

// nolint: gomnd
func (p *Poker) init() {
	for key := range pb.Poker_name {
		poker := pb.Poker(key)
		img, _ := lo.Must2(image.Decode(bytes.NewReader(poker.Bytes())))

		p.images[poker] = ebiten.NewImageFromImage(img)
	}

	pokers := make([]pb.Poker, 54)

	for i := 1; i <= 54; i++ {
		pokers[i-1] = pb.Poker(i)
	}

	pokers = lo.Shuffle(pokers)

	p.backs = desktop.NewList(100, 60, p.images[0].Bounds(), pokers...)
	p.backs.Back = true
	p.my = desktop.NewList(20, 400, p.images[0].Bounds())
	p.out = desktop.NewList(20, 230, p.images[0].Bounds())
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
