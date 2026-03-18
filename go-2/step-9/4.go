package main

/*Напишите тест для функции, которая проверяет, являются ли слова анаграммами:

  func AreAnagrams(str1, str2 string) bool {
      str1 = strings.ToLower(str1)
      str2 = strings.ToLower(str2)

      if len(str1) != len(str2) {
          return false
      }

      // Convert strings to slices of runes for sorting
      r1 := []rune(str1)
      r2 := []rune(str2)

      sort.Slice(r1, func(i, j int) bool {
          return r1[i] < r1[j]
      })

      sort.Slice(r2, func(i, j int) bool {
          return r2[i] < r2[j]
      })

      return string(r1) == string(r2)
  }*/
import "testing"

func TestAreAnagrams(t *testing.T) {
	t.Run("Anograms", func(t *testing.T) {
		t.Parallel()
		s1 := "пила"
		s2 := "липа"
		if !AreAnagrams(s1, s2) {
			t.Errorf("Wrong answer with anograms")
		}
	})
	t.Run("Not Anograms(different lengths)", func(t *testing.T) {
		t.Parallel()
		s1 := "липа"
		s2 := "нухзчетутписать"
		if AreAnagrams(s1, s2) {
			t.Errorf("Wrong answer for different lengths")
		}
	})
	t.Run("Not Anograms(different words, but with equal lengths)", func(t *testing.T) {
		t.Parallel()
		s1 := "липа"
		s2 := "mama"
		if AreAnagrams(s1, s2) {
			t.Errorf("Wrong answer for different lengths")
		}
	})

}
