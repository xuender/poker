package mg

import (
	"fmt"
	"math/rand"
	"time"
)

// north | south | east | west.
type direction int

// Constants to signify which walls of a maze cell have been removed.
const (
	north = 1 << iota
	east
	south
	west
)

// Maps directions to Δx.
// nolint: gochecknoglobals
var dirX = map[direction]int{
	north: 0,
	east:  1,
	south: 0,
	west:  -1,
}

// Maps directions to Δy.
// nolint: gochecknoglobals
var dirY = map[direction]int{
	north: -1,
	east:  0,
	south: 1,
	west:  0,
}

// Opposite directions.
// nolint: gochecknoglobals
var Opposite = map[direction]direction{
	north: south,
	east:  west,
	south: north,
	west:  east,
}

// Cell is a single position in a Maze.
type Cell int

// Maze of N x M dimensions.
type Maze struct {
	width  int
	height int
	cells  [][]Cell
}

// NewMaze creates a new width x height Maze.
func NewMaze(width, height int) *Maze {
	maz := Maze{
		width,
		height,
		make([][]Cell, height),
	}

	for i := range maz.cells {
		maz.cells[i] = make([]Cell, width)
	}

	return &maz
}

func (maze *Maze) Generate() {
	rand.Seed(time.Now().UnixNano())
	maze.carvePassagesFrom(0, 0)
}

func between(potX, min, max int) bool {
	return (potX >= min && potX <= max)
}

// nolint: gosec
func shuffleDirections(slice []direction) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// carvePassagesFrom creates a maze starting from cell cx, cy.
func (maze *Maze) carvePassagesFrom(sizex, sizey int) {
	var (
		dir        direction
		directions = []direction{north, east, south, west}
	)

	shuffleDirections(directions)

	for i := range directions {
		dir = directions[i]
		newX, newY := sizex+dirX[dir], sizey+dirY[dir]

		if between(newX, 0, maze.width-1) && between(newY, 0, maze.height-1) && maze.cells[newY][newX] == 0 {
			// Hacky cast to Cell so that we can encode the carved walls
			maze.cells[sizey][sizex] |= Cell(dir)
			maze.cells[newY][newX] |= Cell(Opposite[dir])
			maze.carvePassagesFrom(newX, newY)
		}
	}
}

func (maze *Maze) isExit(x, y int) bool {
	return x == maze.width-1 && y == maze.height-1
}

// Pretty prints the Maze.
// nolint: forbidigo
func (maze *Maze) Print() {
	fmt.Print("  ") // 2 spaces, one for the left wall & one for the entrance

	for i := 0; i < maze.width*2-2; i++ {
		fmt.Print("_")
	}

	fmt.Print("\n")

	for top, row := range maze.cells {
		fmt.Print("|")

		for left, cell := range row {
			if cell&south != 0 || maze.isExit(left, top) {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}

			if cell&east != 0 {
				if (cell|row[left+1])&south != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("_")
				}
			} else {
				fmt.Print("|")
			}
		}

		fmt.Print("\n")
	}
}
