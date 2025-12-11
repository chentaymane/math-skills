package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(s []int) float64 {
	var total float64
	for _, r := range s {
		total += float64(r)
	}
	return total
}

func main() {
	name := os.Args[1]
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	splite := strings.Fields(string(file))
	numbers := []int{}

	for _, r := range splite {
		number, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, (number))
	}
	var Average float64

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	Average = sum(numbers) / float64(len(numbers))

	median := 0

	if len(numbers)%2 != 0 {
		median = numbers[len(numbers)/2]
	} else {
		left := (len(numbers) / 2) 
		right := (len(numbers) / 2)+1
		median = (numbers[left] + numbers[right]) / 2
	}

	var variance float64


	for _ , r:= range numbers {
		nmbr := r- int(math.Round(Average))
		if nmbr < 0 {
			nmbr = -nmbr
			variance += float64(nmbr*nmbr)
		}else{
			variance += float64(nmbr*nmbr)
		}

	}
	variance = variance / float64(len(numbers))

	SD :=math.Sqrt(float64(variance))




	fmt.Println("Average:",int(math.Round(Average)))
	fmt.Println("Median:",median)
	fmt.Println("Variance:",int(math.Round(variance)))
	fmt.Println("Standard Deviation:",int(math.Round(SD)))
}
