package main

import "codeavdent2022/utils"

/**

--- Part Two ---
Your device's communication system is correctly detecting packets, but still isn't working. It looks like it also needs to look for messages.

A start-of-message marker is just like a start-of-packet marker, except it consists of 14 distinct characters rather than 4.

Here are the first positions of start-of-message markers for all of the above examples:

mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26
How many characters need to be processed before the first start-of-message marker is detected?

*/

func main() {
	input := utils.ReadContent("11.txt")
	arr := []rune(input)

	max := 14
	res := 0
	for i := 0; i < len(arr)-max; i++ {
		fl := arr[i : i+max]
		hash := make(map[rune]bool, 0)
		for _, f := range fl {
			hash[f] = true
		}
		if len(hash) == max {
			res = i + max
			break
		}
	}
	println(res)
}
