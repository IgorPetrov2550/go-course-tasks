// Задание 1: Типизированная ошибка валидации
//
// Реализуй тип ValidationError с полями Field и Message.
// Он должен удовлетворять интерфейсу error (метод Error() string).
//
// Напиши функцию validateUser(name, email string) error:
//   - если name == ""           -> *ValidationError{Field: "name",  Message: "обязательное поле"}
//   - если в email нет символа '@' -> *ValidationError{Field: "email", Message: "неверный формат (нет @)"}
//   - иначе -> nil
//
// В main() вызови validateUser с заведомо невалидными данными,
// поймай ошибку через errors.As и выведи поля.
//
// Ожидаемый вывод:
//   ошибка валидации: поле "email" - неверный формат (нет @)
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
	"strings"
)

// TODO: определи тип ValidationError и метод Error() string

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("поле %q : %s", e.Field, e.Message)
}

// TODO: напиши функцию validateUser(name, email string) error

func validateUser(name, email string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "обязательное поле"}
	}

	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Message: "неверный формат (нет @)"}
	}
	return nil
}

func main() {
	// TODO: вызови validateUser с невалидным email, например validateUser("Аня", "anya.mail")
	//
	// var vErr *ValidationError
	// if errors.As(err, &vErr) {
	//     fmt.Printf(...)
	// }	\
	if err := validateUser("Аня", "anya.mail"); err != nil {
		var vErr *ValidationError
		if errors.As(err, &vErr) {
			//fmt.Println(vErr)
			fmt.Printf("ошибка валидации: поле %q - %s\n", vErr.Field, vErr.Message)
		} else {
			fmt.Println(err)
		}
	}
}
