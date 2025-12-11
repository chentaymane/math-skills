package opertaions

import "math"

func sum(s []int) float64 {
	var total float64
	for _, r := range s {
		total += float64(r)
	}
	return total
}

func Average(numbers []int) int {
	var Average float64

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	Average = sum(numbers) / float64(len(numbers))
	return int(math.Round(Average))
}
