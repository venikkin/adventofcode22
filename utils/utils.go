package utils

import (
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
	i, e := strconv.Atoi(s)
	PanicOnErr(e)
	return i
}
