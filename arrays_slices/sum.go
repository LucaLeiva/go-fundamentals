package arraysslices

func SumArray(numbers [5]int) (sum int) {
	// _ ignora el index
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumSlice(numbers []int) (sum int) {
	// _ ignora el index
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}

	return sums
}
