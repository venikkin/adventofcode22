package main

import (
	"adventofcode22/utils"
	"fmt"
	"strings"
)

/**
--- Day 12: Hill Climbing Algorithm ---
You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

You ask the device for a heightmap of the surrounding area (your puzzle input).
The heightmap shows the local area from above broken into a grid; the elevation of each square of the grid is
given by a single lowercase letter, where a is the lowest elevation, b is the next-lowest, and so on up to the highest elevation, z.

Also included on the heightmap are marks for your current position (S) and the location that should get the best signal (E).
Your current position (S) has elevation a, and the location that should get the best signal (E) has elevation z.

You'd like to reach E, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly
one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be at
most one higher than the elevation of your current square; that is, if your current elevation is m, you could step to elevation n,
but not to elevation o. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

For example:

Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but
eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:

v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^
In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left (<), or right (>).
The location that should get the best signal is still E, and . marks unvisited squares.

This path reaches the goal in 31 steps, the fewest possible.

What is the fewest steps required to move from your current position to the location that should get the best signal?
*/

// https://en.wikipedia.org/wiki/Breadth-first_search
func main() {

	landscape := make([][]rune, 0)
	for _, line := range utils.ReadLines("23-test.txt") {
		landscape = append(landscape, []rune(strings.TrimSpace(line)))
	}

	var start point
	for i := 0; i < len(landscape); i++ {
		for j := 0; j < len(landscape[i]); j++ {
			if string(landscape[i][j]) == "S" {
				start = point{j, i}
			}
		}
	}

	min := move(landscape, start, map[point]bool{}, 0)
	println(min)
}

type point struct {
	x, y int
}

func (p point) add(dx, dy int) point {
	return point{p.x + dx, p.y + dy}
}

func move(land [][]rune, start point, visited map[point]bool, step int) int {
	startRune := land[start.y][start.x]
	if string(startRune) == "E" {
		return step - 1
	}
	if visited[start] {
		return 0
	}
	visited[start] = true
	moves := make([]int, 0)
	// up
	nextUp := start.add(0, -1)
	if start.y > 0 && canMove(startRune, land[start.y-1][start.x]) {
		up := move(land, nextUp, visited, step+1)
		if up != 0 {
			moves = append(moves, up)
		}
	}
	// down
	nextDown := start.add(0, 1)
	if start.y < len(land)-1 && canMove(startRune, land[start.y+1][start.x]) {
		down := move(land, nextDown, visited, step+1)
		if down != 0 {
			moves = append(moves, down)
		}
	}
	// left
	nextLeft := start.add(-1, 0)
	if start.x > 0 && canMove(startRune, land[start.y][start.x-1]) {
		left := move(land, nextLeft, visited, step+1)
		if left != 0 {
			moves = append(moves, left)
		}
	}
	// right
	nextRight := start.add(1, 0)
	if start.x < len(land[0])-1 && canMove(startRune, land[start.y][start.x+1]) {
		right := move(land, nextRight, visited, step+1)
		if right != 0 {
			moves = append(moves, right)
		}
	}

	if len(moves) == 0 {
		return 0
	}

	min := utils.Min(moves...)
	fmt.Printf("%+v -- %d\n", start, min)
	return min
}

func canMove(s, e rune) bool {
	st := s
	if string(st) == "S" {
		st = []rune("a")[0]
	}
	en := e
	if string(e) == "E" {
		en = []rune("z")[0]
	}
	return en <= st+1
}
