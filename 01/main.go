package main

import "github.com/csmith/aoc-2022/common"

func main() {
	elves := common.ReadFileAsSectionedStrings("01/input.txt")
	alpha, beta, gamma := 0, 0, 0
	for i := range elves {
		s := sum(elves[i])
		if s > alpha {
			gamma = beta
			beta = alpha
			alpha = s
		} else if s > beta {
			gamma = beta
			beta = s
		} else if s > gamma {
			gamma = s
		}
	}
	println(alpha)
	println(alpha + beta + gamma)
}

func sum(calories []string) int {
	res := 0
	for i := range calories {
		res += common.MustAtoi(calories[i])
	}
	return res
}
