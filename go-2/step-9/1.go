package main

/*Функция SortIntegers(numbers []int) (пакет main) сортирует переданный слайс в порядке увелечения значений. Напишите тест, который проверит корректности работы функции.*/
import "testing"

func TestSortIntegers(t *testing.T) {
	t.Parallel()
	expected := []int{1, 2, 3, 4, 5}
	data := []int{5, 2, 3, 1, 4}

	SortIntegers(data)

	for i := 0; i < len(data); i++ {
		if expected[i] != data[i] {
			t.Errorf("Mismatch with expected data:\nexpected:%v\ngot:%v", expected, data)
		}
	}

}
