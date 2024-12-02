package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	input_lines := strings.Split(input, "\n")

	result := 0

	for line := 0; line < len(input_lines); line++ {
		array_str := strings.Split(input_lines[line], " ")
		array := make([]int, 0)
		for i := range len(array_str) {
			num, _ := strconv.Atoi(array_str[i])
			array = append(array, num)
		}
		if check_safty(array) {
			result++
		} else {
			for i := range len(array) {
				if check_safty(removeIndex(array, i)) {
					result++
					break
				}
			}

		}
	}

	fmt.Println(result)
}

func check_safty(array []int) bool {
	ok := true
	for i := range len(array) - 1 {
		if !((array[i] > array[i+1]) == (array[0] > array[1])) {
			ok = false
		}

		abs := math.Abs(float64(array[i] - array[i+1]))
		if abs < 1 || abs > 3 {
			ok = false
		}
	}
	return ok
}

func removeIndex(array []int, index int) []int {
	new_array := make([]int, 0)
	for i := range len(array) {
		if i != index {
			new_array = append(new_array, array[i])
		}
	}
	return new_array
}
