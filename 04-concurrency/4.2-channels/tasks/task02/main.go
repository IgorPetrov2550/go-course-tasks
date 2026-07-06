// Задание 2: Fan-In - слить два канала
//
// Создай две функции-генератора:
//   - evenNumbers() <-chan int  - генерирует чётные числа 2, 4, 6, 8, 10
//   - oddNumbers() <-chan int   - генерирует нечётные числа 1, 3, 5, 7, 9
//
// Каждая функция запускает горутину, которая кладёт числа в канал и закрывает его.
//
// Напиши функцию merge(ch1, ch2 <-chan int) <-chan int,
// которая читает из обоих каналов и возвращает один объединённый канал.
//
// В main() слей оба генератора и выведи все числа в одну строку.
//
// Ожидаемый вывод (порядок может быть любым):
//   2 1 4 3 6 5 8 7 10 9
//   (или любой другой порядок - главное все 10 чисел)
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"sync"
)

// TODO: напиши функцию evenNumbers() <-chan int
func evenNumber() <-chan int {
	evenNumbers := make(chan int)
	go func() {
		defer close(evenNumbers)
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				evenNumbers <- i
			}
		}
	}()
	return evenNumbers
}

// TODO: напиши функцию oddNumbers() <-chan int

func oddNumber() <-chan int {
	oddNumbers := make(chan int)
	go func() {
		defer close(oddNumbers)
		for i := 1; i <= 10; i++ {
			if !(i%2 == 0) {
				oddNumbers <- i
			}
		}
	}()
	return oddNumbers
}

// TODO: напиши функцию merge(ch1, ch2 <-chan int) <-chan int
// Подсказка: используй WaitGroup и отдельную горутину для закрытия merged

func merge(ch1, ch2 <-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			merged <- n
		}
	}

	go output(ch1)
	go output(ch2)

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	// TODO: создай два канала через генераторы
	// TODO: слей их через merge
	// TODO: выведи все числа в одну строку
	ch1, ch2 := oddNumber(), evenNumber()
	merged := merge(ch1, ch2)

	for val := range merged {
		fmt.Printf("%d ", val)
	}
}
