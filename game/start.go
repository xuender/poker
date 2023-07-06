package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/xuender/poker/pb"
)

type Start struct {
	bus  *Bus
	keys map[ebiten.Key]func()
}

func NewStart(bus *Bus) *Start {
	start := &Start{bus: bus}
	start.keys = map[ebiten.Key]func(){
		ebiten.KeyEscape: func() { bus.To(pb.Scene_desktop) },
	}

	return start
}

func (p *Start) Update() error               { return nil }
func (p *Start) Keys() map[ebiten.Key]func() { return p.keys }

func (p *Start) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "[ESC] run...")
}
