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
	var sums []int

	for _, numbers := range numbersToSum {
		// append crea un nuevo array, tomando el original y concatenando lo nuevo
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// crea un nuevo slice, aca digo desde el primer elemento hasta el final
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		}
	}

	return sums
}
