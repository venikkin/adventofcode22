package utils

import (
	"os"
	"strings"
)

func ReadLines(file string) []string {
	cwd, err := os.Getwd()
	PanicOnErr(err)

	cnt, err := os.ReadFile(cwd + "/inputs/" + file)
	PanicOnErr(err)

	content := string(cnt)
	return strings.Split(content, "\n")
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
