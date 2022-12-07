package main

import (
	"adventofcode22/utils"
	"sort"
	"strings"
)

/*
--- Part Two ---
Now, you're ready to choose a directory to delete.

The total disk space available to the filesystem is 70000000. To run the update,
you need unused space of at least 30000000. You need to find a directory you can delete that will free up enough space to run the update.

In the example above, the total size of the outermost directory (and thus the total
amount of used space) is 48381165; this means that the size of the unused space must currently be 21618835, which isn't quite the 30000000 required by the update. Therefore, the update still requires a directory with total size of at least 8381165 to be deleted before it can run.

To achieve this, you have the following options:

Delete directory e, which would increase unused space by 584.
Delete directory a, which would increase unused space by 94853.
Delete directory d, which would increase unused space by 24933642.
Delete directory /, which would increase unused space by 48381165.
Directories e and a are both too small; deleting them would not free up enough space.
However, directories d and / are both big enough! Between these, choose the smallest: d, increasing unused space by 24933642.

Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?
*/

func main() {

	root := file{name: "/", size: 0, files: make(map[string]file, 0), folder: true}
	cursor := root
	for _, line := range utils.ReadLines("13.txt") {
		println(line)
		if strings.HasPrefix(line, "$ cd") {
			folderName := line[5:]
			if folderName == "/" {
				cursor = root
			} else if folderName == ".." {
				//println("-- going from " + cursor.name + " to " + cursor.parent.name)
				cursor = *cursor.parent
			} else {
				cursorL, ok := cursor.files[folderName]
				if !ok {
					panic("fail to find " + folderName + " parent: " + cursor.name)
				}
				cursor = cursorL
			}
		} else if strings.HasPrefix(line, "$ ls") {
			// nothing so far
		} else {
			meta := strings.Split(line, " ")
			size := 0
			folder := true
			if meta[0] != "dir" {
				size = utils.AtoiOrPanic(meta[0])
				folder = false
			}
			name := meta[1]
			cur := cursor
			f := file{name: name, size: size, parent: &cur, files: make(map[string]file, 0), folder: folder}
			cursor.files[name] = f
		}
	}

	total := 70_000_000
	used := sumF(root)
	free := total - used

	smallFolders := make([]int, 0)
	findSmall(root, &smallFolders, 30_000_000-free)

	sort.Slice(smallFolders, func(i, j int) bool {
		return smallFolders[i] < smallFolders[j]
	})
	println(smallFolders[0])
}

func sumF(file file) int {
	sum := 0
	if file.folder {
		for _, f := range file.files {
			if f.folder {
				sum += sumF(f)
			} else {
				sum += f.size
			}
		}
	}
	return sum
}

func findSmall(file file, acc *[]int, need int) int {
	sum := 0
	if file.folder {
		for _, f := range file.files {
			if f.folder {
				sum += findSmall(f, acc, need)
			} else {
				sum += f.size
			}
		}
		if sum >= need {
			*acc = append(*acc, sum)
		}
	}
	return sum
}

type file struct {
	name   string
	size   int
	files  map[string]file
	parent *file
	folder bool
}
