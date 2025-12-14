package opertaions

func Variance(numbers []int, Average float64) float64 {
	var variance float64
	for _, r := range numbers {
		nmbr := float64(r) - Average
		variance += nmbr * nmbr
	}
	variance = variance / float64(len(numbers))
	return variance
}
