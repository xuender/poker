package game

import (
	"github.com/samber/lo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	mobile "golang.org/x/mobile/exp/font"
)

const _dpi = 72

type Fonts struct {
	defaultFont *opentype.Font
	monospace   *opentype.Font
}

func NewFonts() *Fonts {
	return &Fonts{
		defaultFont: lo.Must1(opentype.Parse(mobile.Default())),
		monospace:   lo.Must1(opentype.Parse(mobile.Monospace())),
	}
}

func (p *Fonts) DefaultFace(size float64) font.Face {
	return lo.Must1(opentype.NewFace(p.defaultFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     _dpi,
		Hinting: font.HintingVertical,
	}))
}

func (p *Fonts) MonospaceFace(size float64) font.Face {
	return lo.Must1(opentype.NewFace(p.monospace, &opentype.FaceOptions{
		Size:    size,
		DPI:     _dpi,
		Hinting: font.HintingVertical,
	}))
}
