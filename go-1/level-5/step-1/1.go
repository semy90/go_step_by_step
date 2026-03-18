package main

/*Функция Sum(a, b int) int (пакет main) возвращает результат суммирования чисел a и b. Напишите тест TestSum(t *testing.T) для проверки корректности работы.*/
import "testing"

func TestSum(t *testing.T) {
	a := 1
	b := 3
	if a+b != Sum(a, b) {
		t.Fatalf("ошибка!")
	}
}
