package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

type Poker struct {
	Bus    *Bus
	scenes []Scene
}

func NewPoker(bus *Bus, start *StartScene, desktop *DesktopScene, help *HelpScene) *Poker {
	return &Poker{
		Bus:    bus,
		scenes: []Scene{start, desktop, help},
	}
}

func (p *Poker) Layout(_, _ int) (int, int) { return p.Bus.Layout() }
func (p *Poker) Update() error {
	scene := p.scenes[p.Bus.Scene()]
	keys := inpututil.AppendJustPressedKeys(nil)

	if len(keys) > 0 {
		logs.D.Println(p.Bus.Scene(), keys)

		runBus := true

		for key, fun := range scene.Keys() {
			if lo.Contains(keys, key) {
				runBus = false

				fun()
			}
		}

		if runBus {
			for key, fun := range p.Bus.Keys() {
				if lo.Contains(keys, key) {
					fun()
				}
			}
		}
	}

	return scene.Update()
}

func (p *Poker) Draw(screen *ebiten.Image) {
	for _, before := range p.Bus.Scenes() {
		p.scenes[before].Draw(screen)
	}
}
