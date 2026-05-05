package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Input struct {
	Ranges []string
	Items []string
}

type Item int

type Range struct {
	left int
	right int
}

func NewRange(in string) Range {
	s := strings.Split(in, "-")
	zero, _ := strconv.Atoi(s[0])
	one, _ := strconv.Atoi(s[1])
	left := min(zero, one)
	right := max(zero, one)
	return Range{left: left, right: right}
}

func CompareByLeftBorder(a, b Range) int {
	leftCompare := cmp.Compare(a.left, b.left)
	if leftCompare != 0 {
		return leftCompare
	}
	return cmp.Compare(a.right, b.right)
}

func CompareRight(a, b Range) int {
	rightCompare := cmp.Compare(a.right, b.right)
	if rightCompare != 0 {
		return rightCompare
	}
	return cmp.Compare(a.left, b.left)
}

func ComparItemInRange(r Range, i Item) int {
	integer := int(i)
	if integer <= r.right && integer >= r.left {
		return 0
	}
	if integer > r.right {
		return -1
	}
	if integer < r.left {
		return 1
	}
	return 0
}

func main() {
	input, err := getInput("input.txt")
	if err != nil {
		os.Exit(1)
	}

	var items []Item
	for _, i := range(input.Items) {
		integer, _ := strconv.Atoi(i)
		items = append(items, Item(integer))
	}

	var ranges []Range
	for _, i := range(input.Ranges) {
		r := NewRange(i)
		ranges = append(ranges, r)
	}

	revertRange := slices.Clone(ranges)

	slices.SortFunc(ranges, CompareByLeftBorder)
	slices.SortFunc(revertRange, CompareRight)

	for _, r := range(ranges) {
		fmt.Println(r)
	}
	fmt.Println()

	count := CountItemInRanges(items, ranges, revertRange)
	fmt.Printf("Final count: %d\n", count)

	// PART 2
}

func getInput(filename string) (Input, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Input{}, err
	}

	scanner := bufio.NewScanner(file)

	input := Input{}

	isRanges := true

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isRanges = false
			continue
		}
		if isRanges {
			input.Ranges = append(input.Ranges, line)
		} else {
			input.Items = append(input.Items, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return Input{}, err
	}
	return input, nil
}

func CountItemInRanges(items []Item, ranges []Range, revert []Range) int {
	count := 0

	for _, i := range(items) {
		_, found := slices.BinarySearchFunc(ranges, i, ComparItemInRange)
		_, foundRevert := slices.BinarySearchFunc(revert, i, ComparItemInRange)
		if found || foundRevert {
			count += 1
		}
	}

	// for _, i := range(items) {
	// 	for _, r := range(ranges) {
	// 		found := ComparItemInRange(r, i)
	// 		if found == 0 {
	// 			count += 1
	// 			break
	// 		}
	// 	}
	// }
	return count
}
