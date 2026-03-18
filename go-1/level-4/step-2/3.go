package main

/*
Команда Арсения получила заказ на большой проект. Команда уверена в своих силах, но пессимист Арсений заранее готовит код для записи ошибок и уверяет, что их будет предостаточно. Поможете ему подготовиться к худшему?

Создайте интерфейс Logger с методом Log(message string), который будет записывать сообщение в журнал. Создайте пользовательский тип LogLevel типа string, сделайте константные значения типа LogLevel со значениями Error и Info. Создайте структуру Log с полем Level типа LogLevel. Реализуйте метод Log с параметром типа string (текст ошибки), который в зависимости от текущего LogLevel выводит сообщение ERROR: *текст ошибки* или INFO: *текст ошибки*.
*/
import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Info  LogLevel = "INFO"
	Error LogLevel = "ERROR"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(mes string) {
	fmt.Printf("%s: %s", l.Level, mes)
}
