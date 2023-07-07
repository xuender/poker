package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/xuender/poker/fonts"
	"golang.org/x/image/font"
)

const (
	_border = 100
)

type HelpScene struct {
	bus      *Bus
	helps    [][2]string
	keys     map[ebiten.Key]func()
	headFace font.Face
	bodyFace font.Face
}

// nolint: gomnd
func NewHelp(bus *Bus) *HelpScene {
	return &HelpScene{
		bus:      bus,
		headFace: fonts.Head(50),
		bodyFace: fonts.Body(30),
		helps: [][2]string{
			{"", "HELP"},
			{"", "测试"},
			{"[F2]", "帮助"},
			{"[F11]", "进入全屏/退出全屏"},
			{"[ESC]", "退出"},
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
		if help[0] == "" {
			text.Draw(screen, help[1], p.headFace, 300, 150+50*index, col)
		} else {
			text.Draw(screen, help[0], p.bodyFace, 150, 150+50*index, col)
			text.Draw(screen, help[1], p.bodyFace, 300, 150+50*index, col)
		}
	}
}
