// Задача 3: Проверка реализации через компилятор
//
// Ожидаемый вывод:
//   Running fast

package main

import "fmt"

// TODO: объяви интерфейс Runner с методом Run() string

type Runner interface {
	Run() string
}

// TODO: объяви структуру Athlete (без полей)
// TODO: реализуй метод Run() string для Athlete — возвращай "Running fast"

type Athlete struct{}

func (a Athlete) Run() string {
	return "Running fast"
}

// Проверка через компилятор: раскомментируй строку ниже после реализации
var _ Runner = Athlete{}

func main() {
	// TODO: создай Athlete и вызови Run() через интерфейс Runner
	var r Runner = Athlete{}
	fmt.Println(r.Run())
}
