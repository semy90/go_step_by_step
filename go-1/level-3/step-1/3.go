package main

/*
Программист Арсений любит числа, а вот его друг Артём — нет. Они хотят решать домашку вместе, но всё время ссорятся из-за несходства вкусов.

Помогите им и напишите функцию NumbersToLetters(input string) string, которая получает на вход пример в виде цифр и арифметических знаков, заменяет их на слова и возвращает. Замена должна происходить так:

0  |  ноль
1  |  один
2  |  два
3  |  три
4  |  четыре
5  |  пять
6  |  шесть
7  |  семь
8  |  восемь
9  |  девять
+  |  плюс
-  |  минус
*  |  умножить на
/  |  разделить на
Все остальные символы нужно оставить в первоначальном виде: если на вход подаётся (1 + 2) * 3 / 8, нужно вернуть (один плюс два) умножить на три разделить на восемь.
*/
import (
	"strings"
)

func NumbersToLetters(input string) string {
	input = strings.Replace(input, "0", "ноль", -1)
	input = strings.Replace(input, "1", "один", -1)
	input = strings.Replace(input, "2", "два", -1)
	input = strings.Replace(input, "3", "три", -1)
	input = strings.Replace(input, "4", "четыре", -1)
	input = strings.Replace(input, "5", "пять", -1)
	input = strings.Replace(input, "6", "шесть", -1)
	input = strings.Replace(input, "7", "семь", -1)
	input = strings.Replace(input, "8", "восемь", -1)
	input = strings.Replace(input, "9", "девять", -1)
	input = strings.Replace(input, "+", "плюс", -1)
	input = strings.Replace(input, "-", "минус", -1)
	input = strings.Replace(input, "*", "умножить на", -1)
	input = strings.Replace(input, "/", "разделить на", -1)

	return input
}
