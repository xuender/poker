package game

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xuender/poker/pb"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

type Bus struct {
	scenes []pb.Scene
	Screen image.Point
	keys   map[ebiten.Key]func()
}

func NewBus() *Bus {
	bus := &Bus{
		scenes: []pb.Scene{pb.Scene_start},
		Screen: image.Point{X: screenWidth, Y: screenHeight},
		keys:   map[ebiten.Key]func(){},
	}

	bus.keys[ebiten.KeyEscape] = bus.Close
	bus.keys[ebiten.KeyF2] = func() { bus.Pop(pb.Scene_help) }
	bus.keys[ebiten.KeyF11] = func() { ebiten.SetFullscreen(!ebiten.IsFullscreen()) }

	return bus
}

func (p *Bus) Keys() map[ebiten.Key]func() {
	return p.keys
}

func (p *Bus) Scenes() []pb.Scene {
	return p.scenes
}

func (p *Bus) Scene() pb.Scene {
	return p.scenes[len(p.scenes)-1]
}

func (p *Bus) To(scene pb.Scene) {
	p.scenes[len(p.scenes)-1] = scene
}

func (p *Bus) Pop(scene pb.Scene) {
	p.scenes = append(p.scenes, scene)
}

func (p *Bus) Close() {
	if len(p.scenes) <= 1 {
		os.Exit(0)
	}

	p.scenes = p.scenes[:len(p.scenes)-1]
}
