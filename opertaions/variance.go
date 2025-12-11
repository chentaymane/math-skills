package opertaions

import "math"

func Variance(numbers []int, Average int) int {
	var variance float64

	for _, r := range numbers {
		nmbr := r - Average
		if nmbr < 0 {
			nmbr = -nmbr
			variance += float64(nmbr * nmbr)
		} else {
			variance += float64(nmbr * nmbr)
		}

	}
	variance = variance / float64(len(numbers))
	return int(math.Round(variance))
}
