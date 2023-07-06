package game

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/xuender/poker/pb"
)

type Desktop struct {
	bus    *Bus
	images []*ebiten.Image
	backs  *List
	my     *List
	out    *List
	do     bool
}

func NewDesktop(bus *Bus) *Desktop {
	ret := &Desktop{bus: bus}
	ret.images = make([]*ebiten.Image, len(pb.Poker_name))
	ret.init()

	return ret
}

func (p *Desktop) showPoker(screen *ebiten.Image, poker pb.Poker, x, y float64) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(x, y)
	screen.DrawImage(p.images[poker], op)
}

func (p *Desktop) Draw(screen *ebiten.Image) {
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

func (p *Desktop) Keys() map[ebiten.Key]func() { return nil }
func (p *Desktop) Update() error {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		p.do = false

		return nil
	}

	if p.do {
		return nil
	}

	p.do = true

	pox, poy := ebiten.CursorPosition()

	if poker := p.backs.Click(pox, poy); poker != pb.Poker_back {
		p.my.Add(poker)

		return nil
	}

	if poker := p.my.Click(pox, poy); poker != pb.Poker_back {
		p.out.Add(poker)

		return nil
	}

	if poker := p.out.Click(pox, poy); poker != pb.Poker_back {
		p.my.Add(poker)

		return nil
	}

	return nil
}

// nolint: gomnd
func (p *Desktop) init() {
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

	p.backs = NewList(100, 60, p.images[0].Bounds(), pokers...)
	p.backs.Back = true
	p.my = NewList(20, 400, p.images[0].Bounds())
	p.out = NewList(20, 230, p.images[0].Bounds())
}
