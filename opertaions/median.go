package opertaions

import "math"

func Median(numbers []int) int {
	median := 0
	temp := append([]int{}, numbers...)
	for i := 0; i < len(temp); i++ {
		for j := i + 1; j < len(temp); j++ {
			if temp[j] < temp[i] {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
	}
	if len(temp)%2 != 0 {
		median = temp[len(temp)/2]
	} else {
		left := (len(temp) / 2) - 1
		right := (len(temp) / 2)
		median = int(math.Round(float64(temp[left]+temp[right]) / 2.0))
	}
	return median
}