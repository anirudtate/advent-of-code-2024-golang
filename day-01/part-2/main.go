package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	input_lines := strings.Split(input, "\n")

	for line := 0; line < len(input_lines); line++ {
		numbers := strings.Split(input_lines[line], "   ")

		n1, _ := strconv.Atoi((string(numbers[0])))
		a1 = append(a1, n1)

		n2, _ := strconv.Atoi((string(numbers[1])))
		a2 = append(a2, n2)
	}

	slices.Sort(a1)
	slices.Sort(a2)

	result := 0

	for i := 0; i < len(a1); i++ {
		times := 0
		for j := 0; j < len(a2); j++ {
			if a1[i] == a2[j] {
				times++
			}
		}
		result += (a1[i] * times)
	}

	fmt.Println(result)
}
