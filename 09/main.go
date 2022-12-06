package main

import (
	"adventofcode22/utils"
	"fmt"
	"regexp"
	"strings"
)

/**
--- Day 5: Supply Stacks ---
The expedition can depart as soon as the final supplies have been unloaded from the ships.
Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks.
To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps.
After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure,
but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top.
Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of
crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure,
one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3
In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time,
so the first crate to be moved (D) ends up below the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3
Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3
Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3
The Elves just need to know which crate will end up on top of each stack; in this example, the top
crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?
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

		for i := 0; i < amount; i++ {
			crates[destination] = append(crates[destination], crates[target][len(crates[target])-1])
			crates[target] = crates[target][:len(crates[target])-1]
		}
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
