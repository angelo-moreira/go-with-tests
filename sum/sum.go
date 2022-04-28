package main

func Sum(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tl := numbers[1:]
			sums = append(sums, Sum(tl))
		}
	}

	return sums
}
