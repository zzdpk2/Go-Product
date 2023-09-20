package main

func SumArray(numbers [5]int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumSlice(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// lenOfNumbers := len(numbersToSum)
	// sums = make([]int, lenOfNumbers)
	//
	// for i, numbers := range numbersToSum {
	// 	sums[i] += SumSlice(numbers)
	// }
	// return sums
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlice(numbers)
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumSlice(numbers[1:]))
		}
	}
	return
}
