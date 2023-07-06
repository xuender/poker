package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Help struct {
	bus        *Bus
	background *ebiten.Image
	helps      [][2]string
}

// nolint: gomnd
func NewHelp(bus *Bus) *Help {
	background := ebiten.NewImage(bus.Screen.X-200, bus.Screen.Y-200)
	background.Fill(color.RGBA{0xf, 0x60, 0x60, 0x9f})

	return &Help{
		bus:        bus,
		background: background,
		helps: [][2]string{
			{"", "HELP"},
			{"[ESC]", "Exit"},
			{"[F2]", "Help"},
			{"[F11]", "Fullscreen"},
		},
	}
}

func (p *Help) Update() error               { return nil }
func (p *Help) Keys() map[ebiten.Key]func() { return nil }

// nolint: gomnd
func (p *Help) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100)
	screen.DrawImage(p.background, op)

	for index, help := range p.helps {
		ebitenutil.DebugPrintAt(screen, help[0], 150, 150+30*index)
		ebitenutil.DebugPrintAt(screen, help[1], 200, 150+30*index)
	}
}
