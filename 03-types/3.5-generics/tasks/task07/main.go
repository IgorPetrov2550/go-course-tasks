// Задача 7: Интеграционная задача — обобщенный Store
//
// Ожидаемый вывод:
//   strings: [go rust python]
//   ints: [1 2 3]
//   contains "go": true
//   contains "java": false
//   contains 2: true
//   contains 5: false

package main

import "fmt"

// TODO: объяви структуру Store[T any] с полем items []T
// TODO: добавь метод Add(item T) — добавляет элемент в items
// TODO: добавь метод All() []T — возвращает items

type Store[T any] struct {
	items []T
}

func (s *Store[T]) Add(item T) {
	s.items = append(s.items, item)
}

func (s *Store[T]) All() []T {
	return s.items
}

// TODO: напиши функцию Contains[T comparable](items []T, target T) bool
//       возвращает true если target есть в items

func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	// TODO: создай Store[string], добавь "go", "rust", "python"
	//       выведи "strings:", ss.All()

	ss := Store[string]{items: []string{"go", "rust", "python"}}

	fmt.Printf("strings: %v\n", ss.All())

	// TODO: создай Store[int], добавь 1, 2, 3
	//       выведи "ints:", si.All()

	si := Store[int]{items: []int{1, 2, 3}}
	fmt.Printf("strings: %v\n", si.All())

	// TODO: выведи результаты Contains:
	//       Contains(ss.All(), "go")  → "contains \"go\": true"
	//       Contains(ss.All(), "java") → "contains \"java\": false"
	//       Contains(si.All(), 2)     → "contains 2: true"
	//       Contains(si.All(), 5)     → "contains 5: false"

	fmt.Printf("contains \"go\": %v\n", Contains(ss.All(), "go"))
	fmt.Printf("contains \"java\": %v\n", Contains(ss.All(), "java"))
	fmt.Printf("contains 2: %v\n", Contains(si.All(), 2))
	fmt.Printf("contains 5: %v\n", Contains(si.All(), 5))

}
