package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToggle(t *testing.T) {
	mf := minefield{
		height: 2,
		width:  2,
		cells: [][]cell{
			{cell{}, cell{}},
			{cell{}, cell{}},
		},
	}

	mf.ToggleMark(2, 2)

	assert.Equal(t, true, mf.cells[1][1].isMarked)
}

func TestToggleVisible(t *testing.T) {
	mf := minefield{
		height: 2,
		width:  2,
		cells: [][]cell{
			{cell{}, cell{}},
			{cell{}, cell{isVisible: true}},
		},
	}

	mf.ToggleMark(2, 2)

	assert.Equal(t, false, mf.cells[1][1].isMarked)
}

func TestClick(t *testing.T) {
	mf := minefield{
		height: 4,
		width:  4,
		cells: [][]cell{
			{cell{isMine: true}, cell{}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
		},
	}

	result := mf.Click(1, 1)

	assert.Equal(t, true, result)
}

func TestClickSafe(t *testing.T) {
	mf := minefield{
		height: 4,
		width:  4,
		cells: [][]cell{
			{cell{isMine: true}, cell{cntAdjacent: 1}, cell{}, cell{}},
			{cell{cntAdjacent: 1}, cell{cntAdjacent: 1}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
		},
	}

	result := mf.Click(4, 4)

	assert.Equal(t, false, result)

	expectedCells := [][]cell{
		{cell{isMine: true}, cell{cntAdjacent: 1, isVisible: true}, cell{isVisible: true}, cell{isVisible: true}},
		{cell{cntAdjacent: 1, isVisible: true}, cell{cntAdjacent: 1, isVisible: true}, cell{isVisible: true}, cell{isVisible: true}},
		{cell{isVisible: true}, cell{isVisible: true}, cell{isVisible: true}, cell{isVisible: true}},
		{cell{isVisible: true}, cell{isVisible: true}, cell{isVisible: true}, cell{isVisible: true}},
	}

	assert.Equal(t, expectedCells, mf.cells)
}

func TestClickAdjacent(t *testing.T) {
	mf := minefield{
		height: 4,
		width:  4,
		cells: [][]cell{
			{cell{isMine: true}, cell{cntAdjacent: 1}, cell{}, cell{}},
			{cell{cntAdjacent: 1}, cell{cntAdjacent: 1}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
		},
	}

	result := mf.Click(2, 2)

	assert.Equal(t, false, result)

	expectedCells := [][]cell{
		{cell{isMine: true}, cell{cntAdjacent: 1}, cell{}, cell{}},
		{cell{cntAdjacent: 1}, cell{cntAdjacent: 1, isVisible: true}, cell{}, cell{}},
		{cell{}, cell{}, cell{}, cell{}},
		{cell{}, cell{}, cell{}, cell{}},
	}

	assert.Equal(t, expectedCells, mf.cells)
}

func TestClickPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	mf := minefield{
		height: 4,
		width:  4,
		cells: [][]cell{
			{cell{isMine: true}, cell{cntAdjacent: 1}, cell{}, cell{}},
			{cell{cntAdjacent: 1}, cell{cntAdjacent: 1}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
			{cell{}, cell{}, cell{}, cell{}},
		},
	}

	mf.Click(5, 5)
}

func TestString(t *testing.T) {
	mf := minefield{
		height: 4,
		width:  4,
		cells: [][]cell{
			{cell{isMine: true, isMarked: true}, cell{cntAdjacent: 2, isVisible: true}, cell{}},
			{cell{isMine: true, isVisible: true}, cell{cntAdjacent: 2}, cell{}},
			{cell{cntAdjacent: 1}, cell{cntAdjacent: 1}, cell{}},
		},
	}

	result := mf.String()

	expected := `* 2 - 
x - - 
- - - 
`

	assert.Equal(t, expected, result)
}
