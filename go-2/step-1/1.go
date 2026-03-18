package main

/*Напишите функцию

func WriteString(s string, w io.Writer) error
которая записывает строку s в место назначения с помощью интерфейса w.

Если возникает ошибка, функция должна возвращать её, иначе — nil.*/
import (
	"io"
)

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	return err
}
