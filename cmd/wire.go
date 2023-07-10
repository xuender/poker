//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/xuender/poker/game"
)

func InitPoker() *game.Poker {
	wire.Build(
		game.NewPoker,
		game.NewBus,
		game.NewDesktop,
		game.NewHelp,
		game.NewStart,
		game.NewScript,
	)

	return &game.Poker{}
}
