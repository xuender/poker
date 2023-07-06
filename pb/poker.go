package pb

import (
	"embed"
	"fmt"

	"github.com/samber/lo"
)

var (
	//go:embed pokers
	_pokers embed.FS
	// nolint: gochecknoglobals
	_data = make([][]byte, len(Poker_name))
)

// nolint: gochecknoinits
func init() {
	for key := range Poker_name {
		_data[key] = lo.Must1(_pokers.ReadFile(fmt.Sprintf("pokers/%d.png", key)))
	}
}

func (p Poker) Bytes() []byte {
	return _data[p]
}
