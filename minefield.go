package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type cell struct {
	isMine      bool
	isVisible   bool
	isMarked    bool
	cntAdjacent int8
}

type coordinate struct {
	x int
	y int
}

type minefield struct {
	height, width int
	cells         [][]cell
}

// NewMinefield generates new minefield
func NewMinefield(height, width, mineAmount int) minefield {
	mf := minefield{height: height, width: width}

	// init cells
	mf.cells = make([][]cell, height)
	for i := 0; i < height; i++ {
		mf.cells[i] = make([]cell, width)
	}

	// generate mines
	for i := 0; i < mineAmount; i++ {
		mx, my := rand.Intn(mf.height), rand.Intn(mf.width)

		// try again if mine with such coordinate already exist
		if mf.cells[mx][my].isMine {
			i--
			continue
		}

		mf.cells[mx][my].isMine = true
		mf.handleNeighbours(mx, my, func(_, _ int, c *cell) {
			c.cntAdjacent++
		})
	}

	return mf
}

// ToggleMark marks cell as potential bomb
func (mf minefield) ToggleMark(x, y int) {
	// visible cells can not be marked as bomb
	if mf.cells[x-1][y-1].isVisible {
		return
	}

	mf.cells[x-1][y-1].isMarked = !mf.cells[x-1][y-1].isMarked
}

// Click checks the given coordinate and return true if it is a bomb
// otherwise mark field as visible
// if no bomb is nearby, it marks all adjacent safe cells as visible
func (mf minefield) Click(x, y int) bool {
	// given coordinates must be in minefield range
	if x < 1 || y < 1 || x > mf.height || y > mf.width {
		panic(fmt.Sprintf("given coordinates out of minefield range x: %d y: %d", x, y))
	}

	// You are unlucky
	if mf.cells[x-1][y-1].isMine {
		return true // boom
	}

	// mark all adjacent safe cells as visible

	cellsToProcess := []coordinate{
		{x: x - 1, y: y - 1},
	}

	for i := 0; i < len(cellsToProcess); i++ {
		c := cellsToProcess[i]
		mf.cells[c.x][c.y].isVisible = true

		// one of the adjacent cells is a bomb, do not mark neighbors as visible
		if mf.cells[c.x][c.y].cntAdjacent > 0 {
			continue
		}

		// otherwise add to the queue neighbors that not visible yet
		// TODO: (allocates new func every iteration) move callback function initialization outside of the cycle
		mf.handleNeighbours(c.x, c.y, func(x, y int, c *cell) {
			if !c.isVisible { // visible cells already processed so no need adding them to the queue
				cellsToProcess = append(cellsToProcess, coordinate{x, y})
			}
		})
	}

	return false
}

// handleNeighbours iterates all adjacent cells including at given coordinates
// runs callback function for each adjacent cells
func (mf minefield) handleNeighbours(x, y int, callback func(x, y int, c *cell)) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			mx := x + i
			my := y + j

			if mx < 0 || mx > mf.height-1 {
				continue
			}

			if my < 0 || my > mf.width-1 {
				continue
			}

			callback(mx, my, &mf.cells[mx][my])
		}
	}
}

func (mf minefield) String() string {
	builder := strings.Builder{}

	for x := range mf.cells {
		for y := range mf.cells[x] {
			if mf.cells[x][y].isMarked {
				builder.WriteByte('*')
			} else if !mf.cells[x][y].isVisible {
				builder.WriteByte('-')
			} else if mf.cells[x][y].isMine {
				builder.WriteByte('x')
			} else {
				builder.WriteString(strconv.Itoa(int(mf.cells[x][y].cntAdjacent)))
			}

			builder.WriteByte(' ')
		}

		builder.WriteByte('\n')
	}

	return builder.String()
}
