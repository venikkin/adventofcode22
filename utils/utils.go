package utils

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadLines(file string) []string {
	content := ReadContent(file)
	return strings.Split(content, "\n")
}

func ReadContent(file string) string {
	cwd, err := os.Getwd()
	PanicOnErr(err)

	cnt, err := os.ReadFile(cwd + "/inputs/" + file)
	PanicOnErr(err)

	return string(cnt)
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func AtoiOrPanic(s string) int {
	i, e := strconv.Atoi(strings.TrimSpace(s))
	PanicOnErr(e)
	return i
}

func IntAbs(i int) int {
	return int(math.Abs(float64(i)))
}

func ToIntSlice(s string) []int {
	elems := strings.Split(s, ",")
	res := make([]int, len(elems))
	for i, elem := range elems {
		res[i] = AtoiOrPanic(strings.TrimSpace(elem))
	}
	return res
}

func Min(args ...int) int {
	min := math.MaxInt
	for _, i := range args {
		if i < min {
			min = i
		}
	}
	return min
}
