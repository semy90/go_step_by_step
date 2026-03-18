package main

/*
Придумывать надёжные пароли для десятков сервисов и сайтов — непростая задача. Программист Арсений просит помочь с ней.
Напишите программу, которая проверяет пароль на соответствие условиям. Для этого реализуйте функцию checkPassword, которая принимает на вход строку с паролем и возвращает true, если пароль им соответствует, и false, если нет.
Для проверки условий реализуйте дополнительную функцию hasMinimumLength, которая принимает на вход строку с паролем, минимальную длину пароля и проверяет, длиннее ли пароль необходимого значения. В функции hasMinimumLength эту проверку нужно выполнить со значением 8.
Также реализуйте функцию hasUpper, которая проверяет, есть ли в пароле хотя бы одна заглавная буква. Гарантируется, что на вход будут поданы только латинские буквы.
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

func checkPassword(password string) bool {
	return hasUpper(password) && hasMinimumLength(password, 8)
}
