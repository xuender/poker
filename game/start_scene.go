package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/xuender/poker/pb"
)

type StartScene struct {
	bus  *Bus
	keys map[ebiten.Key]func()
}

func NewStart(bus *Bus) *StartScene {
	start := &StartScene{bus: bus}
	start.keys = map[ebiten.Key]func(){
		ebiten.KeyEscape: func() { bus.To(pb.Scene_desktop) },
	}

	return start
}

func (p *StartScene) Update() error               { return nil }
func (p *StartScene) Keys() map[ebiten.Key]func() { return p.keys }

func (p *StartScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "[ESC] run...")
}
