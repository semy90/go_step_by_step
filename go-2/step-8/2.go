package main

/*Напишите программу для чтения JSON-файла и отправкой содержимого в канал. Используйте концепцию тайм-аута вместе с контекстом и отменой при чтении.

Реализуйте функцию:

func readJSON(ctx context.Context, path string, result chan<- []byte)
ctx — контекст path — путь к json result — канал, в который нужно вывести прочитанное значение*/
import (
	"context"
	"os"
)

func read(ctx context.Context, path string, res chan []byte, errch chan error) {
	n, err := os.ReadFile(path)
	defer close(res)
	if err != nil {
		errch <- err
		return
	}
	res <- n
	return
}

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	res := make(chan []byte, 1)
	errch := make(chan error, 1)
	defer close(result)
	go read(ctx, path, res, errch)
	select {
	case <-ctx.Done():
		return
	case n := <-res:
		result <- n
		return
	case <-errch:
		return
	}
}
