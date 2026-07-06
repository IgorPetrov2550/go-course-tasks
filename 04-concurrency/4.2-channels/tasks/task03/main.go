// Задание 3: Pipeline (конвейер на каналах)
//
// Реализуй конвейер из трёх этапов - каждый этап это горутина
// со своим входным и выходным каналом:
//
//   1. generate(nums ...int) <-chan int
//        - запускает горутину, отправляет числа в канал, закрывает канал
//
//   2. square(in <-chan int) <-chan int
//        - читает из in, возводит в квадрат, пишет в свой out
//        - закрывает out когда in закрыт
//
//   3. filter(in <-chan int, pred func(int) bool) <-chan int
//        - пропускает только те значения, для которых pred возвращает true
//
// В main() собери конвейер:
//   generate(1..10) -> square -> filter (только чётные)
//
// Ожидаемый вывод:
//   4 16 36 64 100
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: func generate(nums ...int) <-chan int
func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

// TODO: func square(in <-chan int) <-chan int
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

// TODO: func filter(in <-chan int, pred func(int) bool) <-chan int
func filter(in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			if pred(n) {
				out <- n
			}
		}
	}()

	return out
}
func main() {
	// TODO: собери конвейер и напечатай результат через for range
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	isEven := func(n int) bool {
		return n%2 == 0
	}

	genCh := generate(nums...)
	sqCH := square(genCh)
	filterCh := filter(sqCH, isEven)

	for result := range filterCh {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
