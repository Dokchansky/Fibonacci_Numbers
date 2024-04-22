package main

import "fmt"

func isFibonacci(n int) bool {
	a, b := 0, 1
	for b <= n {
		if b == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}

func fibonacciNeighbors(n int) (int, int) {
	a, b := 0, 1
	for b <= n {
		if b == n {
			return a, a + b
		}
		a, b = b, a+b
	}
	return a, a
}

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	prev, next := fibonacciNeighbors(num)

	if isFibonacci(num) {
		fmt.Printf("Предыдущее число: %d\nСледующее число: %d\n", prev, next)
	} else {
		if num-prev < next-num {
			fmt.Printf("Ближайшее число из ряда Фибоначчи: %d\n", prev)
		} else {
			fmt.Printf("Ближайшее число из ряда Фибоначчи: %d\n", next)
		}
	}
}
