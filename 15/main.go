package main

import (
	"adventofcode22/utils"
	"strings"
)

/*
--- Day 8: Treetop Tree House ---
The expedition comes across a peculiar patch of tall trees all planted carefully in a grid.
The Elves explain that a previous expedition planted these trees as a reforestation effort.
Now, they're curious if this would be a good location for a tree house.

First, determine whether there is enough tree cover here to keep a tree house hidden.
To do this, you need to count the number of trees that are visible from outside the grid when looking directly along a row or column.

The Elves have already launched a quadcopter to generate a map with the height of each tree (your puzzle input). For example:

30373
25512
65332
33549
35390
Each tree is represented as a single digit whose value is its height, where 0 is the shortest and 9 is the tallest.

A tree is visible if all of the other trees between it and an edge of the grid are shorter than it.
Only consider trees in the same row or column; that is, only look up, down, left, or right from any given tree.

All of the trees around the edge of the grid are visible - since they are already on the edge, there
are no trees to block the view. In this example, that only leaves the interior nine trees to consider:

The top-left 5 is visible from the left and top. (It isn't visible from the right or bottom since other trees of height 5 are in the way.)
The top-middle 5 is visible from the top and right.
The top-right 1 is not visible from any direction; for it to be visible, there would need to only be trees of height 0 between it and an edge.
The left-middle 5 is visible, but only from the right.
The center 3 is not visible from any direction; for it to be visible, there would need to be only trees of at most height 2 between it and an edge.
The right-middle 3 is visible from the right.
In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
With 16 trees visible on the edge and another 5 visible in the interior, a total of 21 trees are visible in this arrangement.

Consider your map; how many trees are visible from outside the grid?
*/

func main() {
	trees := make([][]int, 0)
	for _, line := range utils.ReadLines("15.txt") {
		chars := []rune(strings.TrimSpace(line))
		row := make([]int, len(chars))
		for i := 0; i < len(chars); i++ {
			row[i] = utils.AtoiOrPanic(string(chars[i]))
		}
		trees = append(trees, row)
	}
	//fmt.Printf("trees: %+v\n", trees)
	maxRows := len(trees[0])
	maxCols := len(trees)
	println(maxCols, "x", maxRows)
	outer := 2*maxRows + 2*maxCols - 4

	inner := 0
	for i := 1; i < maxCols-1; i++ {
		for j := 1; j < maxRows-1; j++ {
			if isVisible(i, j, trees) {
				inner++
			}
		}
	}

	println("outer: ", outer)
	println("inner: ", inner)
	println("total: ", outer+inner)
}

func isVisible(col, row int, trees [][]int) bool {
	val := trees[col][row]

	left := true
	for i := col - 1; i >= 0; i-- {
		if trees[i][row] >= val {
			left = false
			break
		}
	}
	if left {
		return true
	}

	right := true
	for i := col + 1; i < len(trees[0]); i++ {
		if trees[i][row] >= val {
			right = false
			break
		}
	}
	if right {
		return true
	}

	top := true
	for i := row - 1; i >= 0; i-- {
		if trees[col][i] >= val {
			top = false
			break
		}
	}
	if top {
		return true
	}

	bottom := true
	for i := row + 1; i < len(trees); i++ {
		if trees[col][i] >= val {
			bottom = false
			break
		}
	}

	return bottom
}
