package opertaions

import "math"

func Variance(numbers []int, Average int) int {
	var variance float64

	for _, r := range numbers {
		nmbr := r - Average
		variance += float64(nmbr * nmbr)
	}
	variance = variance / float64(len(numbers))
	return int(math.Round(variance))
}
