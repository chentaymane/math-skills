package opertaions

import "math"

func StandardDeviation(variance int) int {
	StandardDeviation := math.Sqrt(float64(variance))
	return  int(math.Round(StandardDeviation))
}
