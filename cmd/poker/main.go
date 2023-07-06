package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/game"
)

type Poker struct {
	Bus    *game.Bus
	scenes []game.Scene
}

func NewPoker() *Poker {
	bus := game.NewBus()

	return &Poker{
		Bus: bus,
		scenes: []game.Scene{
			game.NewStart(bus),
			game.NewDesktop(bus),
			game.NewHelp(bus),
		},
	}
}

func (p *Poker) Layout(_, _ int) (int, int) { return p.Bus.Screen.X, p.Bus.Screen.Y }
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

func main() {
	flag.Usage = usage
	flag.Parse()

	poker := NewPoker()
	ebiten.SetWindowSize(poker.Bus.Screen.X, poker.Bus.Screen.Y)
	// ebiten.SetWindowSize(400, 300)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("自由扑克")
	lo.Must0(ebiten.RunGame(poker))
}

func usage() {
	fmt.Fprintf(os.Stderr, "poker\n\n")
	fmt.Fprintf(os.Stderr, "显示扑克.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
