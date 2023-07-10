package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/samber/lo"
	"github.com/xuender/poker/fonts"
	"golang.org/x/image/font"
)

const (
	_fontSize = 60
	_two      = 2
)

type StartScene struct {
	bus  *Bus
	face font.Face
}

func NewStart(bus *Bus) *StartScene {
	start := &StartScene{bus: bus, face: fonts.Head(_fontSize)}

	return start
}

func (p *StartScene) Update() error               { return nil }
func (p *StartScene) Keys() map[ebiten.Key]func() { return nil }

func (p *StartScene) Draw(screen *ebiten.Image) {
	max := lo.MaxBy(p.bus.Start, func(a, b string) bool { return len(a) > len(b) })
	width, height := p.bus.Layout()
	left := (width - len(max)*_fontSize/_two) / _two
	top := height/_two - (_fontSize*len(p.bus.Start))/_two

	for index, txt := range p.bus.Start {
		text.Draw(screen, txt, p.face, left, top+index*_fontSize, color.RGBA{0xdf, 0xd0, 0x00, 0xff})
	}
}
