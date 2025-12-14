package main

import (
	"fmt"
	"math"
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
	if len(file) == 0 {
		fmt.Println("incorect input")
		return
	}
	splite := strings.Fields(string(file))
	numbers, err := opertaions.Convert(splite)
	if err != nil {
		fmt.Println(err)
		return
	}

	Average, Average2 := opertaions.Average(numbers)
	median := opertaions.Median(numbers)
	variance := opertaions.Variance(numbers, Average2)
	StandardDeviation := opertaions.StandardDeviation(variance)

	fmt.Println("Average:", Average)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", int(math.Round(variance)))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDeviation)))
}
