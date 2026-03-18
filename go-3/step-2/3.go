package main

/*Реализуйте функцию DecodeStudentFromReader(r io.Reader) (Student, error),
которая принимает источник данных, реализующий интерфейс io.Reader, и читает из него JSON-строку, декодируя её в структуру Student:

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}
Если данные некорректны, функция возвращает ошибку.*/
import (
	"encoding/json"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func DecodeStudentFromReader(r io.Reader) (Student, error) {
	var student Student
	studentBytes, err := io.ReadAll(r)
	if err != nil {
		return Student{}, err
	}
	err = json.Unmarshal(studentBytes, &student)
	if err != nil {
		return Student{}, err
	}
	return student, nil
}
