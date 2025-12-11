package opertaions

import (
	"strconv"
)

func Convert(splite []string) ([]int, error) {
	numbers := []int{}
	for _, r := range splite {
		number, err := strconv.Atoi(r)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, (number))
	}
	return numbers, nil
}
