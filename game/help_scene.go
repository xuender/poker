package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

const (
	_border = 100
)

type HelpScene struct {
	bus   *Bus
	helps [][2]string
	keys  map[ebiten.Key]func()
	face  font.Face
}

// nolint: gomnd
func NewHelp(bus *Bus, fonts *Fonts) *HelpScene {
	return &HelpScene{
		bus:  bus,
		face: fonts.MonospaceFace(26),
		helps: [][2]string{
			{"", "HELP"},
			{"[ESC]", "Exit"},
			{"[F2]", "Help"},
			{"[F11]", "Fullscreen"},
		},
		keys: map[ebiten.Key]func(){
			ebiten.KeyF2: func() {},
		},
	}
}

func (p *HelpScene) Update() error               { return nil }
func (p *HelpScene) Keys() map[ebiten.Key]func() { return p.keys }

func (p *HelpScene) Draw(screen *ebiten.Image) {
	width, height := p.bus.Layout()
	width -= _border
	width -= _border
	height -= _border
	height -= _border

	col := color.RGBA{0xdf, 0xd0, 0x00, 0xff}

	vector.DrawFilledRect(
		screen,
		_border, _border,
		float32(width), float32(height),
		color.RGBA{0xf, 0x60, 0x60, 0xdf},
		false)
	// nolint: gomnd
	for index, help := range p.helps {
		text.Draw(screen, help[0], p.face, 150, 150+30*index, col)
		text.Draw(screen, help[1], p.face, 300, 150+30*index, col)
	}
}
