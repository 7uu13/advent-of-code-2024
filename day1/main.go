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

func partOne() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	var left []int
	var right []int
	var sum int

	for _, line := range lines {
		if line != "" {
			parsedRow := strings.Split(line, "   ")
			num1, _ := strconv.Atoi(parsedRow[0])
			num2, _ := strconv.Atoi(parsedRow[1])

			left = append(left, num1)
			right = append(right, num2)

		}
	}
	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(right); i++ {
		diff := int(math.Abs(float64(left[i] - right[i])))
		sum += diff
	}

	fmt.Println(sum)

}

func partTwo() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	var left []int
	var right []int
	var sum int
	count := make(map[int]int)

	for _, line := range lines {
		if line != "" {
			parsedRow := strings.Split(line, "   ")
			num1, _ := strconv.Atoi(parsedRow[0])
			num2, _ := strconv.Atoi(parsedRow[1])

			left = append(left, num1)
			right = append(right, num2)

		}
	}

	for _, num := range right {
		count[num]++
	}

	for _, num := range left {
		occurences := count[num]
		result := num * occurences
		sum += result
	}

	fmt.Println(sum)

}
func main() {
	partOne()
	partTwo()
}
