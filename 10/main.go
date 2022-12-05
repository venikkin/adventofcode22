package main

import (
	"codeavdent2022/utils"
	"fmt"
	"regexp"
	"strings"
)

/**
--- Part Two ---
As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air conditioning,
leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

Again considering the example above, the crates begin in the same configuration:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3
Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3
However, the action of moving three crates from stack 1 to stack 3 means that those
three moved crates stay in the same order, resulting in this new configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3
Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3
Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3
In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

Before the rearrangement process finishes, update your simulation so that the Elves know where
they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?


*/

func main() {
	lines := utils.ReadLines("09.txt")
	div := emptyString(lines)

	cols := len(lines[div-1])/4 + 1
	crates := make([][]rune, cols)

	for i := div - 2; i >= 0; i-- {
		runes := []rune(lines[i])
		for j := 1; j < len(runes); j += 4 {
			bucket := j / 4
			if string(runes[j]) != " " {
				crates[bucket] = append(crates[bucket], runes[j])
			}
		}
	}
	printByLine(crates)

	commandReg := regexp.MustCompile("^move (\\d+) from (\\d+) to (\\d+)$")
	for _, command := range lines[div+1:] {
		println(command)
		cmd := commandReg.FindStringSubmatch(command)
		amount := utils.AtoiOrPanic(cmd[1])
		target := utils.AtoiOrPanic(cmd[2]) - 1
		destination := utils.AtoiOrPanic(cmd[3]) - 1
		fmt.Printf("%v -(%v)-> %v\n", target, amount, destination)

		crates[destination] = append(crates[destination], crates[target][len(crates[target])-amount:]...)
		crates[target] = crates[target][:len(crates[target])-amount]
		printByLine(crates)
	}

	println()
	printByLine(crates)

	res := ""
	for _, col := range crates {
		res += string(col[len(col)-1])
	}
	println(res)
}

func emptyString(s []string) int {
	for i, l := range s {
		if strings.TrimSpace(l) == "" {
			return i
		}
	}
	panic("cannot find schema")
}

func printByLine(rr [][]rune) {
	for i, ra := range rr {
		print(i, ": ")
		for _, r := range ra {
			print(string(r), ", ")
		}
		println()
	}
}
