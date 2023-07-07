package fonts

import (
	_ "embed"

	"github.com/samber/lo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// nolint: gochecknoglobals
var (
	//go:embed HanZiZhiMeiShenYongTuShengXiao-Shan(God-ShenYongTuGB-Flash)-2.ttf
	_headTtf []byte
	//go:embed PangMenZhengDaoBiaoTiTi-1.ttf
	_bodyTtf  []byte
	_headFont *opentype.Font
	_bodyFont *opentype.Font
)

const _dpi = 72

// nolint: gochecknoinits
func init() {
	_headFont = lo.Must1(opentype.Parse(_headTtf))
	_bodyFont = lo.Must1(opentype.Parse(_bodyTtf))
}

func Head(size float64) font.Face {
	return lo.Must1(opentype.NewFace(_headFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     _dpi,
		Hinting: font.HintingVertical,
	}))
}

func Body(size float64) font.Face {
	return lo.Must1(opentype.NewFace(_bodyFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     _dpi,
		Hinting: font.HintingVertical,
	}))
}
