package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/xuender/poker/cmd"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	ebiten.SetWindowTitle("自由扑克")
	lo.Must0(ebiten.RunGame(cmd.InitPoker()))
}

func usage() {
	fmt.Fprintf(os.Stderr, "poker\n\n")
	fmt.Fprintf(os.Stderr, "显示扑克.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
