package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/xuender/poker/pb"
	"golang.org/x/image/font"
)

const (
	_fontSize = 60
	_two      = 2
)

type StartScene struct {
	bus  *Bus
	face font.Face
	keys map[ebiten.Key]func()
}

func NewStart(bus *Bus, fonts *Fonts) *StartScene {
	start := &StartScene{bus: bus, face: fonts.MonospaceFace(_fontSize)}
	start.keys = map[ebiten.Key]func(){
		ebiten.KeyEscape: func() { bus.To(pb.Scene_desktop) },
	}

	return start
}

func (p *StartScene) Update() error               { return nil }
func (p *StartScene) Keys() map[ebiten.Key]func() { return p.keys }

func (p *StartScene) Draw(screen *ebiten.Image) {
	txt := "[ESC] run..."
	width, height := p.bus.Layout()
	left := (width - len(txt)*_fontSize/_two) / _two
	top := height/_two - _fontSize/_two

	text.Draw(screen, txt, p.face, left, top, color.RGBA{0xdf, 0xd0, 0x00, 0xff})
}
