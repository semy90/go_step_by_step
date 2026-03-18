package main

/*
В прошлой задаче вы уже помогли программисту Арсению написать программу для проверки паролей. Но этого оказалось мало! Доработайте программу так, чтобы она сообщала, насколько хорош предъявленный пароль.
Реализуйте функции hasUpper и hasMinimumLength из прошлой задачи и добавьте к ним функцию hasLowerCase, которая проверяет наличие хотя бы одной строчной буквы, принимая на вход строку с паролем. Напишите функцию ratePassword, которая использует для проверки функции hasUpper, hasLowerCase и hasMinimumLength и, в зависимости от количества успешных проверок, возвращает строку Слабый пароль, Средний пароль или Сложный пароль (одна, две или три успешные проверки соответственно). Если успешно не пройдена ни одна из проверок, нужно вернуть строку: Пароль недостаточно безопасен, придумайте новый.
Все вспомогательные функции также возвращают true, если пароль соответствует условию, и false, если нет.
*/
import (
	"unicode"
)

func hasMinimumLength(password string, lenght int) bool {
	return len(password) >= lenght
}

func hasUpper(password string) bool {
	for _, letter := range password {
		if unicode.IsUpper(letter) {
			return true
		}
	}
	return false
}

func hasLowerCase(password string) bool {
	for _, letter := range password {
		if unicode.IsLower(letter) {
			return true
		}
	}
	return false
}

func ratePassword(password string) string {
	k := 0
	if hasLowerCase(password) {
		k++
	}
	if hasUpper(password) {
		k++
	}
	if hasMinimumLength(password, 8) {
		k++
	}

	if k == 0 {
		return "Пароль недостаточно безопасен, придумайте новый"
	} else if k == 1 {
		return "Слабый пароль"
	} else if k == 2 {
		return "Средний пароль"
	} else {
		return "Сложный пароль"
	}
}
