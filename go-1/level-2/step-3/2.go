package main

/*
Программист Арсений пишет приложения для крупного банка. Ему нужно сделать такую функцию деления, которая правильно возвращает ошибки и никогда не вызывает панику.
Напишите функцию Divide(a, b int) (float64, error), которая получает два целых числа и выводит результат их деления. Если делитель равен нулю, функция должна вернуть в качестве ответа 0 и сообщение об ошибке (division by zero is not allowed). Если второе число не равно нулю, верните в качестве ошибки nil. Ошибку нужно вынести в переменную: DivisionByZeroError.
*/
import (
	"errors"
)

var (
	DivisionByZeroError = errors.New("division by zero is not allowed")
)

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	} else {
		return float64(a) / float64(b), nil
	}
}
