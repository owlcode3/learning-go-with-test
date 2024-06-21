package main

import "fmt"

func main() {
	fmt.Println(SumAll([]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}))
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
