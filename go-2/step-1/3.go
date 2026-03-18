package main

/*Создайте структуру UpperWriter с полем UpperString string

type UpperWriter struct {
    UpperString string
}
и реализуйте интерфейс io.Writer. Метод Write должен переводить строку в верхний регистр и записывать данные в поле UpperString. Если возникает ошибка, функция должна возвращать её.*/
import (
	"strings"
	"unicode/utf8"
)

type UpperWriter struct {
	UpperString string
}

func (uw *UpperWriter) Write(p []byte) (int, error) {
	uw.UpperString = strings.ToUpper(string(p))
	return utf8.RuneCountInString(uw.UpperString), nil
}
