package main

import (
	"adventofcode22/utils"
	"fmt"
	"sort"
	"strings"
)

/**
--- Part Two ---
You're worried you might not ever get your items back. So worried, in fact, that your relief that a
monkey's inspection didn't damage an item no longer causes your worry level to be divided by three.

Unfortunately, that relief was all that was keeping your worry levels from reaching ridiculous levels.
You'll need to find another way to keep your worry levels manageable.

At this rate, you might be putting up with these monkeys for a very long time - possibly 10000 rounds!

With these new rules, you can still figure out the monkey business after 10000 rounds. Using the same example above:

== After round 1 ==
Monkey 0 inspected items 2 times.
Monkey 1 inspected items 4 times.
Monkey 2 inspected items 3 times.
Monkey 3 inspected items 6 times.

== After round 20 ==
Monkey 0 inspected items 99 times.
Monkey 1 inspected items 97 times.
Monkey 2 inspected items 8 times.
Monkey 3 inspected items 103 times.

== After round 1000 ==
Monkey 0 inspected items 5204 times.
Monkey 1 inspected items 4792 times.
Monkey 2 inspected items 199 times.
Monkey 3 inspected items 5192 times.

== After round 2000 ==
Monkey 0 inspected items 10419 times.
Monkey 1 inspected items 9577 times.
Monkey 2 inspected items 392 times.
Monkey 3 inspected items 10391 times.

== After round 3000 ==
Monkey 0 inspected items 15638 times.
Monkey 1 inspected items 14358 times.
Monkey 2 inspected items 587 times.
Monkey 3 inspected items 15593 times.

== After round 4000 ==
Monkey 0 inspected items 20858 times.
Monkey 1 inspected items 19138 times.
Monkey 2 inspected items 780 times.
Monkey 3 inspected items 20797 times.

== After round 5000 ==
Monkey 0 inspected items 26075 times.
Monkey 1 inspected items 23921 times.
Monkey 2 inspected items 974 times.
Monkey 3 inspected items 26000 times.

== After round 6000 ==
Monkey 0 inspected items 31294 times.
Monkey 1 inspected items 28702 times.
Monkey 2 inspected items 1165 times.
Monkey 3 inspected items 31204 times.

== After round 7000 ==
Monkey 0 inspected items 36508 times.
Monkey 1 inspected items 33488 times.
Monkey 2 inspected items 1360 times.
Monkey 3 inspected items 36400 times.

== After round 8000 ==
Monkey 0 inspected items 41728 times.
Monkey 1 inspected items 38268 times.
Monkey 2 inspected items 1553 times.
Monkey 3 inspected items 41606 times.

== After round 9000 ==
Monkey 0 inspected items 46945 times.
Monkey 1 inspected items 43051 times.
Monkey 2 inspected items 1746 times.
Monkey 3 inspected items 46807 times.

== After round 10000 ==
Monkey 0 inspected items 52166 times.
Monkey 1 inspected items 47830 times.
Monkey 2 inspected items 1938 times.
Monkey 3 inspected items 52013 times.
After 10000 rounds, the two most active monkeys inspected items 52166 and 52013 times.
Multiplying these together, the level of monkey business in this situation is now 2713310158.

Worry levels are no longer divided by three after each item is inspected; you'll need to find another
way to keep your worry levels manageable. Starting again from the initial state in your puzzle input, what is the level of monkey business after 10000 rounds?
*/

type monkey struct {
	items        []int
	operation    func(int) int
	divisor      int
	positiveDir  int
	negativeDir  int
	inspectedCnt int
}

func (m *monkey) popAll() []int {
	i := m.items
	m.items = make([]int, 0)
	m.inspectedCnt += len(i)
	return i
}

func (m *monkey) push(i int) {
	m.items = append(m.items, i)
}

func main() {
	lines := utils.ReadLines("21.txt")
	monkeys := make([]monkey, 0)

	for i := 0; i < len(lines); i += 7 {
		items := utils.ToIntSlice(lines[i+1][len("  Starting items: "):])

		operationLine := lines[i+2][len("  Operation: new = old "):]
		operationNumber := strings.TrimSpace(operationLine[2:])
		var operation func(int) int
		if string(operationLine[0]) == "*" {
			operation = func(i int) int { return i * convertToNumber(i, operationNumber) }
		} else if string(operationLine[0]) == "+" {
			operation = func(i int) int { return i + convertToNumber(i, operationNumber) }
		} else if string(operationLine[0]) == "-" {
			operation = func(i int) int { return i - convertToNumber(i, operationNumber) }
		} else {
			operation = func(i int) int { return i / convertToNumber(i, operationNumber) }
		}

		divisor := utils.AtoiOrPanic(lines[i+3][len("  Test: divisible by "):])
		positiveDir := utils.AtoiOrPanic(lines[i+4][len("    If true: throw to monkey "):])
		negativeDir := utils.AtoiOrPanic(lines[i+5][len("    If false: throw to monkey "):])
		monkeys = append(monkeys, monkey{
			items:        items,
			operation:    operation,
			divisor:      divisor,
			negativeDir:  negativeDir,
			positiveDir:  positiveDir,
			inspectedCnt: 0,
		})
	}

	commonDivisor := 1
	for _, monkey := range monkeys {
		commonDivisor *= monkey.divisor
	}
	println(commonDivisor)

	fmt.Printf("monkeys: %+v\n", monkeys)
	for i := 0; i < 10_000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := monkeys[j]
			for _, item := range monkey.popAll() {
				item = item % commonDivisor
				converted := monkey.operation(item)
				var direction int
				if converted%monkey.divisor == 0 {
					direction = monkey.positiveDir
				} else {
					direction = monkey.negativeDir
				}

				monkeys[direction].push(converted)
			}
			monkeys[j] = monkey
		}
	}

	fmt.Printf("monkeys: %+v\n", monkeys)
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCnt > monkeys[j].inspectedCnt
	})

	fmt.Printf("monkeys: %+v\n", monkeys)
	println(monkeys[0].inspectedCnt * monkeys[1].inspectedCnt)
}

func convertToNumber(i int, operationNumber string) int {
	if operationNumber == "old" {
		return i
	}
	return utils.AtoiOrPanic(operationNumber)
}
