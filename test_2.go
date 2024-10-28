package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decode(encoded string) string {
	n := len(encoded) + 1
	nums := make([]int, n)

	nums[0] = 1

	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 'L' {
			nums[i+1] = nums[i] - 1
		} else if encoded[i] == 'R' {
			nums[i+1] = nums[i] + 1
		} else {
			nums[i+1] = nums[i]
		}
	}

	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	shift := -minVal 
	for i := range nums {
		nums[i] += shift
	}

	result := strings.Builder{}
	for _, num := range nums {
		result.WriteString(strconv.Itoa(num))
	}

	return result.String()
}

func main() {
	var input string
	fmt.Print("input = ")
	fmt.Scanln(&input)
	output := decode(input)

	fmt.Println("output = ", output)
}
