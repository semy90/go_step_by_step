package main

/*Напишите функцию

func Copy(r io.Reader, w io.Writer, n uint) error
которая копирует n байт из r в w.

Если количество байт, доступных для чтения, меньше n, функция должна копировать все данные. Если возникает ошибка — возвращать её.*/
import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	lr := io.LimitReader(r, int64(n))
	data, err := io.ReadAll(lr)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
