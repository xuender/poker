// Implements a recursive backtracking maze generation algorithm.
// Based upon http://weblog.jamisbuck.org/2010/12/27/maze-generation-recursive-backtracking
package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/cmd/mg/mg"
)

const usage = "Usage: ./recursive-backtracking <width> <height>"

type dimensions struct {
	width  int
	height int
}

func main() {
	args := os.Args[1:]

	dim, err := parseArgs(args)
	if err != nil {
		logs.E.Println(err)
		logs.I.Println(usage)
		os.Exit(1)
	}

	maze := mg.NewMaze(dim.width, dim.height)
	maze.Generate()
	maze.Print()
}

var (
	ErrWidth  = errors.New("width must be an integer")
	ErrHeight = errors.New("height must be an integer")
	ErrArg    = errors.New("2 arguments required")
)

func parseArgs(args []string) (dimensions, error) {
	var (
		dim dimensions
		err error
	)

	if len(args) < base.Two {
		return dim, ErrArg
	}

	dim.width, err = strconv.Atoi(args[0])
	if err != nil {
		return dim, ErrWidth
	}

	dim.height, err = strconv.Atoi(args[1])
	if err != nil {
		return dim, ErrHeight
	}

	return dim, nil
}
