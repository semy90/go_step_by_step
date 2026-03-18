package main

/*Напишите тест для функции:

    func Contains(numbers []int, target int) bool{
    for _, num := range numbers {
        if num == target {
            return true
        }
    }
    return false
}*/
import "testing"

func TestContains(t *testing.T) {
	numbers := []int{1, 3, 4, 5, 0, 10}
	t.Run("The value was not found, although it is there.", func(t *testing.T) {
		if !Contains(numbers, 4) {
			t.Errorf("Value not found, data: %v, target: 4", numbers)
		}
	})
	t.Run("A value was found although it does not exist.", func(t *testing.T) {
		if Contains(numbers, 10000000) {
			t.Errorf("wrong answer, value: 10000000 not in numbers: %v", numbers)
		}
	})
}
