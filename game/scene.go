package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(*ebiten.Image)
	Keys() map[ebiten.Key]func()
}
