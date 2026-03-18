package main

/*Дополните тест для функции Length из урока, чтобы покрытие кода составляло 100%.

func Length(a int) string {
    switch {
    case a < 0:
        return "negative"
    case a == 0:
        return "zero"
    case a < 10:
        return "short"
    case a < 100:
        return "long"
    }
    return "very long"
}*/
import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
	{0, "zero"},
	{90, "long"},
	{100, "very long"},
}

func TestLenght(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
