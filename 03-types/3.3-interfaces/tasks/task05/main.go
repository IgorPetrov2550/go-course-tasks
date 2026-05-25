// Задача 5: Безопасный type assertion
//
// Ожидаемый вывод:
//   length: 5
//   not a string

package main

import "fmt"

// TODO: напиши функцию printStringLength(x any)
//       используй безопасный type assertion (двузначная форма)
//       если x — string, выводи "length: <len>"
//       иначе выводи "not a string"

func printStringLength(x any) {
	r, ok := x.(string)
	if ok {
		fmt.Printf("length: %d\n", len(r))
		return
	}
	fmt.Printf("not a string\n")
}

func main() {
	printStringLength("hello")
	printStringLength(42)
}
