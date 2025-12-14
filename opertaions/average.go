package opertaions

import "math"

func sum(s []int) float64 {
	var total float64
	for _, r := range s {
		total += float64(r)
	}
	return total
}

func Average(numbers []int) (int, float64) {
	var Average float64
	Average = sum(numbers) / float64(len(numbers))
	return int(math.Round(Average)), Average
}
