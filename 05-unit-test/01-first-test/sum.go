package unittest

func sum(numbers ...int) int {
	var sum int

	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}
