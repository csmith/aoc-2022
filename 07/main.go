package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/csmith/aoc-2022/common"
)

const cdPrefix = "$ cd "

func main() {
	input := common.ReadFileAsStrings("07/input.txt")
	files := make(map[string]int)
	cwd := "/"

	// First build a map of files, so we don't over-count any duplicate `ls` results
	var size int
	var name string
	for i := range input {
		if strings.HasPrefix(input[i], cdPrefix) {
			arg := strings.TrimPrefix(input[i], cdPrefix)
			if strings.HasPrefix(arg, "/") {
				cwd = arg
			} else {
				cwd = filepath.Clean(filepath.Join(cwd, arg))
			}
		} else if n, _ := fmt.Sscanf(input[i], "%d %s", &size, &name); n == 2 {
			files[filepath.Join(cwd, name)] = size
		}
	}

	// Now calculate the total size of each directory
	dirs := make(map[string]int)
	for f := range files {
		dir := filepath.Dir(f)
		for {
			dirs[dir] += files[f]
			if dir == "/" {
				break
			}
			dir = filepath.Dir(dir)
		}
	}

	// And finally do weird stuff to get answers
	sum := 0
	target := dirs["/"] - 40000000
	closest := dirs["/"]
	for d := range dirs {
		if dirs[d] < 100000 {
			sum += dirs[d]
		}
		if dirs[d] > target && dirs[d] < closest {
			closest = dirs[d]
		}
	}

	println(sum)
	println(closest)
}
