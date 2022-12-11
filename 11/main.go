package main

import (
	"fmt"
	"strings"

	"github.com/csmith/aoc-2022/common"
)

type monkey struct {
	inspected   int
	operation   func(int) int
	items       []int
	test        int
	nextIfTrue  *monkey
	nextIfFalse *monkey
}

func (m *monkey) ThrowStuff(divisor int, mod int) {
	for i := range m.items {
		v := (m.operation(m.items[i]) / divisor) % mod
		if v%m.test == 0 {
			m.nextIfTrue.items = append(m.nextIfTrue.items, v)
		} else {
			m.nextIfFalse.items = append(m.nextIfFalse.items, v)
		}
		m.inspected++
	}

	// Assuming we never throw an item to ourselves...
	m.items = m.items[:0]
}

func main() {
	input := common.ReadFileAsSectionedStrings("11/input.txt")
	monkeys := make([]*monkey, len(input))

	// Create all the monkeys, so we can get pointers to them
	for i := range input {
		monkeys[i] = &monkey{}
	}

	testLcm := initMonkeys(input, monkeys)
	println(simulate(monkeys, 20, 3, testLcm))

	resetMonkeys(input, monkeys)
	println(simulate(monkeys, 10000, 1, testLcm))
}

func initMonkeys(input [][]string, monkeys []*monkey) int {
	var tests []int64
	for i := range input {
		m := monkeys[i]
		m.inspected = 0
		m.operation = operation(strings.Split(input[i][2], " = ")[1])
		m.items = common.AtoiSlice(common.TrimSlice(strings.Split(strings.Split(input[i][1], ":")[1], ",")))
		m.test = common.MustAtoi(input[i][3][21:])
		m.nextIfTrue = monkeys[common.MustAtoi(input[i][4][29:])]
		m.nextIfFalse = monkeys[common.MustAtoi(input[i][5][30:])]
		tests = append(tests, int64(m.test))
	}
	return int(common.LCM(tests[0], tests[1], tests[2:]...))
}

func resetMonkeys(input [][]string, monkeys []*monkey) {
	for i := range input {
		m := monkeys[i]
		m.inspected = 0
		m.items = common.AtoiSlice(common.TrimSlice(strings.Split(strings.Split(input[i][1], ":")[1], ",")))
	}
}

func simulate(monkeys []*monkey, rounds, divisor, testLcm int) int {
	for i := 0; i < rounds; i++ {
		for m := range monkeys {
			monkeys[m].ThrowStuff(divisor, testLcm)
		}
	}

	// Calculate monkey business
	a, b := 0, 0
	for m := range monkeys {
		if monkeys[m].inspected > a {
			b = a
			a = monkeys[m].inspected
		} else if monkeys[m].inspected > b {
			b = monkeys[m].inspected
		}
	}
	return a * b
}

func operation(text string) func(int) int {
	if text == "old * old" {
		return func(i int) int {
			return i * i
		}
	} else if strings.HasPrefix(text, "old + ") {
		n := common.MustAtoi(text[6:])
		return func(i int) int {
			return i + n
		}
	} else if strings.HasPrefix(text, "old * ") {
		n := common.MustAtoi(text[6:])
		return func(i int) int {
			return i * n
		}
	}
	panic(fmt.Sprintf("unsupported operation: %s", text))
}
