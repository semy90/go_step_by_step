package main

/*Функция GetUTFLength(input []byte) (int, error) возвращает длину строки UTF8 и ошибку ErrInvalidUTF8 (в случае возникновения). Напишите тест, который бы проверял возвращаемые функцией значения.

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
    if !utf8.Valid(input) {
        return 0, ErrInvalidUTF8
    }

    return utf8.RuneCount(input), nil
}*/
import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{input: []byte("Hello, 世界"), expected: 9, err: nil},

		{input: []byte{0x80}, expected: 0, err: ErrInvalidUTF8},

		{input: []byte{}, expected: 0, err: nil},

		{input: []byte("你"), expected: 1, err: nil},
	}

	for _, tt := range tests {
		got, err := GetUTFLength(tt.input)
		if got != tt.expected {
			t.Errorf("GetUTFLength() got = %v, expected = %v", got, tt.expected)
		}
		if err != tt.err {
			t.Errorf("GetUTFLength() error = %v, expected = %v", err, tt.err)
		}
	}
}
