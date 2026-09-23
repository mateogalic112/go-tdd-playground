package integers

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slicesOfNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range slicesOfNumbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slicesOfNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range slicesOfNumbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
