package game

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xuender/poker/pb"
)

const (
	_screenWidth  = 1024
	_screenHeight = 768
)

type Bus struct {
	scenes []pb.Scene
	keys   map[ebiten.Key]func()
}

func NewBus() *Bus {
	ebiten.SetWindowSize(_screenWidth, _screenHeight)

	bus := &Bus{
		scenes: []pb.Scene{pb.Scene_start},
		keys:   map[ebiten.Key]func(){},
	}

	bus.keys[ebiten.KeyEscape] = bus.Close
	bus.keys[ebiten.KeyF2] = func() { bus.Pop(pb.Scene_help) }
	bus.keys[ebiten.KeyF11] = func() { ebiten.SetFullscreen(!ebiten.IsFullscreen()) }

	return bus
}

func (p *Bus) Layout() (int, int) {
	return _screenWidth, _screenHeight
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
