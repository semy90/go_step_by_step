package main

/*
Функция Multiply(a, b int) int (пакет main) возвращает произведение двух переданных чисел. Напишите тест TestMultiply(t *testing.T) для проверки корректности работы.
*/
import "testing"

type Test struct {
	a   int
	b   int
	out int
}

var tests = []Test{
	{1, 1, 1},
	{1, 2, 2},
	{5, 5, 25},
	{160, 10, 1600},
}

func TestMultiply(t *testing.T) {
	for i, test := range tests {
		if test.out != Multiply(test.a, test.b) {
			t.Errorf("Номер теста %d. Нужно: %d, Выдает %d. При a=%d, b=%d", i+1, test.out, Multiply(test.a, test.b), test.a, test.b)
		}
	}
}
