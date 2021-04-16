package arrays

func SumAllTails(allnumbers ...[]int) (tail []int) {

	for _, numbers := range allnumbers {
		if len(numbers) <= 1 {
			tail = append(tail, 0)
		} else {
			tail = append(tail, Sum(numbers[1:]))
		}

	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
