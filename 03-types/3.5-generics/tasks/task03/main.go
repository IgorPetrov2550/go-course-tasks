// Задача 3: Сумма чисел
//
// Ожидаемый вывод:
//   sum int: 15
//   sum float: 7.5

package main

import "fmt"

// TODO: объяви ограничение Number для ~int | ~float64

type Number interface {
	~int | ~float64
}

// TODO: напиши функцию Sum[T Number](items []T) T
//       возвращает сумму всех элементов среза

func Sum[T Number](items []T) T {
	var total T // инициализируется нулем (0 или 0.0 в зависимости от типа)
	for _, v := range items {
		total += v
	}
	return total
}

func main() {
	// TODO: вызови Sum для []int{1, 2, 3, 4, 5} и выведи "sum int: <результат>"
	// TODO: вызови Sum для []float64{1.5, 2.0, 4.0} и выведи "sum float: <результат>"
	fmt.Printf("sum int: %v\n", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("sum float: %v\n", Sum([]float64{1.5, 2.0, 4.0}))
}
