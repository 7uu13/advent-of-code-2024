package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) {
	var sum int
	file, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		if line != "" {
			slc := strings.Split(line, " ")
			for i := 1; i < len(slc); i++ {
				intnum, _ := strconv.Atoi(slc[i])
				sum += intnum
			}

		}
	}
}

func main() {
	fmt.Println("day2")
	parseInput("input.txt")
}
