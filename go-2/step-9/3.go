package main

/*Напишите тест для функции:

  func ReverseString(input string) string {
      runes := []rune(input)
      for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
          runes[i], runes[j] = runes[j], runes[i]
      }
      return string(runes)
  }*/
import "testing"

func TestContains(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		excepted := "1234567890"
		input := "0987654321"
		input = ReverseString(input)
		if input != excepted {
			t.Errorf("Excepted not equal to input!")
		}
	})
}
