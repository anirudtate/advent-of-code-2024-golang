package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	r := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\))|(do\(\))|(don\'t\(\))`)
	muls := r.FindAllString(input, -1)
	result := 0

	on := true

	for i := 0; i < len(muls); i++ {
		if muls[i] == "do()" {
			on = true
			continue
		}
		if muls[i] == "don't()" {
			on = false
			continue
		}
		if !on {
			continue
		}
		r2 := regexp.MustCompile(`[0-9]+`)
		nums := r2.FindAllString(muls[i], -1)
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		result += n1 * n2
	}
	fmt.Println(result)
}
