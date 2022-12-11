package common

import (
	"regexp"
	"strconv"
	"strings"
)

// MustAtoi converts the input string to an integer and panics on failure.
func MustAtoi(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}

// AtoiSlice calls MustAtoi for each element in the given slice.
func AtoiSlice(input []string) []int {
	res := make([]int, len(input))
	for i := range input {
		res[i] = MustAtoi(input[i])
	}
	return res
}

// TrimSlice executes strings.TrimSpace on each input element. It returns the input slice for convenience.
func TrimSlice(input []string) []string {
	for i := range input {
		input[i] = strings.TrimSpace(input[i])
	}
	return input
}

// TokeniseLine finds all matches of the given regex in the input string, and returns a slice of non-empty
// groups from within those matches.
func TokeniseLine(line string, re *regexp.Regexp) []string {
	var res []string
	matches := re.FindAllStringSubmatch(line, -1)
	for i := range matches {
		for j := range matches[i] {
			if j > 0 && matches[i][j] != "" {
				res = append(res, matches[i][j])
			}
		}
	}
	return res
}

// TokeniseLines applies TokeniseLine for each line provided.
func TokeniseLines(lines []string, re *regexp.Regexp) [][]string {
	var res [][]string
	for i := range lines {
		res = append(res, TokeniseLine(lines[i], re))
	}
	return res
}
