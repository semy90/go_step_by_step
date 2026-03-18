package main

/*Напишите функцию

func ReadString(r io.Reader) (string, error)
которая читает данные с помощью r и возвращает их в строковом виде. Если возникает ошибка, функция должна возвращать пустую строку и ошибку, иначе — строку и nil.*/
import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	info, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(info), nil
}
