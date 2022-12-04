package main

import (
	"codeavdent2022/utils"
	"strconv"
	"strings"
)

/*
--- Part Two ---
It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.

In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs
(5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:

5-7,7-9 overlaps in a single section, 7.
2-8,3-7 overlaps all of the sections 3 through 7.
6-6,4-6 overlaps in a single section, 6.
2-6,4-8 overlaps in sections 4, 5, and 6.
So, in this example, the number of overlapping assignment pairs is 4.

In how many assignment pairs do the ranges overlap?
*/

func main() {

	lines := utils.ReadLines("07.txt")

	res := 0
	for _, line := range lines {
		//println(line)
		thisPair := parseLine(line)

		include := false
		if overlap(thisPair[0], thisPair[1]) {
			include = true
		}
		if include {
			res++
		}
	}
	println(res)
}

func parseLine(line string) [][]int {
	elfs := strings.Split(line, ",")
	elf1 := strings.Split(elfs[0], "-")
	elf2 := strings.Split(elfs[1], "-")

	return [][]int{
		{atoiOrPanic(elf1[0]), atoiOrPanic(elf1[1])},
		{atoiOrPanic(elf2[0]), atoiOrPanic(elf2[1])},
	}
}

func atoiOrPanic(s string) int {
	i, e := strconv.Atoi(s)
	utils.PanicOnErr(e)
	return i
}

func overlap(a, b []int) bool {
	return (a[0] <= b[1] && a[1] >= b[0]) ||
		(b[0] <= a[1] && b[1] >= a[0])
}
