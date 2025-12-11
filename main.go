package main

import (
	"fmt"
	"os"
	"strings"

	"opertaions/opertaions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage error: please provide exactly 1 file name as argument.")
		return
	}
	name := os.Args[1]
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	splite := strings.Fields(string(file))
	numbers, err := opertaions.Convert(splite)
	if err != nil {
		fmt.Println(err)
		return
	}

	Average := opertaions.Average(numbers)
	median := opertaions.Median(numbers)
	variance := opertaions.Variance(numbers, Average)
	StandardDeviation := opertaions.StandardDeviation(variance)

	fmt.Println("Average:", Average)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", StandardDeviation)
}
