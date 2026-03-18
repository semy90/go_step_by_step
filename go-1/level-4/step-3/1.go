package main

/*
Программист Арсений решил помочь районной библиотеке автоматизировать выдачу и приём книг. Для этого, перво-наперво, нужно сохранить цифровую информацию о каждой книге. Арсений поручил это сложное задание вам.

Создайте структуру Book с полями:

Title (string)
Author (string)
Year (int)
Genre (string)
Вам нужно создать конструктор для этой структуры. Он должен инициализировать поля структуры, когда будет появляться её новый экземпляр, принимать значения всех полей и возвращать указатель на экземпляр структуры.
*/
type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(title, author string, year int, genre string) *Book {
	return &Book{Title: title, Author: author, Year: year, Genre: genre}
}
