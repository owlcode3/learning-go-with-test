package main

import "fmt"

func main() {

	fmt.Println(SumAll([]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}))
	fmt.Println(SumAllTails([]int{}, []int{0, 9}))

}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}

	}

	return sums
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}
