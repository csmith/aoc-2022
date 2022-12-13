package main

import (
	"sort"

	"github.com/csmith/aoc-2022/common"
)

func main() {
	pairs := common.ReadFileAsSectionedStrings("13/input.txt")
	part1 := 0
	var allPackets []item
	for p := range pairs {
		l := parse(pairs[p][0])
		r := parse(pairs[p][1])

		if compare(l, r) <= 0 {
			part1 += p + 1 // Stupid 1-based indexing
		}

		allPackets = append(allPackets, l, r)
	}

	divider1 := parse("[[2]]")
	divider2 := parse("[[6]]")
	divider1.divider = true
	divider2.divider = true
	allPackets = append(allPackets, divider1, divider2)

	sort.Slice(allPackets, func(i, j int) bool {
		return compare(allPackets[i], allPackets[j]) < 0
	})

	part2 := 1
	for i := range allPackets {
		if allPackets[i].divider {
			part2 *= i + 1
		}
	}

	println(part1)
	println(part2)
}

type item struct {
	isInt    bool
	val      int
	contents []item
	divider  bool
}

func parse(packet string) item {
	var stack []item

	for i := range packet {
		if packet[i] == '[' {
			stack = append(stack, item{})
		} else if packet[i] == ']' {
			// If we were parsing an int, shove it into its parent
			if stack[len(stack)-1].isInt {
				stack[len(stack)-2].contents = append(stack[len(stack)-2].contents, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 1 {
				// If this is the last thing on the stack we're done, return it
				return stack[0]
			} else {
				// Otherwise chop our list off the end and shove it into its parent
				stack[len(stack)-2].contents = append(stack[len(stack)-2].contents, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		} else if packet[i] == ',' {
			if stack[len(stack)-1].isInt {
				stack[len(stack)-2].contents = append(stack[len(stack)-2].contents, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		} else {
			if stack[len(stack)-1].isInt {
				stack[len(stack)-1].val = stack[len(stack)-1].val*10 + common.MustAtoi(string(packet[i]))
			} else {
				stack = append(stack, item{isInt: true, val: common.MustAtoi(string(packet[i]))})
			}
		}
	}

	panic("reached end of packet without finishing an item")
}

func compare(a, b item) int {
	if a.isInt && b.isInt {
		// If both values are integers, the lower integer should come first
		if a.val < b.val {
			return -1
		} else if a.val > b.val {
			return 1
		} else {
			return 0
		}
	} else if !a.isInt && !b.isInt {
		for i := 0; i < len(a.contents) && i < len(b.contents); i++ {
			if c := compare(a.contents[i], b.contents[i]); c != 0 {
				return c
			}
		}
		return len(a.contents) - len(b.contents)
	} else if a.isInt {
		return compare(item{contents: []item{a}}, b)
	} else {
		return compare(a, item{contents: []item{b}})
	}
}
