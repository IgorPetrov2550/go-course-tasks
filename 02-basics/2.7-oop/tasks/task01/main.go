// Задание 1: Стек
//
// Реализуй структуру Stack для int с методами:
//   - Push(val int)           - добавить элемент на вершину
//   - Pop() (int, error)      - забрать верхний элемент; ошибка, если стек пуст
//   - Peek() (int, error)     - посмотреть верхний элемент без удаления
//   - Len() int               - количество элементов
//   - IsEmpty() bool          - true, если стек пуст
//
// Методы должны быть определены на *Stack (указателе), чтобы Push/Pop
// изменяли внутреннее состояние.
//
// Заведи переменную ErrEmptyStack через errors.New — её будут возвращать Pop/Peek.
//
// В main() продемонстрируй работу: запушь 3 числа, вызови Peek, сделай Pop
// до опустошения и один Pop на пустом стеке — он должен вернуть ошибку.
//
// Ожидаемый вывод:
//   Len=3 Top=30
//   Pop: 30
//   Pop: 20
//   Pop: 10
//   Pop пустого: stack is empty
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
)

// TODO: var ErrEmptyStack = errors.New("stack is empty")

var ErrEmptyStack = errors.New("stack is empty")

// TODO: type Stack struct { items []int }

type Stack struct {
	items []int
}

// TODO: методы Push, Pop, Peek, Len, IsEmpty на *Stack

func (p *Stack) Push(val int) {
	p.items = append(p.items, val)
}

func (p *Stack) Pop() (int, error) {
	if len(p.items) == 0 {
		return 0, ErrEmptyStack
	}
	last := p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1]
	return last, nil
}

func (p *Stack) Peek() (int, error) {
	if len(p.items) == 0 {
		return 0, ErrEmptyStack
	}
	return p.items[len(p.items)-1], nil
}

func (p *Stack) Len() int {
	return len(p.items)
}

func (p *Stack) IsEmpty() bool {
	return len(p.items) == 0
}

func main() {
	// TODO: s := &Stack{}
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	top, _ := s.Peek()
	fmt.Printf("Len=%d Top=%d\n", s.Len(), top)
	for !s.IsEmpty() {
		pop, _ := s.Pop()
		fmt.Printf("Pop=%d\n", pop)
	}
	if _, err := s.Pop(); err != nil {
		fmt.Printf("Pop пустого: %v\n", err)
	}

	// В main() продемонстрируй работу: запушь 3 числа, вызови Peek, сделай Pop
	// до опустошения и один Pop на пустом стеке — он должен вернуть ошибку.
	//
	// Ожидаемый вывод:
	//   Len=3 Top=30
	//   Pop: 30
	//   Pop: 20
	//   Pop: 10
	//   Pop пустого: stack is empty
	//

	// s.Push(10); s.Push(20); s.Push(30)
	// ...

}
