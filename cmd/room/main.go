package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/samber/lo"
	"github.com/xuender/kit/base"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

type Room struct {
	texts []string
}

func NewRoom() *Room {
	return &Room{
		texts: []string{"test", "aaa"},
	}
}

func (p *Room) Update() error              { return nil }
func (p *Room) Layout(_, _ int) (int, int) { return screenWidth, screenHeight }
func (p *Room) Draw(screen *ebiten.Image) {
	for index, msg := range p.texts {
		ebitenutil.DebugPrintAt(screen, msg, 0, index*base.Hundred)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("聊天室")
	lo.Must0(ebiten.RunGame(NewRoom()))
}

func usage() {
	fmt.Fprintf(os.Stderr, "room\n\n")
	fmt.Fprintf(os.Stderr, "聊天室.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
