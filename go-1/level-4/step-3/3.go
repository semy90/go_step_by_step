package main

/*
Библиотеке нужно не только выдавать книги читателям, но и ежемесячно отчитываться о работе. Вам с программистом Арсением нужно создать систему генерации отчётов.

Разработайте программу, которая будет создавать отчёты о пользователях. Внутри у неё должны быть структуры User с данными о пользователях и Report, которая встраивает в себя структуру User. Отчёты о пользователях должны выводиться на основе данных этих структур.

Создайте структуру User с полями:

ID (int)
Name (string)
Email (string)
Age (int)
Реализуйте структуру Report, которая встраивает в себя User и добавляет поля:

ReportID (int)
Date (string)
Добавьте функцию CreateReport(user User, reportDate string) Report, которая принимает пользователя и дату и возвращает отчёт с заполненными данными. Уникальный ReportID можно сгенерировать, например, с учётом времени регистрации (ReportID может быть произвольным, на ваш вкус).

Создайте функцию GenerateUserReports(users []User, reportDate string) []Report, которая принимает список пользователей и дату и возвращает список отчётов о пользователях. Для каждого пользователя в списке создайте отчёт с помощью функции CreateReport и добавьте его в итоговый список.
*/
var (
	count = 0
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}
type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	count++
	return Report{User: user, ReportID: count, Date: reportDate}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	rp := []Report{}
	for i := range users {
		rp = append(rp, CreateReport(users[i], reportDate))
	}
	return rp
}
