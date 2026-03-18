package main

/*Реализуйте функцию EncodeStudentsToWriter(w io.Writer, students []Student) error,
которая принимает io.Writer и слайс структур Student, и записывает этот слайс в формате JSON в writer с помощью json.Encoder.
Функция должна возвращать ошибку, если сериализация или запись не удалась.

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}*/
import (
	"encoding/json"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func EncodeStudentsToWriter(w io.Writer, students []Student) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&students)
	if err != nil {
		return err
	}
	return nil
}
