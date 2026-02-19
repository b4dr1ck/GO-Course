package main

import "fmt"

type transformFn func(int) int

func doubleNumber(n int) int {
	return n * 2
}

func tripleNumber(n int) int {
	return n * 3
}

func addTen(n int) int {
	return n + 10
}

func transformNumbers(numbers []int, transform transformFn) []int {
	transformedNumbers := []int{}
	for _, value := range numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}
	return transformedNumbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	double := transformNumbers(numbers, doubleNumber)
	triple := transformNumbers(numbers, tripleNumber)
	plusTen := transformNumbers(numbers, addTen)
	multiplyByTen := transformNumbers(numbers, func(n int) int { return n * 10 })
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Double: ", double)
	fmt.Println("Triple: ", triple)
	fmt.Println("Plus 10: ", plusTen)
	fmt.Println("Multiply 10: ", multiplyByTen)
}
