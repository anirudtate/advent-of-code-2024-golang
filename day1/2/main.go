package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	for {
		b1 := make([]byte, 5)
		_, err := f.Read(b1)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		n1, _ := strconv.Atoi((string(b1)))
		a1 = append(a1, n1)

		_, err = f.Seek(3, io.SeekCurrent)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		b2 := make([]byte, 5)
		_, err = f.Read(b2)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		n2, _ := strconv.Atoi((string(b2)))
		a2 = append(a2, n2)

		_, err = f.Seek(1, io.SeekCurrent)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
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
		fmt.Println(times)
		result += (a1[i] * times)
	}

	fmt.Println(result)
}
