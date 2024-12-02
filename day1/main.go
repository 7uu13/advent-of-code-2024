package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	left  []int
	right []int
)

func parseInput(input string) ([]int, []int) {
	file, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		if line != "" {
			parsedRow := strings.Split(line, "   ")
			num1, _ := strconv.Atoi(parsedRow[0])
			num2, _ := strconv.Atoi(parsedRow[1])

			left = append(left, num1)
			right = append(right, num2)

		}
	}
	return left, right
}

func partOne(input string) int {
	right, left := parseInput(input)
	sort.Ints(left)
	sort.Ints(right)
	sum := 0

	for i := 0; i < len(right); i++ {
		diff := int(math.Abs(float64(left[i] - right[i])))
		sum += diff
	}
	return sum
}

func partTwo(input string) int {
	right, left := parseInput(input)
	count := make(map[int]int)
	sum := 0

	for _, num := range right {
		count[num]++
	}

	for _, num := rangerange left {
		occurences := count[num]
		result := num * occurences
		sum += result
	}

	return sum

}
func main() {
	input := "input.txt"
	part1 := partOne(input)
	part2 := partTwo(input)

	fmt.Println("Answer for day1 part1:", part1)
	fmt.Println("Answer for day1 part2:", part2)
}
