// Задание 4: Контекст через три слоя
//
// Реализуй цепочку handler -> service -> repository, где каждая функция
// принимает ctx context.Context первым параметром.
//
// 1. Объяви свой тип ключа и константу:
//      type ctxKey string
//      const requestIDKey ctxKey = "request-id"
//
// 2. handler(ctx): кладёт в контекст requestID="req-99" через context.WithValue
//    и вызывает service(ctx).
//
// 3. service(ctx): вызывает repository(ctx). Может добавить WithTimeout.
//
// 4. repository(ctx):
//      - достаёт requestID из контекста
//      - печатает "[req-99] получаем данные из БД..."
//      - имитирует работу через time.Sleep
//      - проверяет ctx.Done(): если контекст отменён - возвращает ctx.Err()
//      - иначе печатает "[req-99] готово" и возвращает nil
//
// Ожидаемый вывод:
//   [req-99] получаем данные из БД...
//   [req-99] готово
//
// Запусти: go run main.go

package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: объяви тип ключа и константу requestIDKey

type ctxKey string

const requestIDKey ctxKey = "request-id"

// TODO: handler(ctx context.Context) error

func handler(ctx context.Context) error {
	ctx = context.WithValue(context.Background(), requestIDKey, "req-99")
	return service(ctx)
}

// TODO: service(ctx context.Context) error

func service(ctx context.Context) error {
	ctx, cansel := context.WithTimeout(ctx, 1*time.Second)
	defer cansel()
	return repository(ctx)
}

// TODO: repository(ctx context.Context) error

func repository(ctx context.Context) error {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		requestID = "unknown-req"
	}

	fmt.Printf("[%s]  получаем данные из БД...\n", requestID)

	select {
	case <-time.After(1 * time.Second):
		fmt.Printf("[%s] готово\n", requestID)
		return nil
	case <-ctx.Done():
		fmt.Printf("[%s] ошибка: %v\n", requestID, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// TODO: вызови handler(context.Background()) и обработай ошибку
	if err := handler(context.Background()); err != nil {
		fmt.Println("Запрос завершился с ошибкой:", err)
	}
}
