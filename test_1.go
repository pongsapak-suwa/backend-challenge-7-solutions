package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)
func themMostValuablePath(Array [][]int) int {
	h := len(Array) - 1
	if h == 0 {
		return 0
	}
	for ch := h - 1; ch >= 0; ch--{
		for cw := 0; cw < len(Array[ch]); cw++ {
			Array[ch][cw] += int(math.Max(float64(Array[ch+1][cw]), float64(Array[ch+1][cw+1])))
		}
	}
	return Array[0][0]
}

func one(Array [][]int, example int) {
	result := themMostValuablePath(Array)
	fmt.Printf("Example 1.%d - output = %d\n", example, result)
}

func main() {
	Array := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	one(Array, 1)

	filePath := "files/hard.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err)
	}
	err = json.Unmarshal(bytes, &Array)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err)
	}
	one(Array, 2)
}
