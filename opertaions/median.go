package opertaions

func Median (numbers []int)int {
	median := 0

	if len(numbers)%2 != 0 {
		median = numbers[len(numbers)/2]
	} else {
		left := (len(numbers) / 2)
		right := (len(numbers) / 2) + 1
		median = (numbers[left] + numbers[right]) / 2
	}
	return median
}