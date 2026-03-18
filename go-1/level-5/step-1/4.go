package main

/*Напишите тест для функции DeleteVowels(s string) string, которая должна удалять все гласные из строки английского языка (y гласной не считается).*/
import "testing"

func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "aaaaiiiiioooooooouuuuuuuuueeeeeeeeeeee",
			out: "",
		},
		{
			in:  "aaay",
			out: "y",
		},
		{
			in:  "mama clear rama",
			out: "mm clr rm",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "y",
			out: "y",
		},
	}
	for i, test := range tests {
		if test.out != DeleteVowels(test.in) {
			t.Errorf("Номер теста %d. Нужно: %s, Выдает %s. При s=%s", i+1, test.out, DeleteVowels(test.in), test.in)
		}
	}
}
